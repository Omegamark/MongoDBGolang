package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"fmt"
	"testing"
)

func TestMongoClient(t *testing.T) {
	// var client *mongo.Client = &mongo.Client{}
	client := mongoapi.MongoClient()
	// fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	if client != mongoapi.MongoClient() {
		fmt.Println("Failed to obtain a Mongo Client")
	}
}
