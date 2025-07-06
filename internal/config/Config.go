package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Configurate() {
	_ = godotenv.Load()

	env := os.Getenv("APP_ENV")
	log.Println("APP_ENV:", env)
	if env == "Dev" {
		viper.SetConfigName("config")
	}
	if env == "Docker" {
		viper.SetConfigName("config_docker")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
