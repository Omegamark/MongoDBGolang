package app

import (
	appconfig "MongoDBGolang/config"
	"MongoDBGolang/handlers"
	mongoapiredux "MongoDBGolang/mongoAPIRedux"
	"fmt"
	"log"
	"net/http"
)

// StartUp starts the app
func StartUp(config *appconfig.Config) {
	// TODO: Have this return an error
	// TODO: Consider moving the Mongo hanlder struct elsewhere. Actually reconsider this whole concept, it is confusing.
	gamerHandler := &handlers.MongoHandler{DB: &mongoapiredux.MongoDatabase{DBConnection: mongoapiredux.MongoConnection{
		URI:        config.MongoCfg.MongoURI,
		Database:   config.MongoCfg.Database,
		Collection: config.MongoCfg.Collection,
	}}}

	router := handlers.InitRoutes(gamerHandler)

	fmt.Printf("Server is running at %v", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, router))
}
