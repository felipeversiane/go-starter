package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	once sync.Once
)

type Config struct {
	Database     DatabaseConfig
	Server       ServerConfig
	Log          LogConfig
}

type ConfigInterface interface {
	GetDatabaseConfig() DatabaseConfig
	GetServerConfig() ServerConfig
	GetLogConfig() LogConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

type LogConfig struct {
	Level string
}


func NewConfig() ConfigInterface {
	var cfg *Config
	once.Do(func() {
		cfg = &Config{
			Database: DatabaseConfig{
				Host:     getEnvOrDie("POSTGRES_HOST"),
				Port:     getEnvOrDie("POSTGRES_PORT"),
				User:     getEnvOrDie("POSTGRES_USER"),
				Password: getEnvOrDie("POSTGRES_PASSWORD"),
				Name:     getEnvOrDie("POSTGRES_DB"),
			},
			Server: ServerConfig{
				Port: getEnvOrDie("API_PORT"),
			},
			Log: LogConfig{
				Level: getEnvOrDie("LOG_LEVEL"),
			},
		}
	})
	return cfg
}

func getEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("missing environment variable %s", key))
	}
	return value
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *Config) GetServerConfig() ServerConfig {
	return c.Server
}

func (c *Config) GetLogConfig() LogConfig {
	return c.Log
}
