.PHONY: build test lint

build:
	@go build -v ./...

test:
	@go test -v ./...

lint:
	@golangci-lint run --config .golangci.yml
