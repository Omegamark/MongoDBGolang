package mongoapi

import "go.mongodb.org/mongo-driver/mongo"

type IDatabaseHelper interface {
	NewClient() (*mongo.Client, error)
	Collection(name string) (*mongo.Collection, error)
}

type DummyData struct {
}

func (dd DummyData) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}

}

func (dd DummyData) Collection(name string) (*mongo.Collection, error) {
	return &mongo.Collection{}, nil

}

type RealData struct {
}

func (rd RealData) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}
}

func (rd RealData) Collection(name string) (*mongo.Collection, error) {
	return &mongo.Collection{}, nil
}
