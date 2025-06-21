package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "GRPCAgentHub/GRPCAgentHub/agent"
)

// server is used to implement agent.AgentServiceServer.
type server struct {
	pb.UnimplementedAgentServiceServer
}

// ProcessTask implements agent.AgentServiceServer
func (s *server) ProcessTask(ctx context.Context, in *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Received data from client: %v", in.GetData())
	return &pb.TaskResponse{Result: "Task processed successfully by mock server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAgentServiceServer(s, &server{})
	// Enable server reflection
	reflection.Register(s)
	log.Printf("Server with reflection listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
