package mongoapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DummyCollectionHelper struct {
	data []interface{}
}

func (dd DummyCollectionHelper) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	dd.data = append(dd.data, data)
	return &mongo.InsertOneResult{}, nil
}

func (dd DummyCollectionHelper) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	dd.data = append(dd.data, data)
	return &mongo.InsertManyResult{}, nil
}
