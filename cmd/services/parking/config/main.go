package main

import (
	"dka-go-microservices/cmd/services/parking/config/services"
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
	// Create New Server GPRC
	grpcServer := grpc.NewServer()
	go func() {
		// Adding Registration Services
		services.Service(grpcServer)
	}()
	// Channel to capture termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// Goroutine for serving gRPC
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("SERVER: Successfully started microservices")
	// Wait for a termination signal
	<-stop
	// Gracefully stop the gRPC server
	log.Println("SERVER: Shutting down gracefully...")
	grpcServer.GracefulStop()
	// Optionally, wait a little to ensure shutdown is clean
	time.Sleep(time.Second)
	log.Println("SERVER: Stopped.")
}
