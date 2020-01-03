package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoClientSuccess(t *testing.T) {
	// var client *mongo.Client = &mongo.Client{}
	client := mongoapi.MongoClient()
	// client2 := mongo.Client{}
	// fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	// Use
	if client == (&mongo.Client{}) {
		t.Errorf("Failed to return a mongo client:\n\t Wanted: *mongo.Client\n\n \tGot: %v", reflect.TypeOf(client))
	}

}

func TestMongoClientFailure(t *testing.T) {
	client := mongoapi.MongoClient()
	if client == (string) {
		t.Error("Client is not a string")
	}
}
