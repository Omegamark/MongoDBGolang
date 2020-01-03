package handlers

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AddToCollection is a handler which writes JSON POST data to a database.
func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read Request Body: ", err)
	}
	fmt.Println(body)
	var gamer models.Gamer
	err = json.Unmarshal(body, &gamer)
	if err != nil {
		fmt.Println("handlers.go ln 24, failed to unmarshal JSON with error: ", err)
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read Request Body: ", err)
	}
	bodyJSON := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyJSON)
	fmt.Println(bodyJSON)

}

// func FindInCollection(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Sanity Check")

// }
