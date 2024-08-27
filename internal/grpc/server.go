package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "go-apache-age/go-apache-go/proto"
	"go-apache-age/internal/db"
)

type server struct {
	pb.UnimplementedUserServiceServer
	DB *db.DB
}

func (s *server) CreateGraph(ctx context.Context, req *pb.CreateGraphRequest) (*pb.CreateGraphResponse, error) {
	err := s.DB.CreateGraph(req.GraphName)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGraphResponse{Message: "Graph created successfully"}, nil
}

func (s *server) AddVertex(ctx context.Context, req *pb.AddVertexRequest) (*pb.AddVertexResponse, error) {
	err := s.DB.AddVertex(req.GraphName, req.Label, req.Properties)
	if err != nil {
		return nil, err
	}
	return &pb.AddVertexResponse{Message: "Vertex added successfully"}, nil
}

func (s *server) AddEdge(ctx context.Context, req *pb.AddEdgeRequest) (*pb.AddEdgeResponse, error) {
	err := s.DB.AddEdge(req.GraphName, req.Label, req.Properties, req.FromVertexId, req.ToVertexId)
	if err != nil {
		return nil, err
	}
	return &pb.AddEdgeResponse{Message: "Edge added successfully"}, nil
}

func StartGRPCServer(db *db.DB) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{DB: db})
	log.Printf("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
