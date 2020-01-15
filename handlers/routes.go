package handlers

import (
	mongoapiredux "MongoDBGolang/mongoAPIRedux"

	"github.com/gorilla/mux"
)

// I hate this, must be a better way.
var gamerHandler = &MongoHandler{&mongoapiredux.MongoDatabase{DBConnection: mongoapiredux.MongoConnection{
	URI:        "mongodb://localhost:27017",
	Database:   "BR",
	Collection: "Gamers",
}}}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/addGamer", gamerHandler.AddToCollection).Methods("POST")
	// r.HandleFunc("/findGamer", FindOneInCollection).Methods("GET")
	// r.HandleFunc("/updateGamer", UpdateGamer).Methods("PUT")
	// r.HandleFunc("/addGameToList", UpdateGamerGamelist).Methods("PUT")
	// r.HandleFunc("/deleteGamer", DeleteOneFromCollection).Methods("DELETE")

	return r
}
