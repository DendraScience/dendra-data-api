/*
 *
 * Influx data provider service
 *
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/flux"
	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/influxql"
	"github.com/dendrascience/dendra-data-api/packages/go/provider/influx/types"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

var (
	drivers = map[string]types.Driver{}
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	drivers["/influx/flux"] = flux.NewDriver()
	drivers["/influx/select"] = influxql.NewDriver()

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

	pb.RegisterProviderServiceServer(s, &server{})
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
	pb.UnimplementedProviderServiceServer
}

func (s *server) StreamDatapoints(request *pb.ProviderStreamDatapointsRequest, srv pb.ProviderService_StreamDatapointsServer) error {
	fmt.Println("stream datapoints request received")

	if driver, found := drivers[request.GetConfigInstance().GetPath()]; found {
		if err := driver.StreamDatapoints(request, srv); err != nil {
			return err
		}
	} else {
		return &types.PathError{
			Path: request.GetConfigInstance().GetPath(),
		}
	}

	return nil
}
