package main

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", mongoapi.HelloWorld)

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

	client := mongoapi.MongoClient()

	gamerCollection := client.Database("BR").Collection("Gamers")

	// err := mongoapi.AddToCollection(gamerCollection, gamer1, gamer2, gamer3)
	// if err != nil {
	// 	fmt.Println("Failed to add items to collection", err)
	// }

	// err = mongoapi.AddToCollection(gamerCollection, gamer4)

	// fmt.Printf("Inserted Records with IDs: %v", results)

	filter := bson.D{{"name", "Bean"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := gamerCollection.UpdateOne(context.TODO(), filter, update)

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
	http.ListenAndServe("localhost:8080", r)

}
