PROJECT_BASE := $(shell pwd)
PROJECT_OUTPUT := $(PROJECT_BASE)/output
# PROJECT_NAME := $(shell basename $(PROJECT_BASE))
# PROJECT_TARGET := $(PROJECT_BIN)/$(PROJECT_NAME)
# GO_BIN := $(GOPATH)/bin
# GO_BIN_TARGET := $(GO_BIN)/$(PROJECT_NAME)
# SECONDS := $(shell date +%s)
# BUILD_NUMBER := $(SECONDS)
# BUILD_TIME := $(shell date -r $(SECONDS) -u +%FT%TZ)
# GIT_REVISION := $(shell git rev-parse --short HEAD 2> /dev/null)
# GIT_TAG := $(shell git describe --tags 2> /dev/null)
# VERSION := $(if $(GIT_TAG),$(GIT_TAG),v0)
# LDFLAGS := -ldflags "-X=main.BuildNumber=$(BUILD_NUMBER) -X=main.BuildTime=$(BUILD_TIME) -X=main.GitRevision=$(GIT_REVISION) -X=main.GitTag=$(GIT_TAG) -X=main.ProjectName=$(PROJECT_NAME) -X=main.Version=$(VERSION)"

PKG_METADATA := packages/go/metadata
PKG_PROVIDER_INFLUX := packages/go/provider/influx

main: print-vars build-metadata build-provider-influx

print-vars:
	@echo "Make vars..."
	@echo PROJECT_BASE=$(PROJECT_BASE)
	@echo PROJECT_OUTOUT=$(PROJECT_OUTPUT)
	@echo PKG_METADATA=$(PKG_METADATA)
	@echo PKG_PROVIDER_INFLUX=$(PKG_PROVIDER_INFLUX)


##
# metadata-service
##

.PHONY: build-metadata
build-metadata: fmt-metadata
	@echo "Running go build..."
	go build -o $(PROJECT_OUTPUT)/metadata-service -C $(PKG_METADATA) main.go
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: fmt-metadata
fmt-metadata:
	@echo "Running gofmt..."
	gofmt -s -w $(PKG_METADATA)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-metadata
run-metadata: build-metadata
	@echo "Running service..."
	$(PROJECT_OUTPUT)/metadata-service


##
# provider-service-influx
##

.PHONY: build-provider-influx
build-provider-influx: fmt-provider-influx
	@echo "Running go build..."
	go build -o $(PROJECT_OUTPUT)/provider-service-influx -C $(PKG_PROVIDER_INFLUX) main.go
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: fmt-provider-influx
fmt-provider-influx:
	@echo "Running gofmt..."
	gofmt -s -w $(PKG_PROVIDER_INFLUX)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-provider-influx
run-provider-influx: build-provider-influx
	@echo "Running service..."
	$(PROJECT_OUTPUT)/provider-service-influx


##
# protoc
##

.PHONY: clean-protoc
clean-protoc:
	rm -rf release/go/v3
	rm -rf release/node/v3
	rm -rf release/python/v3

.PHONY: protoc
protoc: protoc-go protoc-node protoc-python

.PHONY: protoc-go
protoc-go:
	protoc -I=proto \
		--go_out=release/go \
		--go_opt=paths=source_relative \
		--go-grpc_out=release/go \
		--go-grpc_opt=paths=source_relative \
		proto/v3/*.proto

# Requires https://github.com/grpc/grpc-node/tree/master/packages/grpc-tools
# See also https://github.com/grpc/grpc/tree/v1.54.0/examples/node/static_codegen
.PHONY: protoc-node
protoc-node:
	grpc_tools_node_protoc -I=proto \
		--js_out=import_style=commonjs,binary:release/node \
		--grpc_out=grpc_js:release/node \
		proto/v3/*.proto

# See https://grpc.io/docs/languages/python/quickstart/
.PHONY: protoc-python
protoc-python:
	python3 -m grpc_tools.protoc -I=proto \
		--python_out=release/python \
		--pyi_out=release/python \
		--grpc_python_out=release/python \
		proto/v3/*.proto