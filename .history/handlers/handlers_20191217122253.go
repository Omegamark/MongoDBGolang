package handlers

import (
	"fmt"
	"net/http"
)

func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")
	err := r.ParseForm()
	fmt.Println(test)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
