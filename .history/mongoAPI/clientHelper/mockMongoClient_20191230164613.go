package clienthelper

import collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

type MockClientHelper struct {
}

func (mc MockClientHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	return &collectionhelper.MockCollectionHelper{}, nil
}
