package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"fmt"
	"reflect"
	"testing"
)

func TestMongoClient(t *testing.T) {
	client := mongoapi.MongoClient()
	fmt.Println("I'm a test!: ", reflect.TypeOf(client))
}
