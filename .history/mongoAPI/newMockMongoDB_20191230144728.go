package mongoapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabaseHelper interface {
	NewClient(uri string) (*mongo.Client, error)
	Collection(name string) (ICollectionHelper, error)
}

type DummyDatabaseHelper struct {
}

type RealDatabaseHelper struct {
	*mongo.Database
}

type ICollectionHelper interface {
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type DummyCollectionHelper struct {
	data []interface{}
}

type RealCollectionHelper struct {
	*mongo.Collection
}

func (dd DummyCollectionHelper) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	dd.data = append(dd.data, data)
	return &mongo.InsertOneResult{}, nil
}

func (dd DummyCollectionHelper) InsertMany(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertManyResult, error)

func (dd DummyDatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil

}

func (dd DummyDatabaseHelper) Collection(name string) (ICollectionHelper, error) {
	return &DummyCollectionHelper{}, nil
}

type RealData struct {
}

func (rd RealData) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func (rd RealData) Collection(name string) (ICollectionHelper, error) {
	return RealCollectionHelper{}, nil
}
