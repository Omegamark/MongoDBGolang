package handlers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes
func InitRoutes(gamerHandler *MongoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/addGamer", gamerHandler.AddToCollection).Methods("POST")
	// r.HandleFunc("/findGamer", FindOneInCollection).Methods("GET")
	// r.HandleFunc("/updateGamer", UpdateGamer).Methods("PUT")
	// r.HandleFunc("/addGameToList", UpdateGamerGamelist).Methods("PUT")
	r.HandleFunc("/deleteGamer", gamerHandler.DeleteOneFromCollection).Methods("DELETE")

	return r
}

func initRouterConfig() {

}
