package config

import (
	"fmt"
	"os"
)

type Config struct {
	pgUrl    string
	httpPort string
}

func MustNewConfigWithEnv() *Config {
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
