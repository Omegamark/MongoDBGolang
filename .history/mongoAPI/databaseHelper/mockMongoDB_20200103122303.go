package databasehelper

import clienthelper "MongoDBGolang/mongoAPI/clientHelper"

// MockDatabaseHelper ***needs to be renamed to make more sense. Returns a mock instance of a MongoDB Client.
type MockDatabaseHelper struct {
}

func (dd MockDatabaseHelper) NewClient(uri string) (clienthelper.IClientHelper, error) {
	return clienthelper.MockClientHelper{}, nil
}
