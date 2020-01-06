package databasehelper

import (
	clienthelper "MongoDBGolang/mongoAPI/clientHelper"
)

// IDatabaseHelper is an interface to which makes common real and mock instances of MongoDB allowing for unit tests without using real data.
type IDatabaseHelper interface {
	NewClient(uri string) (clienthelper.IClientHelper, error)
}
