package main

import (
	"MongoDBGolang/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	gamerCollection := client.Database("BR").Collection("Gamers")

	games1 := []models.Game{{"Megaman", "Platformer", "NES"}, {"Daggerfall", "CRPG", "PC"}, {"Zelda64", "Adventure", "N64"}}
	gamer1 := models.Gamer{
		"Mark",
		38,
		games1,
	}
	games2 := []models.Game{{"Pocky & Rocky", "Action", "SNES"}, {"Final Fantasy III", "JRPG", "SNES"}, {"Contra Hard Corps", "Action", "Sega Genesis"}}
	gamer2 := models.Gamer{
		"Brian",
		37,
		games2,
	}
	games3 := []models.Game{{"Shining Force", "TRPG", "Sega Genesis"}, {"Shinobi", "Action", "Arcade"}, {"Phantasy Star", "JRPG", "Sega Master System"}}
	gamer3 := models.Gamer{
		"Mark",
		28,
		games3,
	}

	results, err := gamerCollection.InsertMany(context.TODO(), []interface{}{gamer1, gamer2, gamer3})
	if err != nil {
		log.Fatal("Failed to insert documents", err)
	}

	// filter := bson.M{{"name", "Ash"}}
	// update := bson.M{
	// 	{"$inc", bson.M{
	// 		{"age", 1},
	// 	}},
	// }

	// updateResult, err := collection.UpdateMany(context.TODO(), filter, update)

	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// findOptions := options.Find()
	// findOptions.SetLimit(2)

	// cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// 	log.Fatal("error 3", err)
	// }

	// for cur.Next(context.TODO()) {

	// 	var elem models.Gamer
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal("error 2", err)
	// 	}

	// 	// results = append(results, &elem)
	// }

	// err = cur.Err()
	// if err != nil {
	// 	log.Fatal("error 1", err)
	// }

	// cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents: %v\n", results)

}
