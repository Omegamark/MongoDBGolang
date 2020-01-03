package main

import (
	"MongoDBGolang/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/addGamer", handlers.AddToCollection).Methods("POST")
	r.HandleFunc("/findGamer", handlers.FindOneInCollection).Methods("GET")
	// r.HandleFunc("/updateGamer", handlers.UpdateGamer).Methods("PUT")
	// r.HandleFunc("/addGameToList", handlers.UpdateGamerGamelist).Methods("PUT")
	r.HandleFunc("/deleteGamer", handlers.DeleteOneFromCollection).Methods("DELETE")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
