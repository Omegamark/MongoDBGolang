package app

import (
	"MongoDBGolang/handlers"
	"fmt"
	"log"
	"net/http"
)

// StartUp starts the app
func StartUp() {
	// TODO: Have this return an error
	router := handlers.InitRoutes()

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
