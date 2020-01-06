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
		err = errors.New("Failed to delete")
		return &mongo.DeleteResult{}, err
	}
	return &mongo.DeleteResult{}, nil
}

func (dd MockCollectionHelper) FindOne(ctx context.Context, data interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	dd.data = append(dd.data, data)
	fmt.Println(data)
	if data == nil {
		fmt.Println("SANITY CHECK")
	}
	return &mongo.SingleResult{}
}

func (dd MockCollectionHelper) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*UpdateResult, error) {
	dd.data = append(dd.data, update)
}
