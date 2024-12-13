package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ConfigAuth struct {
	Username string
	Password string
}

// MongoDBConfig contains MongoDB settings
type Config struct {
	URI      *string
	Database string
	Port     int
	Auth     *ConfigAuth
}

var mongoClient *mongo.Client

// MongoDB connects with the given configuration
func MongoDB(config Config) *mongo.Client {
	// Set up client options and authentication
	clientOptions := options.Client()

	if config.URI != nil {
		clientOptions.ApplyURI(*config.URI)
	}
	if config.Auth != nil {
		clientOptions.SetAuth(options.Credential{
			Username: config.Auth.Username,
			Password: config.Auth.Password,
		})
	}
	// Set timeout (optional)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	// Check connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping MongoDB:", err)
	}
	log.Println("Connected to MongoDB!")
	mongoClient = client
	return mongoClient
}
