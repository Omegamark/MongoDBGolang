package mongoapi_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMongoClient(t *testing.T) {
	client := MongoClient()
	fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	if reflect.TypeOf(client) != *Client {
		fmt.Println("Failed to obtain a Mongo Client")
	}
}
