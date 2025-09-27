BUILD_DIR = ./bin
BINARY_NAME = aircraft-departure-simulator
GATEWAY_PATH = ./cmd/ads-gateway/main.go
FLIGHT_PATH = ./cmd/ads-flight/main.go
MANAGE_PATH = ./cmd/ads-manage/main.go
SEED_PATH = ./cmd/ads-seed/main.go

## help: get info about the targets within this makefile
.phony: help
help:
	@echo "ads-backend usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
	@echo "ads-backend variables:"
	@echo "  BUILD_DIR: ${BUILD_DIR}"
	@echo "  BINARY_NAME: ${BINARY_NAME}"
	@echo "  GATEWAY_PATH: ${GATEWAY_PATH}"
	@echo "  FLIGHT_PATH: ${FLIGHT_PATH}"
	@echo "  MANAGE_PATH: ${MANAGE_PATH}"
	@echo "  SEED_PATH: ${SEED_PATH}"

## gateway: runs the api gateway app
.phony: gateway
gateway: generate
	@go run ${GATEWAY_PATH}

## flight: runs the flight service app
.phony: flight
flight: generate
	@go run ${FLIGHT_PATH}

## manage: runs the manage service app
.phony: manage
manage: generate
	@go run ${MANAGE_PATH}

## seed: runs the seed service app, for populating the database with mock data
.phony: seed
seed: generate
	@go run ${SEED_PATH}

## build: build the application
.phony: build
build: generate
	@mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_DIR}/${BINARY_NAME}-darwin ${GATEWAY_PATH}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME}-linux ${GATEWAY_PATH}
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe ${GATEWAY_PATH}

## generate: this application uses sqlc, so this target will generate the code using config from sqlc.yaml
.phony: generate
generate:
	@sqlc generate

.phony: confirm
confirm:
	@echo -n "are you sure? [y/n] " && read ans && [ $${ans:-n} = y ]

## clean: clean up any build artifacts in the build directory
.phony: clean
clean: confirm
	@echo "cleaning up..."
	@rm -rf ${BUILD_DIR}

## tidy: tidies up the module and the test cache
.phony: tidy
tidy:
	@go clean -testcache
	@go mod tidy

## install: installs all dependencies for the module
.phony: install
install: tidy
	@go install ./...

## update: updates all packages used within the module
.phony: update
update: tidy
	@go get -u ./...

## test: runs unit tests and gives a coverage report
.phony: test
test: generate tidy
	@go test -cover ./...

## test: generates a coverage report as coverage.out
.phony: coverage
coverage: generate tidy
	@go test -coverprofile=coverage.out ./...

## fmt: format the code
.phony: fmt
fmt:
	@gofmt -l .

## info: get info about the build environment, go version, packages, etc.
.phony: info
info:
	@go version
	@go env
	@go vet ./...
	@go list ./...
