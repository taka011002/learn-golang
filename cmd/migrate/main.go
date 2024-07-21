package main

import (
	"learn-golang/src/db"
	"log/slog"
)

func main() {
	config := db.NewConfig()
	err := db.Migrate(config)
	if err != nil {
		slog.Error(err.Error())
	}
}
