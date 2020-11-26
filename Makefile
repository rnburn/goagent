.PHONY: default help test-unit docker-test bench lint deps ci-deps check-examples generate-config

default: help

help: ## Prints this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

test: test-unit test-docker

test-unit: ## Runs unit tests.
	@go test -count=1 -v -race -cover ./...

test-docker: ## Runs the docker container attribute test.
	@./tests/docker/test.sh

bench: ## Runs benchmarks.
	go test -v -run - -bench . -benchmem ./...

lint: ## Runs linters for Go code.
	@echo "Running linters..."
	@golangci-lint run ./... && echo "Done."

deps: ## Pulls the deps.
	@go get -v -t -d ./...

deps-ci: ## Pulls the CI deps.
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint

check-examples: ## Makes sure examples can be compiled.
	go build -o ./examples/http_client instrumentation/hypertrace/net/hyperhttp/examples/client/main.go && rm ./examples/http_client
	go build -o ./examples/http_server instrumentation/hypertrace/net/hyperhttp/examples/server/main.go && rm ./examples/http_server
	go build -o ./examples/grpc_client instrumentation/hypertrace/google.golang.org/hypergrpc/examples/client/main.go && rm ./examples/grpc_client
	go build -o ./examples/grpc_server instrumentation/hypertrace/google.golang.org/hypergrpc/examples/server/main.go && rm ./examples/grpc_server
	go build -o ./examples/query instrumentation/hypertrace/database/hypersql/examples/query/main.go && rm ./examples/query
	go build -o ./examples/http_client instrumentation/opentelemetry/net/hyperhttp/examples/client/main.go && rm ./examples/http_client
	go build -o ./examples/http_server instrumentation/opentelemetry/net/hyperhttp/examples/server/main.go && rm ./examples/http_server
	go build -o ./examples/grpc_client instrumentation/opentelemetry/google.golang.org/hypergrpc/examples/client/main.go && rm ./examples/grpc_client
	go build -o ./examples/grpc_server instrumentation/opentelemetry/google.golang.org/hypergrpc/examples/server/main.go && rm ./examples/grpc_server
	go build -o ./examples/http_client instrumentation/opencensus/net/hyperhttp/examples/client/main.go && rm ./examples/http_client
	go build -o ./examples/http_server instrumentation/opencensus/net/hyperhttp/examples/server/main.go && rm ./examples/http_server
	go build -o ./examples/grpc_client instrumentation/opencensus/google.golang.org/hypergrpc/examples/client/main.go && rm ./examples/grpc_client
	go build -o ./examples/grpc_server instrumentation/opencensus/google.golang.org/hypergrpc/examples/server/main.go && rm ./examples/grpc_server

generate-config: ## Generates config object for Go
	@echo "Compiling the proto file"
	cd config/agent-config; protoc --go_out=plugins=grpc,paths=source_relative:.. config.proto
	@echo "Generating the loaders"
	cd config; go run cmd/generator/main.go agent-config/config.proto
	@echo "Done."
