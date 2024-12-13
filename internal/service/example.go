package service

import (
	"context"
	"dka-go-microservices/generated/example" // Ganti dengan path sesuai proyek Anda
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// Server struct mengimplementasikan layanan gRPC
type Server struct {
	example.UnimplementedExampleServiceServer
}

// Implementasi SayHello RPC
func (s *Server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	// Menerima permintaan SayHello dan memprosesnya
	message := fmt.Sprintf("Hello %s!", req.GetName())
	log.Printf("Received SayHello request: %s", req.GetName())

	// Mengembalikan respons dengan pesan
	return &example.HelloResponse{
		Message: message,
	}, nil
}

// NewServer untuk membuat dan menginisialisasi server gRPC
func NewServer() *Server {
	return &Server{}
}

// RegisterGRPCServer untuk mendaftarkan layanan ke server gRPC
func RegisterGRPCServer(grpcServer *grpc.Server, server *Server) {
	example.RegisterExampleServiceServer(grpcServer, server)
}
