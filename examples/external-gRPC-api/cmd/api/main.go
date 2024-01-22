package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yourusername/user-service/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Implement logic to create user
	user := req.GetUser()
	log.Printf("Received: %v", user.GetName())
	return &pb.CreateUserResponse{User: user}, nil
}

// Implement other CRUD methods (GetUser, UpdateUser, DeleteUser)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
