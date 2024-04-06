.PHONY: build

build: 
		go build -v ./cmd/address_service

.PHONY: run

run: 
		go run -v ./cmd/address_service

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: proto_generate
proto_generate:
		protoc -I proto --go_out=plugins=grpc:pkg/api proto/address.proto

.PHONY: sqlc_generate
sqlc_generate:
		sqlc generate

.PHONY: deps
deps:
		go mod tidy



DEFAULT_GOAL := build