package main

import (
	"context"
	"learn-golang/src/di"
	"learn-golang/src/graph"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	resolver, cleanup, err := di.InitializeResolver(ctx)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer cleanup()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("connect to http://localhost:8080/ for GraphQL playground")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {

		slog.Error("failed to listen and serve")
		return
	}
}
