package services

import (
	"dka-go-microservices/internal/service/core/parking/config"
	"google.golang.org/grpc"
)

func Service(grpcServer *grpc.Server) {
	// Initialize the service server
	srv := config.NewServer()
	// Register the service with the gRPC server
	config.RegisterGRPCServer(grpcServer, srv)
}
