package main

import (
	"context"
	"dka-go-microservices/generated/example"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Connect to the gRPC server on port 5051
	conn, err := grpc.Dial(":5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Create a new ExampleService client
	client := example.NewExampleServiceClient(conn)

	// Send a SayHello request
	req := &example.HelloRequest{Name: "World"}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response message
	fmt.Println(res.Message)
}
