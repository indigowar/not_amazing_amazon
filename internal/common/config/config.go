package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LogFile  string
	Postgres Postgres
	Redis    Redis
	Minio    Minio
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int
	Db       int
	Password string
}

type Minio struct {
	Host         string
	Port         int
	RootUser     string
	RootPassword string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/appname")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.SetDefault("LogFile", "app.log")
	viper.SetDefault("Postgres.Port", 5432)
	viper.SetDefault("Redis.Port", 6379)
	viper.SetDefault("Minio.Port", 9000)

	// Read configuration from file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}
