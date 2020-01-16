package main

import (
	"MongoDBGolang/app"
	appconfig "MongoDBGolang/config"

	"github.com/spf13/viper"
)

func main() {
	config := setConfig()
	// TODO: Have StartUp return an error and shutdown the app.
	app.StartUp(config)
}

// TODO: Consider setting a config and passing it down to the app.
// func setConfig() (*app.Config, error) {
func setConfig() *appconfig.Config {
	config := appconfig.Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	config.MongoCfg = &appconfig.MongoConfig{
		MongoURI:   viper.GetString("mongo_uri"),
		Database:   viper.GetString("mongo_database"),
		Collection: viper.GetString("mongo_collection"),
	}

	config.Port = viper.GetString("port")

	return &config
}
