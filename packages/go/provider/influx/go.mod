module github.com/dendrascience/dendra-data-api/packages/go/provider/influx

go 1.20

require (
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.12.2 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

require (
	github.com/dendrascience/dendra-data-api/release/go v0.0.0
	google.golang.org/grpc v1.53.0
)

replace github.com/dendrascience/dendra-data-api/release/go => ../../../../release/go
