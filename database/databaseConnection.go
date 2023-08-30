package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance(ctx context.Context) (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	mongoURL := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB server: %w", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	databaseName := "super_cluster"
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
