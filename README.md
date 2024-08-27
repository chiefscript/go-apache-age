# Go project with Apache AGE, PostgreSQL, Docker, and gRPC

This project is a Go application that interacts with three PostgreSQL databases using Apache AGE for graph database functionalities. The application is containerized using Docker and communicates via gRPC.

## Project Structure

- `proto/`: Contains the gRPC `.proto` files.
- `cmd/`: Contains the main entry point for the Go application.
- `internal/`: Contains the internal logic, including database connections and gRPC server implementation.
- `Dockerfile`: Defines the Docker image for the Go application.
- `docker-compose.yml`: Defines the multi-container Docker application, including three PostgreSQL databases and the Go app.

## Requirements

- Docker & Docker Compose installed
- Go installed on your local machine

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/chiefscript/go-apache-age.git
cd go-apache-age
```

This guide covers the steps to set up and run the project.

### 2. Ensure Protobuf code is generated

Make sure you have generated the Go code from the`.proto` file. This step is meant to avoid unresolved type errors.

Run the following command from the project root:
```bash
    protoc --go_out=. --go-grpc_out=. proto/service.proto
```

    This will generate two files:
    - `proto/service.pb.go`: Contains the generated code for messages and service definitions.
    - `proto/service_grpc.pb.go`: Contains the generated code for gRPC server and client stubs.


### 3. Rebuild the project
After generating the protobuf files, rebuild the project to ensure everything is set up correctly:

```bash
go build ./...
```
This step will help resolve any import-related issues.


### 4. Run the Docker Compose setup
Run the following command to build and start the Docker containers:
```docker
docker compose up --build
```

### 5. Test the gRPC endpoints
You can use grpcurl or a gRPC client to test the endpoints:
```bash
grpcurl -plaintext -d '{"graphName": "testGraph"}' localhost:50051 proto.UserService/CreateGraph
```

Postman has added support for gRPC, and you can test your gRPC server using it.

Steps:
- In Postman, start a new request and switch the request type to gRPC.
- Enter localhost:50051 as the server address.
- Load the .proto file to define the service and methods.
- Choose a method, fill in the request data, and send the request

### 6. Stopping the Containers
To stop the running containers, use:
```bash
docker-compose down
```

### 7. Cleaning Up
To remove all containers, networks, and volumes created by Docker Compose, use:
```bash
docker-compose down --volumes
```