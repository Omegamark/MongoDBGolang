package main

import (
	"MongoDBGolang/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

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

	gamerCollection, err = client.Database("BR").Collection("Gamers")

	games1 := []models.Game{{"Megaman", "Platformer", "NES"}, {"Daggerfall", "CRPG", "PC"}}
	// gamer1 := models.Gamer{
	// 	"Mark",
	// 	38,
	// 	games1,
	// }

	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	// thing := map[string]interface{}{
	// 	"Name": "Jack",
	// 	"Age":  "Flack",
	// }

	// if err != nil {
	// 	fmt.Println("I'm an error man.", err)
	// }

	// manyThings := []interface{}{ash, misty, brock}

	// manyThingsResult, err := client.Database("test").Collection("trainers").InsertMany(context.TODO(), manyThings)
	// if err != nil {
	// 	fmt.Println("FAILED MANY")
	// }

	// fmt.Printf("Inserted item IDs: %s", manyThingsResult.InsertedIDs)

	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	collection := client.Database("test").Collection("trainers")
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("error 4", err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal("error 3", err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Gamer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("error 2", err)
		}

		// results = append(results, &elem)
	}

	err = cur.Err()
	if err != nil {
		log.Fatal("error 1", err)
	}

	cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents: %v\n", results)

}
