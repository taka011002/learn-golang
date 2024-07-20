package main

import (
	"context"
	"learn-golang/src/db"
	"learn-golang/src/db/sqlc"
	"learn-golang/src/graph"
	"learn-golang/src/usecase"
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

	conn, err := db.NewDbClient(ctx)
	defer conn.Close(ctx)
	if err != nil {
		slog.Error("failed to connect to db")
		return
	}
	queries := sqlc.New(conn)
	useCase := usecase.NewUseCase(queries)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(*useCase)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	slog.Info("connect to http://localhost:8080/ for GraphQL playground")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {

		slog.Error("failed to listen and serve")
		return
	}
}
