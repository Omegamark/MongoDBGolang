package mongoapi

import (
	"MongoDBGolang/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient returns a connection to an instance of Mongo
func MongoClient() *mongo.Client {
	// TODO: Make the URI here configurable.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client

}

// AddToCollection adds a new gamer to the database.
func AddToCollection(collection *mongo.Collection, gamer ...models.Gamer) {
	// TODO: Make the clientDB into an interface so these methods could be used with any DB.
	if len(gamer) > 1
	collection.InsertMany(context.TODO(), []interface{gamer})

}
