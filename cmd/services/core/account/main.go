package main

import (
	"dka-go-microservices/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Setup gRPC server
	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// Initialize the service server
	srv := service.NewServer()
	// Register the service with the gRPC server
	service.RegisterGRPCServer(grpcServer, srv)
	// Channel to capture termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine for serving gRPC
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("Successfully started microservices")
	// Wait for a termination signal
	<-stop
	// Gracefully stop the gRPC server
	log.Println("Shutting down gracefully...")
	grpcServer.GracefulStop()
	// Optionally, wait a little to ensure shutdown is clean
	time.Sleep(time.Second)
	log.Println("Server stopped.")
}
