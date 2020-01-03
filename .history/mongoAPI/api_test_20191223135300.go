package mongoapi_test

import (
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoClient(t *testing.T) {
	var client *mongo.Client = &mongo.Client{}
	// client = mongoapi.MongoClient()
	fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	if client != *mongo.Client() {
		fmt.Println("Failed to obtain a Mongo Client")
	}
}
