type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}