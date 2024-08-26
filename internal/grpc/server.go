package grpc

import (
	"context"
	"log"
	"net"

	pb "go-apache-age/go-apache-go/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	// You can add references to your databases here
}

func (s *server) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// Implement CreateUser logic
	// Placeholder response
	return &pb.UserResponse{
		Id:    1, // Example ID
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.UserId) (*pb.UserResponse, error) {
	// Implement GetUser logic
	// Placeholder response
	return &pb.UserResponse{
		Id:    req.Id,
		Name:  "John Doe",     // Example name
		Email: "john@doe.com", // Example email
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UserRequest) (*pb.Empty, error) {
	// Implement UpdateUser logic
	// Placeholder response
	return &pb.Empty{}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	// Implement DeleteUser logic
	// Placeholder response
	return &pb.Empty{}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
