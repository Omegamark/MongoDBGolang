package mongoapiredux

import (
	"MongoDBGolang/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDatabase represents a connection to a mongo database
type MongoDatabase struct {
	DBConnection MongoConnection
}

/*

	CREATE

*/

// AddToGamerCollection adds one or many new gamer(s) to the database. You may add an arbitrary number of gamers.
func (db *MongoDatabase) AddToGamerCollection(gamer ...interface{}) error {
	collection, err := db.DBConnection.GetCollection()
	if err != nil {
		return err
	}

	if gamer[0] == nil {
		return errors.New("Nothing to add to DB")
	}

	if len(gamer) > 1 {
		_, err := collection.InsertMany(context.TODO(), gamer)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = collection.InsertOne(context.TODO(), gamer[0])
	if err != nil {
		return err
	}
	return nil
}

/*

	READ

*/

// FindOneInCollection returns a gamer containing the fields
func (db *MongoDatabase) FindOneInCollection(gamerName interface{}, projections []interface{}) (*models.Gamer, error) {
	collection, err := db.DBConnection.GetCollection()

	if gamerName == nil {
		return nil, errors.New("Must enter a gamer name")
	}

	var result models.Gamer

	// TODO: Make filters build more dynamically.
	filter := bson.M{
		"name": gamerName,
	}

	projectionOpts := bson.M{}
	for _, key := range projections {
		projectionOpts[key.(string)] = 1
	}

	err = collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectionOpts)).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// /*

// 	UPDATE

// */

// // UpdateOneGamerByName allows the gamers name and age to be changed.
// func (db *MongoDatabase) UpdateOneGamerByName(collection collectionhelper.ICollectionHelper, gamerName string, updateInfo interface{}) error {
// 	if gamerName == "" {
// 		return errors.New("Must enter a gamer name")
// 	}

// 	if updateInfo == nil {
// 		return errors.New("Must enter update information")
// 	}

// 	filter := bson.M{
// 		"name": gamerName,
// 	}

// 	update := bson.M{
// 		"$set": bson.M{
// 			"age": updateInfo,
// 		},
// 	}

// 	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	log.Printf("Matched %v document and updated %v document.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
// 	return nil
// }

// // AddGameToGamerGamelist adds a game to a gamer's game list.
// func (db *MongoDatabase) AddGameToGamerGamelist(collection collectionhelper.ICollectionHelper, listUpdate models.GamelistUpdate) error {
// 	// TODO: Make filter and update more dynamic
// 	filter := bson.M{
// 		"name": listUpdate.Name,
// 	}

// 	log.Println("List Update: ", listUpdate)

// 	update := bson.M{
// 		"$push": bson.M{
// 			"gamelist": listUpdate.Game,
// 		},
// 	}

// 	_, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		fmt.Println("Failed to update: ", err)
// 		return err
// 	}
// 	return nil
// }

/*

	DELETE

*/

// DeleteOneGamerFromCollectionByName deletes any gamers passed as arguments.
func (db *MongoDatabase) DeleteOneGamerFromCollectionByName(gamerName interface{}) error {
	// TODO: Return a success response.
	collection, err := db.DBConnection.GetCollection()
	if err != nil {
		return err
	}

	var filter bson.M

	if gamerName != nil {
		filter = bson.M{
			"name": gamerName,
		}
	} else {
		return errors.New("Name argument cannot be blank or <nil>")
	}

	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// // DropCollection drops the passed collection from the database.
// func (db *MongoDatabase) DropCollection(collection *mongo.Collection) error {
// 	err := collection.Drop(context.TODO())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
