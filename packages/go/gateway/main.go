/*
 *
 * Gateway reverse-proxy service
 *
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	gw "github.com/dendrascience/dendra-data-api/release/go-gw/v3"
)

type ConnectedClient[T any] struct {
	conn   *grpc.ClientConn
	client T
}

var (
	datapoints *ConnectedClient[gw.DatapointsServiceClient]
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

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// set up a connection to the datapoints server
	datapoints = dialClient[gw.DatapointsServiceClient]("DATAPOINTS_SERVICE")
	defer datapoints.conn.Close()
	datapoints.client = gw.NewDatapointsServiceClient(datapoints.conn)

	// set up server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8051"
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	// register services
	err := gw.RegisterDatapointsServiceHandlerClient(ctx, mux, datapoints.client)
	if err != nil {
		log.Fatalf("failed to register datapoints service: %v\n", err)
	}

	// SEE: https://github.com/grpc-ecosystem/grpc-gateway/issues/1889
	withCors := cors.AllowAll().Handler(mux)

	server := &http.Server{Addr: ":" + port, Handler: withCors}
	go server.ListenAndServe()

	<-ctx.Done()
	stop()

	fmt.Println("stopping server...")

	// shuts down the server with 5 seconds of grace period
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("shutdown error: %v\n", err)
	}

	fmt.Println("bye!")
}
