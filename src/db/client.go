package db

import (
	"context"
	"database/sql"
	"fmt"
	"learn-golang/src/db/sqlc"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func newDbClient(ctx context.Context, c *Config) (*pgx.Conn, func(), error) {
	conn, err := pgx.Connect(ctx, c.toUrl())
	if err != nil {
		return nil, nil, err
	}

	return conn, func() {
		conn.Close(ctx)
	}, nil
}

func NewQueries(ctx context.Context, c *Config) (*sqlc.Queries, func(), error) {
	conn, cleanup, err := newDbClient(ctx, c)
	if err != nil {
		return nil, nil, err
	}

	queries := sqlc.New(conn)
	return queries, cleanup, nil

}

func NewDb(c *Config) (*sql.DB, func(), error) {
	db, err := sql.Open("pgx", c.toUrl())
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		db.Close()
	}, nil
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
	return fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		c.Host, c.User, c.Password, c.Port, c.DbName)
}
