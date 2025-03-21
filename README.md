# 7 Solutions Backend Challenge

This repository contains solutions to the 7 Solutions backend challenge. It includes multiple Go-based projects, each solving a specific problem or implementing a service.

### 1. Problem 1

This program fetches a JSON input from a URL and calculates the best path in a triangle of numbers.

#### Running the Program

```bash
go run 1/main.go
```

### 2. Problem 2

This program calculates the minimum sum sequence based on a given input of constraints (`L`, `R`, `=`).

#### Running the Program

```bash
go run 2/main.go
```

### 3. PieFireDire (gRPC and HTTP Service)

This project implements a gRPC and HTTP service for summarizing text by counting occurrences of specific allowed words.

## Project Structure

```
7sol-challenge/
├── PieFireDire/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   └── services/
│   │       ├── summary_service.go
│   │       └── summary_service_test.go
│   ├── proto/
|   |   └── summary.proto
│   │   └── summary/
│   │       ├── summary.pb.go
│   │       └── summary_grpc.pb.go
│   ├── go.mod
│   ├── go.sum
├── 1/
│   └── main.go
├── 2/
│   └── main.go
└── README.md
```

#### Features

- **gRPC Service**: Exposes a `GetSummary` RPC method.
- **HTTP API**: Provides a `/beef/summary` endpoint for text summarization.
- **Unit Tests**: Includes comprehensive tests for the gRPC service.

#### Running the Service

1. **Start the gRPC Server**:

   ```bash
   go run cmd/server/main.go
   ```

   The gRPC server will run on `localhost:50051`.

2. **Access the HTTP API**:
   The HTTP server will run on `localhost:8080`. Use the `/beef/summary` endpoint to summarize text:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"text": "t-bone fatback pastrami"}' http://localhost:8080/beef/summary
   ```

3. **Run Tests**:
   ```bash
   go test ./internal/services/...
   ```
