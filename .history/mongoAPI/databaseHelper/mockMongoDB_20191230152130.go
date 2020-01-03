package databasehelper

import (
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

	"go.mongodb.org/mongo-driver/mongo"
)

type DummyDatabaseHelper struct {
}

func (dd DummyDatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func (dd DummyDatabaseHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	return &collectionhelper.DummyCollectionHelper{}, nil
}
