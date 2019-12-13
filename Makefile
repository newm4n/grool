GO111MODULE=on

.PHONY: all test clean build docker

build:
	export GO111MODULE on; \
	go build ./...

test: build
	go test ./... -v -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out