package mongoapi

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseHelper interface {
	Collection(name string) *mongo.Collection
	NewClient() (*mongo.Client, error)
}

type DummyData struct {
}

func (dd DummyData) NewClient(uri string) *mongo.Client {

}

func (dd DummyData) Collection(name string) (*mongo.Collection, error) {

}

type RealData struct {
}

func (rd RealData) NewClient(uri string) *mongo.Client {
	return &mongo.Client{}
}

func (rd RealData) Collection(name string) (*mongo.Collection, error) {
	return &mongo.Collection{}, nil
}
