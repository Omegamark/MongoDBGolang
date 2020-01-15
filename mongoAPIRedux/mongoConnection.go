package mongoapiredux

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is a struct with mongo variables
type MongoConnection struct {
	URI        string `json:"uri"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

func (db *MongoConnection) connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(db.URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetCollection returns a mongo collection client
func (db *MongoConnection) GetCollection() (*mongo.Collection, error) {
	client, err := db.connect()
	if err != nil {
		return nil, err
	}
	return client.Database(db.Database).Collection(db.Collection), nil
}
