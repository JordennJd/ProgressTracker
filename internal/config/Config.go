package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Configurate() {
	_ = godotenv.Load()
	//if err != nil {
	//	panic("Error loading .env file")
	//}
	env := os.Getenv("APP_ENV")
	log.Println("APP_ENV:", env)
	if env == "Dev" {
		viper.SetConfigName("config") // имя файла без расширения
	}
	if env == "Docker" {
		viper.SetConfigName("config_docker") // имя файла без расширения
	}
	viper.SetConfigType("yaml") // или json, toml
	viper.AddConfigPath(".")    // путь к файлу конфигурации

	// Автоматическое чтение переменных окружения
	viper.AutomaticEnv()

	// Чтение конфиг-файла
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
