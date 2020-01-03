package handlers

import (
	"fmt"
	"net/http"
)

func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}
