package mongoapi

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	NewClient() (mongo.Client, error)
}
