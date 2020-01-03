package databasehelper

import (
	clienthelper "MongoDBGolang/mongoAPI/clientHelper"
)

type IDatabaseHelper interface {
	NewClient(uri string) (clienthelper.IClientHelper, error)
	// Collection(name string) (collectionhelper.ICollectionHelper, error)
}
