package appconfig

import (
	"github.com/spf13/viper"
)

// Config is a struct with commong variables for the app.
type Config struct {
	Port     string
	MongoCfg *MongoConfig
}

// MongoConfig is a struct defining the MongoDB address, database name, and collection name.
type MongoConfig struct {
	MongoURI   string
	Database   string
	Collection string
}

// TODO: Consider setting a config and passing it down to the app.
// func setConfig() (*app.Config, error) {
func setConfig() *Config {
	config := Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	config.MongoCfg = &MongoConfig{
		MongoURI:   viper.GetString("mongo_uri"),
		Database:   viper.GetString("mongo_database"),
		Collection: viper.GetString("mongo_collection"),
	}

	config.Port = viper.GetString("port")

	return &config
}
