package handlers

import (
	"MongoDBGolang/models"
	mongoapiredux "MongoDBGolang/mongoAPIRedux"
	"encoding/json"
	"log"
	"net/http"
)

// MongoHandler is a mongo handler
type MongoHandler struct {
	DB mongoapiredux.MongoInterface
}

// AddToCollection is a handler which writes JSON POST data to a database.
func (c *MongoHandler) AddToCollection(w http.ResponseWriter, r *http.Request) {
	// NOTE: Close the request body to avoid memory leaks.
	defer r.Body.Close()

	var gamer *models.Gamer
	err := json.NewDecoder(r.Body).Decode(&gamer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.DB.AddToGamerCollection(gamer)

	response, err := json.Marshal(gamer)
	if err != nil {
		// TODO: Add a response helper utility.
		// TODO: This is returning as text, make it into JSON key:value.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	successResponseHelper(response, w)
}

// DeleteOneFromCollection removes one or more records from a collection.
func (c *MongoHandler) DeleteOneFromCollection(w http.ResponseWriter, r *http.Request) {
	// NOTE: This is how one can decode JSON with an unknown shape into go. It is difficult to work with and will really only show a reproduction of the JSON object.
	// However some rudamentory operations can be done with it as demonstrated below.
	defer r.Body.Close()

	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gamerName := bodyJSON["name"]

	err = c.DB.DeleteOneGamerFromCollectionByName(gamerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// FindOneInCollection returns a single gamer from the collection.
func (c *MongoHandler) FindOneInCollection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// NOTE: This is to show how to decode JSON if there is no struct to unmarshal into.
	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := bodyJSON["name"]
	opts := bodyJSON["opts"]

	gamer, _ := c.DB.FindOneInCollection(name, opts.([]interface{}))

	response, err := json.Marshal(gamer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	successResponseHelper(response, w)
}

// UpdateGamer is a handler which updates the info of one gamer by name.
func (c *MongoHandler) UpdateGamer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyJSON := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&bodyJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// NOTE: Asserting that this JSON value is a string.
	// NOTE: If unmarshalling into a struct, these kinds of tricks are unnecessary.
	name := bodyJSON["name"].(string)
	// NOTE: Make this into a field and value to update.
	infoToUpdate := bodyJSON["age"]

	c.DB.UpdateOneGamerByName(name, infoToUpdate)
}

// UpdateGamerGamelist adds a game to a gamer's gamelist.
func (c *MongoHandler) UpdateGamerGamelist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Having this extra type is unnecessary. The Gamer type would work just as well for this update.
	var gamerUpdate models.GamelistUpdate

	err := json.NewDecoder(r.Body).Decode(&gamerUpdate)
	if err != nil {
		log.Println("Failed to read request Body: ", err)
	}

	c.DB.AddGameToGamerGamelist(gamerUpdate)
}

// NOTE: Helper functions
func successResponseHelper(responseBytes []byte, w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}
