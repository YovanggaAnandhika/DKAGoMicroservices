package database

import (
	"context"
	"dka-go-microservices/internal/connection/MongoDB"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	ctx context.Context
}

func (db DB) GetDatabase(dbName string) (*mongo.Database, error) {
	// Add Default error nil
	var err error = nil
	// MongoDB configuration
	config := MongoDB.Config{
		URI:      "mongodb://127.0.0.1:27017/",
		Database: "dka_parking",
		Auth: &MongoDB.ConfigAuth{
			Username: "developer",
			Password: "Cyberhack2010",
		},
	}
	// Get MongoDB client
	client, err := MongoDB.MongoDB(config)
	// Ensure the client disconnects on function exit
	dbInstance := client.Database(dbName)
	// Return the database instance and nil error
	return dbInstance, err
}

func Client(ctx context.Context) DB {
	return DB{ctx: ctx}
}
