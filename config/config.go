package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost          string
	DBMigrationHost string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	config := &Config{
		DBHost:          os.Getenv("DB_HOST"),
		DBMigrationHost: os.Getenv("DB_MIGRATION_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBUser:          os.Getenv("DB_USER"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
	}

	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfig(config *Config) error {
	variables := reflect.ValueOf(config).Elem()
	numberOfFields := variables.NumField()

	for i := 0; i < numberOfFields; i++ {
		field := variables.Field(i)
		fieldName := variables.Type().Field(i).Name

		if field.Kind() == reflect.String && field.String() == "" {
			return fmt.Errorf("the env variable %s is required", fieldName)
		}
	}

	return nil
}
