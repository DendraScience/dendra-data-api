package flux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/types"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

type Driver struct {
	datapointsBatchSize int
	instances           map[string]Instance
}

type Instance struct {
	client   influxdb2.Client
	queryAPI api.QueryAPI
}

type queryParams struct {
	// assigned from config instance query params
	bucket      string
	measurement string
	fields      []string
	tagSet      map[string]string
	coalesce    bool
	utcOffset   int
	vField      string

	// assigned from request query params
	startTime time.Time
	stopTime  time.Time
	sortAsc   bool
	hasLimit  bool
	limit     uint32
}

const pipe = " |> "

var unsafeRe = regexp.MustCompile(`\W`)

func safeName(s string) string {
	return unsafeRe.ReplaceAllString(s, "_")
}

func encodeString(s string) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(s)
	return string(buffer.Bytes()), err
}

func fromFn(p *queryParams) string {
	return `from(bucket: "` + safeName(p.bucket) + `")`
}

func rangeFn(p *queryParams) (string, error) {
	const errPrefix = "flux.rangeFn"

	start, err := p.startTime.MarshalText()
	if err != nil {
		return "", fmt.Errorf("%s: start time error: %w", errPrefix, err)
	}
	stop, err := p.stopTime.MarshalText()
	if err != nil {
		return "", fmt.Errorf("%s: stop time error: %w", errPrefix, err)
	}

	return `range(start: ` + string(start) + `, stop: ` + string(stop) + `)`, nil
}

func filterMeasurementFn(p *queryParams) string {
	return `filter(fn: (r) => r._measurement == "` + safeName(p.measurement) + `")`
}

func filterTagSetFn(p *queryParams) (string, error) {
	const errPrefix = "flux.filterTagSetFn"

	parts := make([]string, 0, len(p.tagSet))
	for key, value := range p.tagSet {
		str, err := encodeString(value)
		if err != nil {
			return "", fmt.Errorf("%s: encoding error: %w", errPrefix, err)
		}
		parts = append(parts, `r["`+safeName(key)+`"] == `+str)
	}

	if len(parts) > 0 {
		return `filter(fn: (r) => ` + strings.Join(parts, " and ") + `)`, nil
	}
	return "", nil
}

func filterFieldsFn(p *queryParams) string {
	parts := make([]string, len(p.fields))
	for i, key := range p.fields {
		parts[i] = `r._field == "` + safeName(key) + `"`
	}

	if len(parts) > 0 {
		return `filter(fn: (r) => ` + strings.Join(parts, " or ") + `)`
	}
	return ""
}

func limitFn(p *queryParams) string {
	if p.hasLimit {
		if p.sortAsc {
			return fmt.Sprintf("limit(n: %d)", p.limit)
		} else {
			return fmt.Sprintf("tail(n: %d)", p.limit)
		}
	}
	return ""
}

func sortFn(p *queryParams) string {
	if !p.sortAsc {
		return `sort(columns: ["_time"], desc: true)`
	}
	return ""
}

func pivotFn() string {
	return `pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")`
}

func dropFn(p *queryParams) string {
	parts := make([]string, 3, len(p.tagSet)+3)
	parts[0] = `"_measurement"`
	parts[1] = `"_start"`
	parts[2] = `"_stop"`

	for key := range p.tagSet {
		parts = append(parts, `"`+safeName(key)+`"`)
	}
	return `drop(columns: [` + strings.Join(parts, ",") + `])`
}

func (driver *Driver) instanceAndQueryParams(paramsStruct *structpb.Struct, query *pb.DatapointsQuery) (*Instance, *queryParams, error) {
	p := queryParams{}

	//
	// process config instance query params
	//

	// get fields from params.query
	paramsFields := paramsStruct.GetFields()
	if paramsFields == nil {
		return nil, nil, &types.ParamsError{
			Detail: "params fields are nil",
		}
	}
	queryStruct := paramsFields["query"].GetStructValue()
	if queryStruct == nil {
		return nil, nil, &types.ParamsError{
			Detail: "query cannot be empty",
		}
	}
	queryFields := queryStruct.GetFields()
	if queryFields == nil {
		return nil, nil, &types.ParamsError{
			Detail: "query fields are nil",
		}
	}

	// get database and associated influx instance
	databaseString := queryFields["database"].GetStringValue()
	instance, found := driver.instances[strings.ToUpper(databaseString)]
	if !found {
		instance, found = driver.instances["DEFAULT"]
		if !found {
			return nil, nil, &types.ParamsError{
				Detail: "database instance not found",
			}
		}
	}

	// get bucket and measurement
	p.bucket = queryFields["bucket"].GetStringValue()
	if p.bucket == "" {
		return nil, nil, &types.ParamsError{
			Detail: "query.bucket cannot be empty",
		}
	}
	p.measurement = queryFields["measurement"].GetStringValue()
	if p.measurement == "" {
		p.measurement = "logger_data"
	}

	// get fields (at least one required)
	fieldsList := queryFields["fields"].GetListValue()
	if fieldsList != nil {
		values := fieldsList.GetValues()
		p.fields = make([]string, len(values))
		for i, value := range values {
			str := value.GetStringValue()
			if str == "" {
				return nil, nil, &types.ParamsError{
					Detail: "query.fields[] value cannot be empty",
				}
			}
			p.fields[i] = str
		}
	}
	if len(p.fields) == 0 {
		return nil, nil, &types.ParamsError{
			Detail: "query.fields[] must have at least one field specified",
		}
	}

	// get tag set (optional)
	tagSetStruct := queryFields["tag_set"].GetStructValue()
	if tagSetStruct != nil {
		fields := tagSetStruct.GetFields()
		p.tagSet = make(map[string]string)
		for key, value := range fields {
			str := value.GetStringValue()
			if str == "" {
				return nil, nil, &types.ParamsError{
					Detail: "query.tag_set[] value cannot be empty",
				}
			}
			p.tagSet[key] = str
		}
	}

	p.coalesce = queryFields["coalesce"].GetBoolValue()
	p.utcOffset = int(queryFields["utc_offet"].GetNumberValue())
	p.vField = queryFields["v_field"].GetStringValue()

	//
	// process request query params
	//

	if query.GtTime == nil {
		p.startTime = time.Date(1800, time.January, 2, 0, 0, 0, 0, time.UTC)
	} else if query.GtEqual {
		p.startTime = query.GtTime.AsTime() // >=
	} else {
		p.startTime = query.GtTime.AsTime().Add(time.Millisecond) // >
	}

	if query.LtTime == nil {
		p.stopTime = time.Date(2200, time.January, 2, 0, 0, 0, 0, time.UTC)
	} else if query.LtEqual {
		p.stopTime = query.LtTime.AsTime().Add(time.Millisecond) // <=
	} else {
		p.stopTime = query.LtTime.AsTime() // <
	}

	p.sortAsc = query.SortAsc

	if query.Limit != nil {
		p.hasLimit = true
		p.limit = *query.Limit
	}

	return &instance, &p, nil
}

// TODO: handle shift, utc_offset, etc.
func (driver *Driver) StreamDatapoints(request *pb.StreamDatapointsRequest, srv pb.ProviderService_StreamDatapointsServer) error {
	const errPrefix = "flux.StreamDatapoints"

	fmt.Println("flux stream datapoints")

	// build the query
	instance, p, err := driver.instanceAndQueryParams(request.ConfigInstance.Params, request.Query)
	if err != nil {
		return fmt.Errorf("%s: %w", errPrefix, err)
	}
	rangeExpr, err := rangeFn(p)
	if err != nil {
		return fmt.Errorf("%s: %w", errPrefix, err)
	}

	query := fromFn(p) +
		pipe + rangeExpr +
		pipe + filterMeasurementFn(p)

	filterTagSetExpr, err := filterTagSetFn(p)
	if err != nil {
		return fmt.Errorf("%s: %w", errPrefix, err)
	}
	if filterTagSetExpr != "" {
		query = query + pipe + filterTagSetExpr
	}

	filterFieldsExpr := filterFieldsFn(p)
	if filterFieldsExpr != "" {
		query = query + pipe + filterFieldsExpr
	}

	limitExpr := limitFn(p)
	if limitExpr != "" {
		query = query + pipe + limitExpr
	}

	sortExpr := sortFn(p)
	if sortExpr != "" {
		query = query + pipe + sortExpr
	}

	query = query +
		pipe + pivotFn() +
		pipe + dropFn(p)

	// DEBUG
	fmt.Printf("prepared query: %s\n", query)

	// execute the query
	ctx := srv.Context()
	result, err := instance.queryAPI.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: query error: %w", errPrefix, err)
	}
	datapoints := []*pb.Datapoint{}

	// iterate over query result lines, shape and send datapoints
	for result.Next() {
		record := result.Record()
		time := record.Time()
		values := record.Values()

		delete(values, "_time")
		delete(values, "result")
		delete(values, "table")

		var v *float64

		if p.vField != "" {
			value, found := values[p.vField]
			if found {
				pbValue, err := structpb.NewValue(value)
				if err != nil {
					return fmt.Errorf("%s: v field error: %w", errPrefix, err)
				}
				// adapted from: https://github.com/protocolbuffers/protobuf-go/blob/v1.30.0/types/known/structpb/struct.pb.go#L488
				if x, ok := pbValue.GetKind().(*structpb.Value_NumberValue); ok {
					v = &x.NumberValue
					delete(values, p.vField)
				}
			}
		}

		d, err := structpb.NewStruct(values)
		if err != nil {
			return fmt.Errorf("%s: d struct error: %w", errPrefix, err)
		}

		datapoint := pb.Datapoint{
			T: timestamppb.New(time),
			D: d,
			V: v,
		}
		datapoints = append(datapoints, &datapoint)

		if len(datapoints) >= driver.datapointsBatchSize {
			resp := pb.StreamDatapointsResponse{Datapoints: datapoints}
			err = srv.Send(&resp)
			if err != nil {
				return fmt.Errorf("%s: batch send error: %w", errPrefix, err)
			}
			datapoints = nil
		}
	}
	if result.Err() != nil {
		return fmt.Errorf("%s: result error: %w", errPrefix, result.Err())
	} else if len(datapoints) > 0 {
		// flush final batch
		resp := pb.StreamDatapointsResponse{Datapoints: datapoints}
		err = srv.Send(&resp)
		if err != nil {
			return fmt.Errorf("%s: flush send error: %w", errPrefix, err)
		}
	}

	return nil
}

func NewDriver() *Driver {
	const itemPrefix = "FLUX_DATABASE_"

	datapointsBatchSize, err := strconv.Atoi(os.Getenv("FLUX_DATAPOINTS_BATCH_SIZE"))
	if err != nil || datapointsBatchSize <= 0 {
		datapointsBatchSize = 24
	}
	driver := Driver{
		datapointsBatchSize: datapointsBatchSize,
		instances:           map[string]Instance{},
	}
	items := strings.Split(os.Getenv("FLUX_DATABASES"), ",")

	for _, item := range items {
		key := strings.ToUpper(strings.TrimSpace(item))
		if key != "" {
			prefix := itemPrefix + key + "_"

			org := os.Getenv(prefix + "ORG")
			token := os.Getenv(prefix + "TOKEN")
			url := os.Getenv(prefix + "URL")
			client := influxdb2.NewClient(url, token)
			queryAPI := client.QueryAPI(org)

			driver.instances[key] = Instance{
				client:   client,
				queryAPI: queryAPI,
			}

			fmt.Printf("registered flux instance %s for org %q at %s\n", key, org, url)
		}
	}

	return &driver
}
