PROJECT_BASE := $(shell pwd)
PROJECT_OUTPUT := $(PROJECT_BASE)/output
PKG_CONVERTER_PINT := packages/python/converter/pint
PKG_DATAPOINTS := packages/go/datapoints
PKG_GATEWAY := packages/go/gateway
PKG_METADATA := packages/go/metadata
PKG_PROVIDER_INFLUX := packages/go/provider/influx
PKG_TRANSFORMER_MATHJS := packages/node/transformer/mathjs

main: print-vars build-datapoints build-gateway build-metadata build-provider-influx

print-vars:
	@echo "Make vars..."
	@echo PROJECT_BASE=$(PROJECT_BASE)
	@echo PROJECT_OUTOUT=$(PROJECT_OUTPUT)
	@echo PKG_CONVERTER_PINT=$(PKG_CONVERTER_PINT)
	@echo PKG_DATAPOINTS=$(PKG_DATAPOINTS)
	@echo PKG_GATEWAY=$(PKG_GATEWAY)
	@echo PKG_METADATA=$(PKG_METADATA)
	@echo PKG_PROVIDER_INFLUX=$(PKG_PROVIDER_INFLUX)
	@echo PKG_TRANSFORMER_MATHJS=$(PKG_TRANSFORMER_MATHJS)


##
# setup
##

# IMPORTANT: Install Protocol Buffer Compiler beforehand
# https://grpc.io/docs/protoc-installation/

.PHONY: setup
setup: setup-go setup-node setup-python

# See https://grpc.io/docs/languages/go/quickstart/
# See https://grpc-ecosystem.github.io/grpc-gateway/
.PHONY: setup-go
setup-go:
	@echo "Installing required tools..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@printf "\e[32mSuccess!\e[39m\n"

# See https://github.com/grpc/grpc-node/tree/master/packages/grpc-tools
# See https://github.com/grpc/grpc/tree/v1.54.0/examples/node/static_codegen
.PHONY: setup-node
setup-node:
	@echo "Installing required tools..."
	npm install -g grpc-tools
	@printf "\e[32mSuccess!\e[39m\n"

# See https://pipenv.pypa.io/en/latest/
# See https://black.readthedocs.io/en/stable/index.html
# See https://grpc.io/docs/languages/python/quickstart/
.PHONY: setup-python
setup-python:
	@echo "Installing required tools..."
	pip3 install black
	pip3 install grpcio
	pip3 install grpcio-tools
	pip3 install pipenv --user
	@printf "\e[32mSuccess!\e[39m\n"


##
# converter-service-pint
##

.PHONY: build-converter-pint
build-converter-pint: fmt-converter-pint
	@echo "Skipping build"

.PHONY: fmt-converter-pint
fmt-converter-pint:
	@echo "Running black..."
	black $(PKG_CONVERTER_PINT)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: install-converter-pint
install-converter-pint:
	@echo "Installing dependencies..."
	(cd $(PKG_CONVERTER_PINT) && python3 -m pipenv sync)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-converter-pint
run-converter-pint: export PIPENV_PIPFILE=$(PKG_CONVERTER_PINT)/Pipfile
run-converter-pint: export PIPENV_VENV_IN_PROJECT=$(PKG_CONVERTER_PINT)
run-converter-pint: export PYTHONPATH=$(PROJECT_BASE)/release/python
run-converter-pint: build-converter-pint
	@echo "Running service..."
	python3 -m pipenv run python3 $(PKG_CONVERTER_PINT)/main.py


##
# datapoints-service
##

.PHONY: build-datapoints
build-datapoints: fmt-datapoints
	@echo "Running go build..."
	go build -o $(PROJECT_OUTPUT)/datapoints-service -C $(PKG_DATAPOINTS) main.go
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: fmt-datapoints
fmt-datapoints:
	@echo "Running gofmt..."
	gofmt -s -w $(PKG_DATAPOINTS)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-datapoints
run-datapoints: build-datapoints
	@echo "Running service..."
	$(PROJECT_OUTPUT)/datapoints-service


##
# gateway
##

.PHONY: build-gateway
build-gateway: fmt-gateway
	@echo "Running go build..."
	go build -o $(PROJECT_OUTPUT)/gateway -C $(PKG_GATEWAY) main.go
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: fmt-gateway
fmt-gateway:
	@echo "Running gofmt..."
	gofmt -s -w $(PKG_GATEWAY)
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-gateway
run-gateway: build-gateway
	@echo "Running gateway..."
	$(PROJECT_OUTPUT)/gateway


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
# transformer-service-mathjs
##

.PHONY: build-transformer-mathjs
build-transformer-mathjs: fmt-transformer-mathjs
	@echo "Running npm run build..."
	npm run --prefix $(PKG_TRANSFORMER_MATHJS) build
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: fmt-transformer-mathjs
fmt-transformer-mathjs:
	@echo "Running npm run fmt..."
	npm run --prefix $(PKG_TRANSFORMER_MATHJS) fmt
	@printf "\e[32mSuccess!\e[39m\n"

.PHONY: run-transformer-mathjs
run-transformer-mathjs: export NODE_PATH=$(PROJECT_BASE)/$(PKG_TRANSFORMER_MATHJS)/node_modules
run-transformer-mathjs: build-transformer-mathjs
	@echo "Running service..."
	npm run --prefix $(PKG_TRANSFORMER_MATHJS) start


##
# protoc
##

.PHONY: clean-protoc
clean-protoc:
	rm -rf release/go/v3
	rm -rf release/go-gw/v3
	rm -rf release/node/v3
	rm -rf release/openapiv2/v3
	rm -rf release/python/v3

.PHONY: protoc
protoc: protoc-go protoc-go-gw protoc-node protoc-python

.PHONY: protoc-go
protoc-go:
	protoc -I=proto \
		--go_out=release/go \
		--go_opt=paths=source_relative \
		--go-grpc_out=release/go \
		--go-grpc_opt=paths=source_relative \
		proto/v3/*.proto

# See https://github.com/grpc-ecosystem/grpc-gateway/blob/main/examples/internal/proto/examplepb/a_bit_of_everything.proto
# See https://grpc-ecosystem.github.io/grpc-gateway/docs/development/grpc-gateway_v2_migration_guide/
.PHONY: protoc-go-gw
protoc-go-gw:
	protoc -I=proto \
		--go_out=release/go-gw \
		--go_opt=paths=source_relative \
		--go-grpc_out=release/go-gw \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=release/go-gw \
	    --grpc-gateway_opt=logtostderr=true \
	    --grpc-gateway_opt=paths=source_relative \
	    --grpc-gateway_opt=generate_unbound_methods=true \
	    --openapiv2_out=release/openapiv2 \
	    --openapiv2_opt=logtostderr=true,json_names_for_fields=false \
		proto/v3/*.proto

# See https://github.com/AckeeCZ/grpc-server-reflection
.PHONY: protoc-node
protoc-node:
	grpc_tools_node_protoc -I=proto \
		--js_out=import_style=commonjs,binary:release/node \
		--grpc_out=grpc_js:release/node \
		proto/v3/*.proto
	grpc_tools_node_protoc -I=proto \
		--descriptor_set_out=release/node/descriptor_set.bin \
		--include_imports \
		proto/v3/transformer.proto proto/v3/types.proto

.PHONY: protoc-python
protoc-python:
	python3 -m grpc_tools.protoc -I=proto \
		--python_out=release/python \
		--pyi_out=release/python \
		--grpc_python_out=release/python \
		proto/v3/*.proto
