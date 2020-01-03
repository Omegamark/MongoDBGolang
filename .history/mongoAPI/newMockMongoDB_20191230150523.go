package mongoapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type IDatabaseHelper interface {
// 	NewClient(uri string) (*mongo.Client, error)
// 	Collection(name string) (ICollectionHelper, error)
// }

type DummyDatabaseHelper struct {
}

type RealDatabaseHelper struct {
	*mongo.Database
}

type ICollectionHelper interface {
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	InsertMany(context.Context, []interface{}, ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
}

type DummyCollectionHelper struct {
	data []interface{}
}

type RealCollectionHelper struct {
	*mongo.Collection
}

// func (dd DummyCollectionHelper) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
// 	dd.data = append(dd.data, data)
// 	return &mongo.InsertOneResult{}, nil
// }

// func (dd DummyCollectionHelper) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
// 	dd.data = append(dd.data, data)
// 	return &mongo.InsertManyResult{}, nil
// }

func (dd DummyDatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil

}

func (dd DummyDatabaseHelper) Collection(name string) (ICollectionHelper, error) {
	return &DummyCollectionHelper{}, nil
}

type RealData struct {
}

func (rd RealDatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func (rd RealDatabaseHelper) Collection(name string) (ICollectionHelper, error) {
	return RealCollectionHelper{}, nil
}
