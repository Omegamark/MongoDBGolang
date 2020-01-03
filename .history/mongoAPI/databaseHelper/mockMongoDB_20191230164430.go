package databasehelper

import clienthelper "MongoDBGolang/mongoAPI/clientHelper"

type MockDatabaseHelper struct {
}

func (dd MockDatabaseHelper) NewClient(uri string) (*clienthelper.IClientHelper, error) {
	return &clienthelper.MockClientHelper{}, nil
}
