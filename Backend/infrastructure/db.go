package infrastructure

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Collection {
	//FIND .ENV
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}
	//getting the uri from .env
	MONGO_URI := os.Getenv("MONGO_URI")

	//connceting to th database
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//checking the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	return client.Database("TheraMap").Collection("users")
}
