/*
 *
 * InfluxDB data provider
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/flux"
	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/influxql"
	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/types"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	drivers = map[string]types.Driver{}
)

func main() {
	flag.Parse()

	drivers["/influx/flux"] = flux.NewDriver()
	drivers["/influx/select"] = influxql.NewDriver()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterProviderServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

type server struct {
	pb.UnimplementedProviderServiceServer
}

func (s *server) StreamDatapoints(request *pb.StreamDatapointsRequest, srv pb.ProviderService_StreamDatapointsServer) error {
	fmt.Println("stream datapoints request received")

	if driver, found := drivers[request.ConfigInstance.Path]; found {
		if err := driver.StreamDatapoints(request, srv); err != nil {
			return err
		}
	} else {
		return &types.PathError{
			Path: request.ConfigInstance.Path,
		}
	}

	return nil
}
