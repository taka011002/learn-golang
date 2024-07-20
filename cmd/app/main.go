package main

import (
	"log/slog"
	"net/http"
	"os"
)

const (
	exitOK = iota
	exitErr
)

func main() {
	os.Exit(run())
}

func run() int {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			return
		}
	})

	slog.Info("Server Start", slog.String("port", ":8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return exitErr
	}

	return exitOK
}
