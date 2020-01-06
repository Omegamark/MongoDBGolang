package clienthelper

import (
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

	"go.mongodb.org/mongo-driver/mongo"
)

type RealClientHelper struct {
	Client *mongo.Client
}

// Collection returns a MongoDB Collection, *currently only from the "BR" database.
func (rc RealClientHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	// TODO: Make the Database and Collection names dynamic, potentially use constants for a dropdown.
	return rc.Client.Database("BR").Collection(name), nil
}

// func (rc RealClientHelper) NewClient() {
// 	// TODO : Make the URI here configurable.
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	return client
// }
