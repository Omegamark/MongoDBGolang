package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoClient(t *testing.T) {
	// var client *mongo.Client = &mongo.Client{}
	client := mongoapi.MongoClient()
	// client2 := mongo.Client{}
	// fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	// Use
	if client != (&mongo.Client{}) {
		fmt.Println("Failed to obtain a Mongo Client")
		t.Error("Failed to return a mongo client:\n\t Wanted: *mongo.Client\n\n \tGot: %v", client)
	}
}
