package config

import (
	"fmt"
	"os"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDbName   string
}

func ReadConfig() *Config {
	return &Config{
		PostgresUser:     os.Getenv("DATABASE_USERNAME"),
		PostgresPassword: os.Getenv("DATABASE_PASSWORD"),
		PostgresHost:     os.Getenv("DATABASE_HOST"),
		PostgresPort:     os.Getenv("DATABASE_PORT"),
		PostgresDbName:   os.Getenv("DATABASE_DBNAME"),
	}
}

func (c *Config) GetPostgresConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDbName,
	)
}