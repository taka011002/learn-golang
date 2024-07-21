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

	_ "github.com/jackc/pgx/v5/stdlib"
)

func initDbContainer(ctx context.Context) (*db.Config, error) {
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

	// Clean up the container
	//defer func() {
	//	if err := postgresContainer.Terminate(ctx); err != nil {
	//		log.Fatalf("failed to terminate container: %s", err)
	//	}
	//}()

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

func newUseCase(ctx context.Context, config db.Config) (*UseCase, error) {
	conn, err := db.NewDbClient(ctx, &config)
	if err != nil {
		return nil, err
	}
	queries := sqlc.New(conn)
	//defer conn.Close(ctx)

	useCase := UseCase{
		queries: queries,
	}
	return &useCase, nil
}

func TestUseCase_CreateUser(t *testing.T) {
	ctx := context.Background()
	congig, err := initDbContainer(ctx)
	if err != nil {
		t.Errorf("failed to init db container: %s", err)
	}
	useCase, err := newUseCase(ctx, *congig)
	if err != nil {
		t.Errorf("failed to create usecase: %s", err)
	}

	user, err := useCase.CreateUser(ctx, "test")
	if err != nil {
		t.Errorf("failed to create user: %s", err)
	}

	if user.Name != "test" {
		t.Errorf("failed to create user: %s", err)
	}
}
