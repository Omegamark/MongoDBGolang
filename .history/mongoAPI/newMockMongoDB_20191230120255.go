package mongoapi

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseHelper interface {
	Collection(name string) mongo.Collection
	NewClient() (mongo.Client, error)
}

type DummyData struct {
}

func (dd DummyData) NewClient() mongo.Collection {

}

func (dd DummyData) Collection() (mongo.Client, error) {

}
