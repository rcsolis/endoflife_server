package main

import (
	"log"
	"net"

	s "github.com/rcsolis/endoflife_server/internal/server"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

func main() {
	// Get a new listener on the port
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed connection on port: %v", err)
	}
	// Crete a grpc server and cycle server instance
	grpcServer := grpc.NewServer()
	// Register the TodoServer implementation
	s.RegisterGrpcServer(grpcServer)
	// Start the server
	log.Printf("Starting server: %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
