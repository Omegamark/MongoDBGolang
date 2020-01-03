package mongoapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabaseHelper interface {
	NewClient(uri string) (*mongo.Client, error)
	Collection(name string) (*mongo.Collection, error)
}

type DummyData struct {
}

type ICollectionHelper interface {
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type DummyCollectionHelper struct {
	data []interface{}
}

type RealCollectionHelper struct {
}

func (dd DummyCollectionHelper) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &mongo.InsertOneResult{}, nil
}

func (dd DummyData) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil

}

func (dd DummyData) Collection(name string) (*mongo.Collection, error) {
	return &mongo.Collection{}, nil
}

type RealData struct {
}

func (rd RealData) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func (rd RealData) Collection(name string) (*mongo.Collection, error) {
	return &mongo.Collection{}, nil
}
