package mockmongoapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}

type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

func NewClient() (ClientHelper, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	return &mongoClient{cl: c}, err
}

func NewDatabase(databaseName string, client ClientHelper) DatabaseHelper {
	return client.Database(databaseName)
}

func (mc *mongoClient) Connect() error {
	return mc.cl.Connect(nil)
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)

	return id.InsertedID, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
