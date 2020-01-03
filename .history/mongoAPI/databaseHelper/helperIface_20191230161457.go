package databasehelper

import (
	clienthelper "MongoDBGolang/mongoAPI/clientHelper"
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"
)

type IDatabaseHelper interface {
	NewClient(uri string) (clienthelper.IClientHelper, error)
	Collection(name string) (collectionhelper.ICollectionHelper, error)
}
