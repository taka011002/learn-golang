package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDbClient(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, getUrl())
}

func NewDb() (*sql.DB, error) {
	return sql.Open("pgx", getUrl())
}

func getUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), "localhost", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}
