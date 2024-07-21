package usecase

import (
	"context"
	"learn-golang/src/db"
	"learn-golang/src/db/sqlc"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func initDbContainer(t *testing.T, ctx context.Context) (*db.Config, error) {
	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16.3",
		//postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("Could not stop posql: %s", err)
		}
	})

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return nil, err
	}

	config := db.Config{
		Host:     host,
		Port:     port.Port(),
		User:     dbUser,
		Password: dbPassword,
		DbName:   dbName,
	}

	// DB migrationも実行する
	err = db.Migrate(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func setUpIT(t *testing.T) *UseCase {
	t.Helper()

	ctx := context.Background()

	config, err := initDbContainer(t, ctx)
	if err != nil {
		t.Fatal("failed to initialize db container")
	}

	conn, cleanup, err := db.NewDbClient(ctx, config)
	if err != nil {
		t.Fatal("failed to connect db")
	}
	t.Cleanup(cleanup)

	queries := sqlc.New(conn)
	useCase := UseCase{
		queries: queries,
	}

	return &useCase
}
