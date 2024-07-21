package db

import (
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(config *Config) error {
	d, err := NewDb(config)
	defer d.Close()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(d, &postgres.Config{})
	if err != nil {
		return err
	}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	url := "file://" + basePath + "/migrations"

	m, err := migrate.NewWithDatabaseInstance(
		url,
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
