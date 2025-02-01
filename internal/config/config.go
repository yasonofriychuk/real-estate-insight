package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	pgUrl    string
	httpPort string
}

func MustNewConfigWithEnv() *Config {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error loading .env file: %s", err))
	}
	return &Config{
		pgUrl:    mustGetEnv("PG_URL"),
		httpPort: mustGetEnv("HTTP_PORT"),
	}
}

func (c *Config) PgUrl() string {
	return c.pgUrl
}

func (c *Config) HttpPort() string {
	return c.httpPort
}

func mustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}
	return value
}
