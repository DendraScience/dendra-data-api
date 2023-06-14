module github.com/dendrascience/dendra-data-api/packages/go/provider/influx

go 1.20

require (
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)

require (
	github.com/dendrascience/dendra-data-api/release/go v0.0.0
	github.com/influxdata/influxdb-client-go/v2 v2.12.2
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

replace github.com/dendrascience/dendra-data-api/release/go => ../../../../release/go
