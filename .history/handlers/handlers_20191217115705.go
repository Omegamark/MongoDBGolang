package handlers

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"net/http"
)

func AddToCollectionHandler(w http.ResponseWriter, r *http.Request) {
	mongoapi.HelloWorld
}
