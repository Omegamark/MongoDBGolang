package handlers

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	databasehelper "MongoDBGolang/mongoAPI/databaseHelper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// AddToCollection is a handler which writes JSON POST data to a database.
func AddToCollection(w http.ResponseWriter, r *http.Request) {
	var gamer models.Gamer
	err := json.NewDecoder(r.Body).Decode(&gamer)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}

	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("I'm an error 1")
	}
	collection, err := client.Collection("Gamers")
	if err != nil {
		fmt.Println("I'm an error 2")
	}
	err = mongoapi.AddToGamerCollection(collection, gamer)
	if err != nil {
		fmt.Println("handlers.go ln 33, failed to add gamer to database with error: ", err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	// fmt.Fprint(w, gamer)
	fmt.Println(gamer)
}

// DeleteOneFromCollection removes one or more records from a collection.
func DeleteOneFromCollection(w http.ResponseWriter, r *http.Request) {
	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	gamerName := bodyJSON["name"]

	collection, err := client.Collection("Gamers")
	if err != nil {
		log.Printf("Failed to initialize Mongo collection: %s", err)
	}
	err = mongoapi.DeleteOneGamerFromCollectionByName(collection, gamerName)
	if err != nil {
		log.Printf("Failed to delete from Mongo: %s", err)
	}
}

// // FindOneInCollection returns a single gamer from the collection.
func FindOneInCollection(w http.ResponseWriter, r *http.Request) {
	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}

	name := bodyJSON["name"]
	opts := bodyJSON["opts"]

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	mongoapi.FindOneInCollection("BR", "Gamers", name, opts.([]interface{}))
}

// // UpdateGamer is a handler which updates the info of one gamer by name.
// func UpdateGamer(w http.ResponseWriter, r *http.Request) {
// 	bodyJSON := make(map[string]interface{})
// 	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
// 	if err != nil {
// 		fmt.Println("Failed to read request Body: ", err)
// 	}

// 	// Asserting that this JSON value is a string.
// 	name := bodyJSON["name"].(string)
// 	// Make this into a field and value to update.
// 	infoToUpdate := bodyJSON["age"]

// 	fmt.Println(bodyJSON)
// 	mongoapi.UpdateOneGamerByName("BR", "Gamers", name, infoToUpdate)
// }

// // UpdateGamerGamelist adds a game to a gamer's gamelist.
// func UpdateGamerGamelist(w http.ResponseWriter, r *http.Request) {
// 	var gamerUpdate models.GamelistUpdate

// 	err := json.NewDecoder(r.Body).Decode(&gamerUpdate)
// 	if err != nil {
// 		fmt.Println("Failed to read request Body: ", err)
// 	}

// 	mongoapi.AddGameToGamerGamelist("BR", "Gamers", gamerUpdate)
// }
