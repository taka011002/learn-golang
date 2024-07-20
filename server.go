package main

import (
	"learn-golang/graph"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {

		slog.Error("failed to listen and serve")
		return
	}
}
