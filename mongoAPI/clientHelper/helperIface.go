package clienthelper

import collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"

type IClientHelper interface {
	Collection(name string) (collectionhelper.ICollectionHelper, error)
}
