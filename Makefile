.PHONY: build

build: 
		go build -v ./cmd/address_service

.PHONY: run

run: 
		go run -v ./cmd/address_service

.PHONY: test
test:
		go test -v -race -timeout 30s ./...


DEFAULT_GOAL := build