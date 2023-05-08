/*
 *
 * Metadata service
 *
 */

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/dendrascience/dendra-data-api/packages/go/metadata/types"

	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

var (
	apiClient = http.Client{
		Timeout: time.Second * 20,
	}
	apiURL             string
	getDatastreamQuery = "?$select%5B%5D=_id" +
		"&$select%5B%5D=datapoints_config" +
		"&$select%5B%5D=datapoints_config_built" +
		"&$select%5B%5D=datapoints_config_refd" +
		"&$select%5B%5D=derivation_method" +
		"&$select%5B%5D=derived_from_datastream_ids" +
		"&$select%5B%5D=name" +
		"&$select%5B%5D=station_lookup" +
		"&$select%5B%5D=terms" +
		"&$select%5B%5D=terms_info" +
		"&$select%5B%5D=version_id"
	unmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	apiURL = os.Getenv("WEB_API_URL")
	if apiURL == "" {
		apiURL = "https://api.dendra.science/v2"
	}

	fmt.Printf("using web api at %s\n", apiURL)

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

	pb.RegisterMetadataServiceServer(s, &server{})
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
	pb.UnimplementedMetadataServiceServer
}

func (s *server) GetDatastream(ctx context.Context, request *pb.GetDatastreamRequest) (*pb.GetDatastreamResponse, error) {
	fmt.Println("get datastream request received")

	if request.GetDatastreamId() == "" {
		return nil, types.ErrDatastreamIdEmpty
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL+
		"/datastreams/"+
		request.GetDatastreamId()+getDatastreamQuery,
		nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}
	// TODO: Handle auth
	// req.Header.Set("Authorization", "Bearer "+apiToken)

	// send request
	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-success status code %d", resp.StatusCode)
	}

	// decode response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read all error: %w", err)
	}
	datastream := pb.Datastream{}
	if err := unmarshalOptions.Unmarshal(respBody, &datastream); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &pb.GetDatastreamResponse{
		Datastream: &datastream,
	}, nil
}

func (s *server) ListUoms(ctx context.Context, _ *emptypb.Empty) (*pb.ListUomsResponse, error) {
	fmt.Println("list uoms request received")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL+
		"/uoms",
		nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}

	// send request
	resp, err := apiClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned non-success status code %d", resp.StatusCode)
	}

	// decode response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read all error: %w", err)
	}
	listResp := pb.ListUomsResponse{}
	if err := unmarshalOptions.Unmarshal(respBody, &listResp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &listResp, nil
}
