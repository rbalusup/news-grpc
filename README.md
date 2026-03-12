# news-grpc

A gRPC microservice for managing news articles, built with Go and Protocol Buffers.

## Overview

`news-grpc` exposes a `NewsService` with two RPCs:

| RPC | Request | Response | Description |
|-----|---------|----------|-------------|
| `Create` | `NewRequest` | `NewsResponse` | Create a new news article |
| `Get` | `NewsID` | `NewsResponse` | Retrieve a news article by ID |

The server listens on `localhost:50051` and includes a standard gRPC health check endpoint.

## Prerequisites

- **Go 1.25+** — [Download](https://go.dev/dl/)
- **Buf CLI** (optional, for regenerating proto code) — installed automatically via `go tool buf`

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/rbalusup/news-grpc.git
cd news-grpc
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Run the server

```bash
go run cmd/server/main.go
```

The server starts and listens on `tcp://localhost:50051`.

### 4. Build a binary

```bash
go build -o news-grpc ./cmd/server
./news-grpc
```

## Protocol Buffers

Proto definitions live in `proto/news/v1/`. Generated Go code is in `api/news/v1/` and is committed to the repository — you do **not** need to regenerate it to run the server.

To regenerate after editing `.proto` files:

```bash
make generate-proto
```

This runs `go tool buf generate --template buf.gen.yaml` and outputs Go files to `./api`.

## Project Structure

```
news-grpc/
├── cmd/server/          # Server entry point (main.go)
├── proto/news/v1/       # Protocol Buffer source definitions
│   ├── news.proto       # Message types (NewRequest, NewsResponse, NewsID)
│   └── service.proto    # Service definition (NewsService)
├── api/news/v1/         # Generated Go code from proto files
├── internal/grpc/       # gRPC server implementation
│   └── server.go        # NewsService handler
├── buf.yaml             # Buf linting and breaking-change config
├── buf.gen.yaml         # Buf code generation config
├── Makefile             # Build targets
└── go.mod               # Go module definition
```

## Testing with grpcurl

You can interact with the running server using [`grpcurl`](https://github.com/fullstorydev/grpcurl):

```bash
# List available services
grpcurl -plaintext localhost:50051 list

# Health check
grpcurl -plaintext localhost:50051 grpc.health.v1.Health/Check

# Create a news article
grpcurl -plaintext -d '{
  "id": "1",
  "author": "Jane Doe",
  "title": "Breaking News",
  "summary": "A brief summary",
  "content": "Full article content here.",
  "source": "example.com",
  "tags": ["tech", "go"]
}' localhost:50051 news.v1.NewsService/Create

# Get a news article by ID
grpcurl -plaintext -d '{"id": "1"}' localhost:50051 news.v1.NewsService/Get
```

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make generate-proto` | Regenerate Go code from `.proto` files using Buf |
