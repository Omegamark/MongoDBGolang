package databasehelper

import (
	clienthelper "MongoDBGolang/mongoAPI/clientHelper"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RealDatabaseHelper is a struct whose NewClient method will return a real instance of MongoDB.
type RealDatabaseHelper struct {
}

// NewClient here returns a real instance of a MongoDB Client.
func (dd RealDatabaseHelper) NewClient(uri string) (clienthelper.IClientHelper, error) {
	var clientHelper clienthelper.RealClientHelper = clienthelper.RealClientHelper{}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return clienthelper.RealClientHelper{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return clienthelper.RealClientHelper{}, err
	}

	fmt.Println("Connected to MongoDB!")

	clientHelper.Client = client

	return clientHelper, nil
}
