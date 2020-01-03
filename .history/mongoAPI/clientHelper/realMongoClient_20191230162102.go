package clienthelper

import collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

type RealClientHelper struct {
}

func (rc RealClientHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {

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
