package app

import (
	"adiubaidah/simple-crud/helper"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func GetConfig() Config {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	helper.PanicIfError(err)

	return Config{
		DBHost: config.GetString("DB_HOST"),
		DBPort: config.GetString("DB_PORT"),
		DBUser: config.GetString("DB_USERNAME"),
		DBPass: config.GetString("DB_PASSWORD"),
		DBName: config.GetString("DB_NAME"),
	}

}
