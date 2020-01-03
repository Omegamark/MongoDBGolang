package databasehelper

import "go.mongodb.org/mongo-driver/mongo"

type IDatabaseHelper interface {
	NewClient(uri string) (*mongo.Client, error)
	Collection(name string) (ICollectionHelper, error)
}
