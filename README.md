# Address Microservice

## Prerequisites

- Protobuf - https://protobuf.dev/
- GRPC - https://grpc.io/docs/languages/go/quickstart/
- SQLC - https://sqlc.dev/

## Usage

Create `configs/config.toml` file. You can use `configs/config.example.toml` as example

Install dependencies:

```
make deps
```

Run server:
```
make run
```

Build binary:
```
make build
```

## Development

Update SQLC schema (`schema.sql`) and queries (`query.sql`). Then generate Go code (models, interfaces etc):
```
 make sqlc_generate
```

Update protobuf spec (`proto/address.proto`) and then generate Go code:
```
 make proto_generate
```

GRPC procedures implementation is in `internal/grpcserver/grpcserver.go`

## Tests

Run tests:
```
make test
```

