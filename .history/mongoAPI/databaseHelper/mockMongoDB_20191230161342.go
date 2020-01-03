package databasehelper

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DummyDatabaseHelper struct {
}

func (dd DummyDatabaseHelper) NewClient(uri string) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}
