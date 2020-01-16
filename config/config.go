package appconfig

// Config is a struct with commong variables for the app.
type Config struct {
	MongoCfg *MongoConfig
}

// MongoConfig is a struct defining the MongoDB address, database name, and collection name.
type MongoConfig struct {
	MongoURI   string
	Database   string
	Collection string
}
