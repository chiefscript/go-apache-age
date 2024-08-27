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

### 1. Clone the Repository

```bash
git clone https://github.com/chiefscript/go-apache-age.git
cd go-apache-age
docker-compose up --build
