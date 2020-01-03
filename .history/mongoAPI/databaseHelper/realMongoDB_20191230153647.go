package databasehelper

import (
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHelper struct {
}

func (dd DatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func (dd DatabaseHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	return &collectionhelper.DummyCollectionHelper{}, nil
}

func NewClient() *mongo.Client {
	// TODO : Make the URI here configurable.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
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
