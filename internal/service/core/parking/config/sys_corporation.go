package config

import (
	"context"
	syscorporation "dka-go-microservices/generated/services/parking/config"
	database "dka-go-microservices/internal/database/MongoDB"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net/http"
)

// Server struct implements the gRPC service
type Server struct {
	syscorporation.UnimplementedUserLoginServiceServer
}

// GetAllData Implementasi SayHello RPC
func (s *Server) GetAllData(ctx context.Context, req *syscorporation.GetAllDataRequest) (*syscorporation.GetAllDataResponse, error) {
	// Log the request origin for debugging
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("Request: %s -> %s", p.Addr, p.LocalAddr)
	}

	// Connect to database
	db, err := database.Client(ctx).GetDatabase("dka_parking")
	if err != nil {
		log.Println("Database connection error:", err)
		return &syscorporation.GetAllDataResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed To Get Data",
			Error:  err.Error(),
		}, nil
	}

	// Find documents in the collection
	cursor, err := db.Collection("sys_corporation").Find(ctx, bson.D{})
	if err != nil {
		log.Println("Find query error:", err)
		return &syscorporation.GetAllDataResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to fetch data from database",
			Error:  err.Error(),
		}, nil
	}

	// Ensure cursor is closed when function returns
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var results []*syscorporation.ModelData
	// Use cursor.All to decode all documents at once
	if err := cursor.All(ctx, &results); err != nil { // Perhatikan penggunaan '&' di sini
		log.Println("Error decoding cursor:", err)
		return &syscorporation.GetAllDataResponse{
			Status: false,
			Code:   http.StatusUnprocessableEntity,
			Msg:    "Failed to decode data",
			Error:  err.Error(),
		}, nil
	}

	// Log and return response with fetched data
	log.Printf("Fetched data: %v", results)
	return &syscorporation.GetAllDataResponse{
		Status: true,
		Code:   http.StatusOK,
		Msg:    "Data fetched successfully",
		Data:   results,
	}, nil
}

// NewServer creates and initializes a new gRPC server
func NewServer() *Server {
	return &Server{}
}

// RegisterGRPCServer registers the gRPC service with the server
func RegisterGRPCServer(grpcServer *grpc.Server, server *Server) {
	syscorporation.RegisterUserLoginServiceServer(grpcServer, server)
}
