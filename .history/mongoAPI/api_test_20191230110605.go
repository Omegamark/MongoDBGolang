package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"context"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoClientSuccess(t *testing.T) {
	// var client *mongo.Client = &mongo.Client{}
	var db mongoapi.DataAccess = mongoapi.FakeMongoClient{}
	client, _ := db.Connect(context.TODO(), nil)
	// fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	// Use () parens to encapsulate an initialized type.
	if client == (&mongo.Client{}) {
		t.Errorf("Failed to return a mongo client:\n\t Wanted: *mongo.Client\n\n \tGot: %v", reflect.TypeOf(client))
	}

}

// func TestMongoClientFailure(t *testing.T) {
// 	// This test does not actuall fail.
// 	client := mongoapi.MongoClient()
// 	if client != (&mongo.Client{}) {
// 		t.Log("Derrrrrr")
// 		// t.Error("Client is not a string")
// 	}
// }

// func TestAddGamerToCollection(t *testing.T) {
// 	client, err := mongoapi.NewFakeClient()
// 	if err != nil {
// 		t.Error("Failed to initialize mock mongo")
// 	}

// 	fakeCollection := client.Database("BR").Collection("Gamers")

// }

func TestAddToGamerCollection(t *testing.T) {
	db := mongoapi.FakeMongoClient{}
	client := mongoapi.AddToGamerCollection(db.Connect(context.TODO(), nil))

}
