package config

import (
	"log"
	"net/url"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ConnectionString string `mapstructure:"connection_string"`
	Port             string `mapstructure:"port"`
	Config           []string
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configuration ...")

	// Load .env file into environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, skipping...")
	}

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal(err)
	}

	AppConfig.Config = []string{
		"{MYSQL_DBNAME}", viper.GetString("MYSQL_DBNAME"),
		"{MYSQL_USERNAME}", viper.GetString("MYSQL_USERNAME"),
		"{MYSQL_PASSWORD}", url.QueryEscape(url.QueryEscape(viper.GetString("MYSQL_PASSWORD")))}
}

func LoadEnvDBConfigPropeties() {
	log.Println("Loading env db properties")

}
