package handlers

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// AddToCollection is a handler which writes JSON POST data to a database.
func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")

	var gamer models.Gamer
	err := json.NewDecoder(r.Body).Decode(&gamer)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}

	collection := mongoapi.MongoClient().Database("BR").Collection("Gamers")
	err = mongoapi.AddToGamerCollection(collection, gamer)
	if err != nil {
		fmt.Println("handlers.go ln 29, failed to add gamer to database with error: ", err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, gamer)
	fmt.Println(gamer)
}

// DeleteOneFromCollection removes one or more records from a collection.
func DeleteOneFromCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Route Hit")

	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}
	gamerName := bodyJSON["name"]
	collection := mongoapi.MongoClient().Database("BR").Collection("Gamers")
	err = mongoapi.DeleteOneGamerFromCollectionByName(collection, gamerName)
}

// FindOneInCollection returns a single gamer from the collection.
func FindOneInCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check")
	// := ioutil.ReadAll(r.Body)
	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		fmt.Println("Failed to read request Body: ", err)
	}

	fmt.Println(bodyJSON)
	name := bodyJSON["name"]
	opts := bodyJSON["opts"]
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(opts)
	fmt.Println(reflect.TypeOf(opts))

	// w.WriteHeader(200)
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Fprint(w, bodyJSON)

}
