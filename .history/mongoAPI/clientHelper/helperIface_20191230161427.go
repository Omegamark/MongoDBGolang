package clienthelper

import collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

type IClientHelper interface {
	Collection(name string) (collectionhelper.ICollectionHelper, error)
}

func (dd DummyDatabaseHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	return &collectionhelper.DummyCollectionHelper{}, nil
}
