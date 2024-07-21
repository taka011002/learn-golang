package usecase

import (
	"context"
	"learn-golang/src/db"
	"learn-golang/src/db/sqlc"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestUseCase_CreateUser(t *testing.T) {
	ctx := context.Background()

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
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get container host: %s", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		log.Fatalf("failed to get container port: %s", err)
	}

	config := db.Config{
		Host:     host,
		Port:     port.Port(),
		User:     dbUser,
		Password: dbPassword,
		DbName:   dbName,
	}
	err = db.Migrate(&config)
	if err != nil {
		log.Fatalf("failed to migrate db: %s", err)
	}

	conn, err := db.NewDbClient(ctx, &config)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	defer conn.Close(ctx)
	queries := sqlc.New(conn)

	useCase := UseCase{
		queries: queries,
	}

	user, err := useCase.CreateUser(ctx, "test")
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	if user.Name != "test" {
		t.Errorf("failed to create user: %s", err)
	}
}
