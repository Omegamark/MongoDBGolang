package handlers

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	databasehelper "MongoDBGolang/mongoAPI/databaseHelper"
	"encoding/json"
	"log"
	"net/http"
)

// TODO: Have the JSON body also specify which collection to use so it can be passed to Mongo API.

// AddToCollection is a handler which writes JSON POST data to a database.
func AddToCollection(w http.ResponseWriter, r *http.Request) {
	var gamer models.Gamer
	err := json.NewDecoder(r.Body).Decode(&gamer)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Println("I'm an error 1")
	}
	collection, err := client.Collection("Gamers")
	if err != nil {
		log.Println("I'm an error 2")
	}
	err = mongoapi.AddToGamerCollection(collection, gamer)
	if err != nil {
		log.Println("handlers.go ln 33, failed to add gamer to database with error: ", err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	// fmt.Fprint(w, gamer)
	log.Println(gamer)
}

// DeleteOneFromCollection removes one or more records from a collection.
func DeleteOneFromCollection(w http.ResponseWriter, r *http.Request) {
	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Printf("Failed to initialize a client: %s", err)
	}
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

// FindOneInCollection returns a single gamer from the collection.
func FindOneInCollection(w http.ResponseWriter, r *http.Request) {
	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Printf("Failed to initialize a client: %s", err)
	}

	collection, err := client.Collection("Gamers")
	if err != nil {
		log.Printf("Failed to initialize Mongo collection: %s", err)
	}

	bodyJSON := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	name := bodyJSON["name"]
	opts := bodyJSON["opts"]

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	gamer, _ := mongoapi.FindOneInCollection(collection, name, opts.([]interface{}))

	response, err := json.Marshal(gamer)
	if err != nil {
		log.Println("Failed to marshall json. ", err)
	}

	w.Write(response)
}

// UpdateGamer is a handler which updates the info of one gamer by name.
func UpdateGamer(w http.ResponseWriter, r *http.Request) {
	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Printf("Failed to initialize a client: %s", err)
	}

	collection, err := client.Collection("Gamers")
	if err != nil {
		log.Printf("Failed to initialize Mongo collection: %s", err)
	}

	bodyJSON := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	// Asserting that this JSON value is a string.
	name := bodyJSON["name"].(string)
	// Make this into a field and value to update.
	infoToUpdate := bodyJSON["age"]

	log.Println(bodyJSON)
	mongoapi.UpdateOneGamerByName(collection, name, infoToUpdate)
}

// UpdateGamerGamelist adds a game to a gamer's gamelist.
func UpdateGamerGamelist(w http.ResponseWriter, r *http.Request) {
	var realDB databasehelper.IDatabaseHelper = &databasehelper.RealDatabaseHelper{}

	client, err := realDB.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Printf("Failed to initialize a client: %s", err)
	}

	collection, err := client.Collection("Gamers")
	if err != nil {
		log.Printf("Failed to initialize Mongo collection: %s", err)
	}

	var gamerUpdate models.GamelistUpdate

	err = json.NewDecoder(r.Body).Decode(&gamerUpdate)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	mongoapi.AddGameToGamerGamelist(collection, gamerUpdate)
}
