package mongoapi_test

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

// func Test_Login(t *testing.T) {

// 	var dataAccess = data.DummyUserDataAccess{}
// 	dataAccess.New()
// 	var handler UserHandler = UserHandler{&dataAccess}

// 	user := &models.User{Username: "JohnSmith", Password: "Password"}
// 	jsonPerson, _ := json.Marshal(user)

// 	request, _ := http.NewRequest("POST", "/User/Login", bytes.NewBuffer(jsonPerson))
// 	response := httptest.NewRecorder()
// 	decoder := json.NewDecoder(response.Body)

// 	handler.Login(response, request)

// 	var u models.User
// 	if err := decoder.Decode(&u); err != nil {
// 		t.Errorf("User does not exist.")
// 	}

// 	if u.Address != "200 Place Zero" || u.City != "City Zero" {
// 		t.Errorf("User's information doesnt match.")
// 	}
// }

type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}

type CollectionHelper interface {
	FindOne(context.Context, interface{}) SingleResultHelper
	InsertOne(context.Context, interface{}) (interface{}, error)
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}

type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect() error
	StartSession() {
		cl *mongo.Client
	}
}
