package main

import (
	"MongoDBGolang/handlers"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/addGamer", handlers.AddToCollection).Methods("POST")
	r.HandleFunc("/deleteGamer", handlers.DeleteOneFromCollection).Methods("DELETE")
	r.HandleFunc("/findGamer", handlers.FindOneInCollection).Methods("GET")

	// games1 := []models.Game{{"Megaman", "Platformer", "NES"}, {"Daggerfall", "CRPG", "PC"}, {"Zelda64", "Adventure", "N64"}}
	// gamer1 := models.Gamer{
	// 	"Mark",
	// 	38,
	// 	games1,
	// }
	// games2 := []models.Game{{"Pocky & Rocky", "Action", "SNES"}, {"Final Fantasy III", "JRPG", "SNES"}, {"Contra Hard Corps", "Action", "Sega Genesis"}}
	// gamer2 := models.Gamer{
	// 	"Brian",
	// 	37,
	// 	games2,
	// }
	// games3 := []models.Game{{"Shining Force", "TRPG", "Sega Genesis"}, {"Shinobi", "Action", "Arcade"}, {"Phantasy Star", "JRPG", "Sega Master System"}}
	// gamer3 := models.Gamer{
	// 	"Sean",
	// 	28,
	// 	games3,
	// }
	// games4 := []models.Game{{"Chess", "Board", "Commodore"}, {"Bakarat", "Ancient", "Wood"}, {"Cyberpunk 2077", "WRPG", "PC"}}
	// gamer4 := models.Gamer{
	// 	"Bean",
	// 	65,
	// 	games4,
	// }

	// client := mongoapi.MongoClient()
	// gamerCollection := client.Database("BR").Collection("Gamers")

	// ************************ C - Create ***************************

	// err := mongoapi.AddToCollection(gamerCollection, gamer1, gamer2, gamer3)
	// if err != nil {
	// 	fmt.Println("Failed to add items to collection", err)
	// }

	// err = mongoapi.AddToCollection(gamerCollection, gamer4)

	// fmt.Printf("Inserted Records with IDs: %v", results)

	// filter := bson.M{
	// 	"name": "Sean",
	// }
	// update := bson.M{
	// 	"$set": bson.M{
	// 		"age": 63,
	// 	},
	// }

	// *** Update an individual field in a specific entry in Mongo
	// updateResult, err := gamerCollection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	fmt.Println("Failed to update: ", err)
	// }

	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// ********************** R - READ **********************

	// var result models.Gamer
	// projectionOpts := bson.M{
	// 	"gamelist": 1,
	// 	"age":      1,
	// }

	// *** Target an individual field in a specific record in mongo
	// err := gamerCollection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectionOpts)).Decode(&result)
	// if err != nil {
	// 	fmt.Println("Drrr", err)
	// }

	// fmt.Println("RESULTS: ", result)

	// *** Push a new item to an array to
	// ***********************************
	// game := models.Game{
	// 	Name:     "Do The Thing!",
	// 	Genre:    "Please Just Work Already!",
	// 	Platform: "Don't Care!!!",
	// }

	// update := bson.M{
	// 	"$set": bson.M{
	// 		"age": 63,
	// 	},
	// }
	// *** NOTE: *** This is correct for JSON coming into the application. There will be no reason to marshal what is already json into json.
	updateJSON, err := json.Marshal(game)
	var updateJSONMap map[string]interface{}
	err = json.Unmarshal(updateJSON, &updateJSONMap)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("****************")
	// fmt.Println("Derrr: ", updateJSONMap)
	// fmt.Println("****************")
	// ************************ U - Update ***********************

	// update := bson.M{
	// 	"$gamelist": bson.M{
	// 		"$push": bson.M{
	// 			"$gamelist": updateJSONMap,
	// 		},
	// 	},
	// }
	update := bson.M{
		"$push": bson.M{
			"gamelist": updateJSONMap,
		},
	}

	collection, err := gamerCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Failed to update: ", err)
	}

	// fmt.Println("Collection: ", collection.UpsertedID)
	// *********************** D - Delete ************************
	// filterHuh := bson.M{
	// 	"name": "Mark",
	// }
	// results, err := gamerCollection.DeleteOne(context.TODO(), filterHuh)
	// if err != nil {
	// 	fmt.Println("Error DELETING: ", err)
	// }
	// fmt.Println(results.DeletedCount)

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
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
