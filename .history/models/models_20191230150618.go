package models

import (
	mongoapi "MongoDBGolang/mongoAPI"

	"go.mongodb.org/mongo-driver/mongo"
)

// Gamer is an struct defining an individual user.
type Gamer struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gamelist []Game `json:"gamelist"`
}

// Game is a struct defining a single game.
type Game struct {
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Platform string `json:"platform"`
}

// GamelistUpdate is a struct defining what is to be updated.
type GamelistUpdate struct {
	Name string `json:"name"`
	Game Game   `json:"game"`
}

// IDatabaseHelper is an interface made for aiding with testing the mongo API layer.
type IDatabaseHelper interface {
	NewClient(uri string) (*mongo.Client, error)
	Collection(name string) (mongoapi.ICollectionHelper, error)
}
