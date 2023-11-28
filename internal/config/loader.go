package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Get() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file: %s", err))
	}
	
	return &Config{
		Server: Server{
			Host: viper.GetString("SERVER_HOST"),
			Port: viper.GetString("SERVER_PORT"),
		},
		Database: Database{
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			User:     viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Name:     viper.GetString("DATABASE_NAME"),
		},
		Cloudinary: Cloudinary{
			CloudName: viper.GetString("CLOUDINARY_CLOUD_NAME"),
			ApiKey:    viper.GetString("CLOUDINARY_API_KEY"),
			ApiSecret: viper.GetString("CLOUDINARY_API_SECRET"),
		},
		SecretKey: SecretKey{
			SecretKey: viper.GetString("SECRET_KEY"),
		},
	}
}
