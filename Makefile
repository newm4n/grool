GO111MODULE=on

.PHONY: all test clean build docker

build:
	export GO111MODULE on; \
	go build ./...

test: build
	go test ./... -v -cover