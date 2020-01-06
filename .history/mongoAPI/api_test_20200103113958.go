package mongoapi_test

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	databasehelper "MongoDBGolang/mongoAPI/databaseHelper"
	"fmt"
	"testing"
)

// func TestMongoClientSuccess(t *testing.T) {
// 	// var client *mongo.Client = &mongo.Client{}
// 	var db mongoapi.DataAccess = mongoapi.FakeMongoClient{}
// 	client, _ := db.Connect(context.TODO(), nil)
// 	// fmt.Println("I'm a test!: ", reflect.TypeOf(client))
// 	// Use () parens to encapsulate an initialized type.
// 	if client == (&mongo.Client{}) {
// 		t.Errorf("Failed to return a mongo client:\n\t Wanted: *mongo.Client\n\n \tGot: %v", reflect.TypeOf(client))
// 	}

// }

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

func TestAddToGamerCollectionSuccess(t *testing.T) {
	var games1 = []models.Game{{"Megaman", "Platformer", "NES"}, {"Daggerfall", "CRPG", "PC"}, {"Zelda64", "Adventure", "N64"}}
	var gamer1 = models.Gamer{
		"Mark",
		38,
		games1,
	}
	var fakeDB databasehelper.IDatabaseHelper = &databasehelper.MockDatabaseHelper{}
	fakeClient, _ := fakeDB.NewClient("I'm a fake client!")

	collection, _ := fakeClient.Collection("I'm a fake collection!")

	err := mongoapi.AddToGamerCollection(collection, gamer1)
	if err != nil {
		t.Errorf("Test Failed: %s", err)
	}
}

func TestAddToGamerCollectionFail(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = &databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client!")

	collection, _ := fakeClient.Collection("I'm a fake collection!")

	err := mongoapi.AddToGamerCollection(collection, nil)
	if err == nil {
		t.Errorf("Test Failed: %s", err)
	}
}

func TestDeleteOneGamerFromCollectionByNameSuccess(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	err := mongoapi.DeleteOneGamerFromCollectionByName(collection, "Jack Flack")
	if err != nil {
		t.Errorf("Test Failed: %s", err)
	}
}

func TestDeleteOneGamerFromCollectionByNameFail(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	err := mongoapi.DeleteOneGamerFromCollectionByName(collection, nil)
	if err == nil {
		t.Errorf("Test Failed: %s", err)
	}
}

func TestFindOneInCollectionSuccess(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	result, err := mongoapi.FindOneInCollection(collection, "Jack Flack", []interface{}{"name", "age"})
	fmt.Println("RESULT: ", result)
	if err != nil {
		t.Errorf("Test Failed: %v", result)
	}
}

func TestFindOneInCollectionFail(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	result, err := mongoapi.FindOneInCollection(collection, nil, []interface{}{"name", "age"})
	fmt.Println("RESULT: ", result)
	if err == nil {
		t.Errorf("Test Failed: %v", err)
	}
}

func TestUpdateOneGamerByNameSuccess(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	err := mongoapi.UpdateOneGamerByName(collection, "Oozie Nelson", "some update info")
	if err != nil {
		t.Errorf("Test Failed: %v", err)
	}
}

func TestUpdateOneGamerByNameFail(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	err := mongoapi.UpdateOneGamerByName(collection, "Bebop", nil)
	if err == nil {
		t.Errorf("Test Failed: %v", err)
	}
}

func TestAddGameToGamerGamelist(t *testing.T) {
	var fakeDB databasehelper.IDatabaseHelper = databasehelper.MockDatabaseHelper{}

	fakeClient, _ := fakeDB.NewClient("I'm a fake client mang!")

	collection, _ := fakeClient.Collection("I'm a fake collection yo!")

	models.GamelistUpdate{
		Name: "Rocksteady",
		Game: models.Game{
			Name:     "Doesn't matter",
			Genre:    "Fantastical",
			Platform: "Amiga CD32",
		},
	}

	mongoapi.AddGameToGamerGamelist(collection, "Rocksteady")
}
