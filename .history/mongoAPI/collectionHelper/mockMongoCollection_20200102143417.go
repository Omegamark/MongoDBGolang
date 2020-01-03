package collectionhelper

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollectionHelper struct {
	data []interface{}
}

func (dd MockCollectionHelper) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	dd.data = append(dd.data, data)
	return &mongo.InsertOneResult{}, nil
}

func (dd MockCollectionHelper) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	dd.data = append(dd.data, data)
	return &mongo.InsertManyResult{}, nil
}

func (dd MockCollectionHelper) DeleteOne(ctx context.Context, data interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dd.data = append(dd.data, data)
	fmt.Println(data)
	var err error
	if data == nil {
		fmt.Println("SANITY CHECK!!!!!!")
		err = errors.New("Failed to delete")
		return &mongo.DeleteResult{}, err
	}
	return &mongo.DeleteResult{}, nil
}
