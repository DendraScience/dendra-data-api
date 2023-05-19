/*
 *
 * Datapoints service
 *
 */

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dendrascience/dendra-data-api/packages/go/datapoints/types"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

const mathjs = "MATHJS"

type ConnectedClient[T any] struct {
	conn   *grpc.ClientConn
	client T
}

var (
	metadata     *ConnectedClient[pb.MetadataServiceClient]
	providers    map[string]*ConnectedClient[pb.ProviderServiceClient]
	transformers map[string]*ConnectedClient[pb.TransformerServiceClient]
	converters   map[string]*ConnectedClient[pb.ConverterServiceClient]
)

func dialClient[T any](envName string) *ConnectedClient[T] {
	addr := os.Getenv(envName)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to %s: %v\n", addr, err)
	}

	return &ConnectedClient[T]{
		conn: conn,
	}
}

func dialClients[T any](envName string, itemPrefix string) map[string]*ConnectedClient[T] {
	clients := make(map[string]*ConnectedClient[T])
	items := strings.Split(os.Getenv(envName), ",")

	for _, item := range items {
		key := strings.ToUpper(strings.TrimSpace(item))
		if key != "" {
			target := itemPrefix + key
			addr := os.Getenv(target)
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect to %s: %v\n", addr, err)
			}
			clients[key] = &ConnectedClient[T]{
				conn: conn,
			}
		}
	}

	return clients
}

func queryIntervalFromQuery(query *pb.DatapointsQuery, datastream *pb.Datastream) types.Interval {
	interval := types.NewQueryInterval()

	if query.GetGtTime() != nil {
		interval.Start = query.GetGtTime().AsTime()
		interval.LeftOpen = !query.GetGtEqual()
	}
	if query.GetLtTime() != nil {
		interval.End = query.GetLtTime().AsTime()
		interval.RightOpen = !query.GetLtEqual()
	}

	if query.IsLocal {
		offset := time.Duration(datastream.GetStationLookup().GetUtcOffset()) * time.Second
		interval.Start = interval.Start.Add(-offset)
		interval.End = interval.End.Add(-offset)
	}

	return interval
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// set up a connection to the metadata server
	metadata = dialClient[pb.MetadataServiceClient]("METADATA_SERVICE")
	defer metadata.conn.Close()
	metadata.client = pb.NewMetadataServiceClient(metadata.conn)

	// set up connections to provider clients
	providers = dialClients[pb.ProviderServiceClient]("PROVIDER_SERVICES", "PROVIDER_SERVICE_")
	for _, provider := range providers {
		defer provider.conn.Close()
		provider.client = pb.NewProviderServiceClient(provider.conn)
	}

	// set up connections to transformer clients
	transformers = dialClients[pb.TransformerServiceClient]("TRANSFORMER_SERVICES", "TRANSFORMER_SERVICE_")
	for _, transformer := range transformers {
		defer transformer.conn.Close()
		transformer.client = pb.NewTransformerServiceClient(transformer.conn)
	}

	// resolve mathjs transformer; required for "evaluate" actions
	_, found := transformers[mathjs]
	if !found {
		log.Fatalf("mathjs transformer not configured")
	}

	// set up connections to converter clients
	converters = dialClients[pb.ConverterServiceClient]("CONVERTER_SERVICES", "CONVERTER_SERVICE_")
	for _, converter := range converters {
		defer converter.conn.Close()
		converter.client = pb.NewConverterServiceClient(converter.conn)
	}

	// set up server
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 50051
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterDatapointsServiceServer(s, &server{})
	reflection.Register(s)

	go func(l net.Listener) {
		fmt.Printf("server listening at %v\n", l.Addr())
		if err := s.Serve(l); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}(lis)

	<-ctx.Done()
	stop()

	fmt.Println("stopping server...")
	s.Stop()

	fmt.Println("bye!")
}

type server struct {
	pb.UnimplementedDatapointsServiceServer
}

func (s *server) StreamDatapoints(request *pb.StreamDatapointsRequest, srv pb.DatapointsService_StreamDatapointsServer) error {
	fmt.Println("stream datapoints request received")
	// TODO: add jwt/signing validation

	ctx := srv.Context()
	getDatastreamResp, err := metadata.client.GetDatastream(ctx, &pb.GetDatastreamRequest{
		DatastreamId: request.GetDatastreamId(),
	})
	if err != nil {
		return fmt.Errorf("get datastream error: %w", err)
	}

	// DEBUG
	log.Printf("get datasteam response: %v\n", getDatastreamResp)

	datastream := getDatastreamResp.GetDatastream()
	convert := request.GetConvert()
	query := request.GetQuery()
	config := types.MergeConfig(datastream, !query.GetSortAsc())
	queryInterval := queryIntervalFromQuery(query, datastream)

	// DEBUG
	log.Printf("merged config: %v\n", config)
	log.Printf("query interval: %v\n", queryInterval)

	// resolve mathjs transformer; required for "evaluate" actions
	mathjsTransformer, _ := transformers[mathjs]

	// resolve converter
	var converter *ConnectedClient[pb.ConverterServiceClient]
	convertLibrary := convert.GetLibrary()
	if convertLibrary != "" {
		conv, found := converters[strings.ToUpper(convertLibrary)]
		if !found {
			return fmt.Errorf("converter not valid %q", convertLibrary)
		}
		converter = conv

		if convert.GetFromUnit() == "" {
			// attempt to infer the from unit from metadata
			unitTag := datastream.GetTermsInfo().GetUnitTag()
			if unitTag == "" {
				return types.ErrConvertFromEmpty
			}
			convert.FromUnit = unitTag
		}
		if convert.GetToUnit() == "" {
			return types.ErrConvertToEmpty
		}
	}

	for _, inst := range config {
		// intersect intervals; skip querying if empty
		interval := queryInterval.Intersect(types.Interval{
			Start:     inst.GetBeginsAt().AsTime(),
			End:       inst.GetEndsBefore().AsTime(),
			LeftOpen:  false,
			RightOpen: true,
		})
		if interval.IsEmpty() {
			continue
		}

		path := inst.GetPath()
		parts := strings.Split(path, "/")
		if len(parts) < 2 {
			return fmt.Errorf("path not valid %q", path)
		}
		provider, found := providers[strings.ToUpper(parts[1])]
		if !found {
			return fmt.Errorf("provider not valid %q", parts[1])
		}

		newQuery := proto.Clone(query).(*pb.DatapointsQuery)
		newQuery.GtTime = timestamppb.New(interval.Start)
		newQuery.GtEqual = !interval.LeftOpen
		newQuery.LtTime = timestamppb.New(interval.End)
		newQuery.LtEqual = !interval.RightOpen

		providerReq := pb.ProviderStreamDatapointsRequest{
			Query:          newQuery,
			ConfigInstance: inst,
		}

		// DEBUG
		log.Printf("provider req: %v\n", providerReq)

		stream, err := provider.client.StreamDatapoints(ctx, &providerReq)
		if err != nil {
			return fmt.Errorf("stream datapoints error: %w", err)
		}

		actions := inst.GetActions()
		evaluate := actions.GetEvaluate()

		for {
			providerResp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("stream datapoints receive error: %w", err)
			}

			datapoints := providerResp.GetDatapoints()

			if evaluate != "" {
				transformerResp, err := mathjsTransformer.client.TransformMany(ctx, &pb.TransformManyRequest{
					Transform: &pb.TransformArgs{
						Expression: evaluate,
					},
					Datapoints: datapoints,
				})
				if err != nil {
					return fmt.Errorf("mathjs transform many error: %w", err)
				}
				datapoints = transformerResp.GetDatapoints()
			}

			if converter != nil {
				converterResp, err := converter.client.ConvertMany(ctx, &pb.ConvertManyRequest{
					Convert:    convert,
					Datapoints: datapoints,
				})
				if err != nil {
					return fmt.Errorf("convert many error: %w", err)
				}
				datapoints = converterResp.GetDatapoints()
			}

			if actions != nil {
				// merge annotation-quality metadata into datapoints
				q := pb.DatapointQuality{
					AnnotationIds: inst.GetAnnotationIds(),
					Attrib:        actions.GetAttrib(),
					Flag:          actions.GetFlag(),
				}
				for _, datapoint := range datapoints {
					datapoint.Q = &q
				}
			}

			// TODO: handle limit

			resp := pb.StreamDatapointsResponse{Datapoints: datapoints}
			err = srv.Send(&resp)
			if err != nil {
				return fmt.Errorf("stream datapoints send error: %w", err)
			}
		}
	}

	return nil
}
