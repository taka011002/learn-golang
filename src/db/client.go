package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDbClient(ctx context.Context) (*pgx.Conn, error) {
	url := "postgres://golang:golang@localhost:5432/golang" // TODO 環境変数から取得する
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	return conn, nil
}
