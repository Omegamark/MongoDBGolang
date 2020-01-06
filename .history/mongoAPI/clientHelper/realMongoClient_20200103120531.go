package clienthelper

import (
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

	"go.mongodb.org/mongo-driver/mongo"
)

// RealClientHelper is a struct which contains the ACTUAL Mongo client for production purposes.
type RealClientHelper struct {
	Client *mongo.Client
}

// Collection returns a MongoDB Collection, *currently only from the "BR" database.
func (rc RealClientHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	// TODO: Make the Database and Collection names dynamic, potentially use constants for a dropdown.
	return rc.Client.Database("BR").Collection(name), nil
}
