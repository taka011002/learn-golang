package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDbClient(ctx context.Context, c *Config) (*pgx.Conn, error) {
	return pgx.Connect(ctx, c.toUrl())
}

func NewDb(c *Config) (*sql.DB, error) {
	return sql.Open("pgx", c.toUrl())
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewConfig() *Config {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	return &Config{
		Host:     host,
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
}

func (c *Config) toUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DbName)
}
