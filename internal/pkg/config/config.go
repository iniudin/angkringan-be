package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"url"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
}

func New() *Config {
	config := Config{}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
