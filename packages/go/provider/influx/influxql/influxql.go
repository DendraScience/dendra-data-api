package influxql

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

type Driver struct {
	datapointsBatchSize int
	instances           map[string]Instance
}

type Instance struct {
	// TODO: Finish!
}

func (driver *Driver) StreamDatapoints(request *pb.ProviderStreamDatapointsRequest, srv pb.ProviderService_StreamDatapointsServer) error {
	fmt.Println("influxql streaming datapoints")

	return nil
}

func NewDriver() *Driver {
	const itemPrefix = "INFLUXQL_API_"

	datapointsBatchSize, err := strconv.Atoi(os.Getenv("INFLUXQL_DATAPOINTS_BATCH_SIZE"))
	if err != nil || datapointsBatchSize <= 0 {
		datapointsBatchSize = 10
	}
	driver := Driver{
		datapointsBatchSize: datapointsBatchSize,
		instances:           map[string]Instance{},
	}
	items := strings.Split(os.Getenv("INFLUXQL_APIS"), ",")

	for _, item := range items {
		key := strings.ToUpper(strings.TrimSpace(item))
		if key != "" {
			prefix := itemPrefix + key + "_"

			// TODO: Finish!
			url := os.Getenv(prefix + "URL")

			driver.instances[key] = Instance{}

			fmt.Printf("registered influxql instance %s at %s\n", key, url)
		}
	}

	return &driver
}
