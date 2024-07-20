package main

import (
	"learn-golang/src/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	d, err := db.NewDb()
	defer d.Close()
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(d, &postgres.Config{})
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
