package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	ConnectionString := "mongodb://localhost:27017"
	fmt.Print(ConnectionString)

	client, err := mongo.NewClient(options.Client().ApplyURI(ConnectionString))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client = DBinstance()

const DATABASE = "simple_bank"

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection = client.Database(DATABASE).Collection(collectionName)

	return collection
}
