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
	// bodyJSON := make(map[string]interface{})
	var bodyJSON models.Gamer
	err = json.Unmarshal(body, &bodyJSON)
	client := mongoapi.AddToGamerCollection()
	fmt.Println(bodyJSON)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
