package handlers

import (
	"fmt"
	"net/http"
)

func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	fmt.Println(r.FormValue("name"))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
