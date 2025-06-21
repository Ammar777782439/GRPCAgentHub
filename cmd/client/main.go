package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "GRPCAgentHub/GRPCAgentHub/agent"
)

const (
	address     = "localhost:50051"
	defaultData = "This is a test task"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAgentServiceClient(conn)

	// Contact the server and print out its response.
	data := defaultData
	if len(os.Args) > 1 {
		data = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ProcessTask(ctx, &pb.TaskRequest{Data: data})
	if err != nil {
		log.Fatalf("could not process task: %v", err)
	}
	log.Printf("Response from server: %s", r.GetResult())
}
