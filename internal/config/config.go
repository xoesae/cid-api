package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type DatabaseConfig struct {
}

type Config struct {
	Port       string `mapstructure:"PORT"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbHostCli  string `mapstructure:"DB_HOST_CLI"`
	DbPort     string `mapstructure:"DB_PORT"`
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading env, %s", err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("error on parse config, %s", err)
		}
	})

	return config
}
