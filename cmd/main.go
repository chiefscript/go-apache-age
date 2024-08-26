package main

import (
	"go-apache-age/internal/db"
	"go-apache-age/internal/grpc"
	"log"
)

func main() {
	// Connect to the three databases
	db1, err := db.ConnectDB("postgres1", 5432, "user", "password", "db1")
	if err != nil {
		log.Fatalf("Error connecting to db1: %v", err)
	}
	defer db1.Close()

	db2, err := db.ConnectDB("postgres2", 5467, "user", "password", "db2")
	if err != nil {
		log.Fatalf("Error connecting to db2: %v", err)
	}
	defer db2.Close()

	db3, err := db.ConnectDB("postgres3", 5478, "user", "password", "db3")
	if err != nil {
		log.Fatalf("Error connecting to db3: %v", err)
	}
	defer db3.Close()

	// Start the gRPC server
	grpc.StartGRPCServer()
}
