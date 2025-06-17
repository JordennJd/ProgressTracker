package config

import (
	"github.com/spf13/viper"
)

func Configurate() {
	viper.SetConfigName("config") // имя файла без расширения
	viper.SetConfigType("yaml")   // или json, toml
	viper.AddConfigPath(".")      // путь к файлу конфигурации

	// Автоматическое чтение переменных окружения
	viper.AutomaticEnv()

	// Чтение конфиг-файла
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
