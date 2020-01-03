package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")
	test, err := json.Unmarshal(r.Body)
	if err != nil {
		fmt.Println("Failed: ", err)
	}
	fmt.Println(test)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
