package mongoapiredux

// MongoInterface is a mongo interface
type MongoInterface interface {
	AddToGamerCollection(gamer ...interface{}) error
	// FindOneInCollection(collection collectionhelper.ICollectionHelper, gamerName interface{}, projections []interface{}) (models.Gamer, error)
	// UpdateOneGamerByName(collection collectionhelper.ICollectionHelper, gamerName string, updateInfo interface{}) error
	// AddGameToGamerGamelist(collection collectionhelper.ICollectionHelper, listUpdate models.GamelistUpdate) error
	DeleteOneGamerFromCollectionByName(gamerName interface{}) error
	// DropCollection(collection *mongo.Collection) error
}
