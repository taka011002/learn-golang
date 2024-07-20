package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	url := "postgres://golang:golang@localhost:5432/golang" // TODO 環境変数から取得する
	db, err := sql.Open("pgx", url)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/db/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
