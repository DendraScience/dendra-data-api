package types

import (
	pb "github.com/dendrascience/dendra-data-api/release/go/v3"
)

type Driver interface {
	StreamDatapoints(*pb.ProviderStreamDatapointsRequest, pb.ProviderService_StreamDatapointsServer) error
}
