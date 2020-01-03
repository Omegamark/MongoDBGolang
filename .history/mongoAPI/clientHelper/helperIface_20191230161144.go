package clienthelper

import "go.mongodb.org/mongo-driver/mongo"

type IClientHelper interface {
	NewClient() *mongo.Client
}
