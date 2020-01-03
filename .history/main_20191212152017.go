package main

import (
	"MongoDBGolang/models"
	mongoapi "MongoDBGolang/mongoAPI"
	"fmt"
)

func main() {

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
		Name:     "Mark",
		Age:      28,
		Gamelist: games3,
	}

	client := mongoapi.MongoClient()

	gamerCollection := client.Database("BR").Collection("Gamers")
	gamerlist := []models.Gamer{}{gamer1, gamer2, gamer3}

	err := mongoapi.AddToCollection(gamerCollection, []models.Gamer{}{gamer1, gamer2, gamer3})
	if err != nil {
		fmt.Println("Failed to add items to collection", err)
	}
	// results, err := gamerCollection.InsertMany(context.TODO(), []interface{}{gamer1, gamer2, gamer3})
	// if err != nil {
	// 	log.Fatal("Failed to insert documents", err)
	// }

	fmt.Printf("Inserted Records with IDs: %v", results)

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
