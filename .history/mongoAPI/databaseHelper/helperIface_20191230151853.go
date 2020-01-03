package databasehelper

import (
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

	"go.mongodb.org/mongo-driver/mongo"
)

type IDatabaseHelper interface {
	NewClient(uri string) (*mongo.Client, error)
	Collection(name string) (collectionhelper.ICollectionHelper, error)
}
