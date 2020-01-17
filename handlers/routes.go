package handlers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes
func InitRoutes(gamerHandler *MongoHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/addGamer", gamerHandler.AddToCollection).Methods("POST")
	r.HandleFunc("/findGamer", gamerHandler.FindOneInCollection).Methods("GET")
	r.HandleFunc("/updateGamer", gamerHandler.UpdateGamer).Methods("PUT")
	r.HandleFunc("/addGameToList", gamerHandler.UpdateGamerGamelist).Methods("PUT")
	r.HandleFunc("/deleteGamer", gamerHandler.DeleteOneFromCollection).Methods("DELETE")

	return r
}

func initRouterConfig() {

}
