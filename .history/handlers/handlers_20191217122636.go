package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func AddToCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sanity Check!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Failed to read Request Body: ", err)
	}
	fmt.Println(body)
	fmt.Println(reflect.TypeOf(body))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
