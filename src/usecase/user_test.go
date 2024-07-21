package usecase

import (
	"context"
	"learn-golang/src/model"
	"learn-golang/src/repository"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func newUserUseCase(t *testing.T) *userUseCase {
	queries := setUpIT(t)
	return &userUseCase{
		repo: repository.NewUserRepository(queries, repository.NewUuidGenerator(), repository.NewTimeGenerator()),
	}
}

func TestUseCase_CreateUser(t *testing.T) {
	ctx := context.Background()
	useCase := newUserUseCase(t)

	cases := []struct {
		name     string
		input    string
		wantName string
	}{
		{
			name:     "test",
			input:    "test",
			wantName: "test",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := useCase.CreateUser(ctx, tc.input)
			if err != nil {
				t.Errorf("CreateUser() error = %v, want nil", err)
			}
			if user.Name != tc.wantName {
				t.Errorf("CreateUser() = %v, want %v", user.Name, tc.wantName)
			}
		})
	}
}

func TestUseCase_GetUser(t *testing.T) {
	ctx := context.Background()
	useCase := newUserUseCase(t)

	// TODO: IDを動的に決定せず、固定値にしてテストしたい
	testUsers := []model.User{
		{Name: "test"},
	}
	for i, tu := range testUsers {
		u, err := useCase.CreateUser(ctx, tu.Name)
		if err != nil {
			t.Fatalf("CreateUser() error = %v, want nil", err)
		}
		testUsers[i].ID = u.ID
		testUsers[i].CreatedAt = u.CreatedAt
	}

	cases := []struct {
		name     string
		id       string
		wantName string
	}{
		{
			name:     "success",
			id:       testUsers[0].ID,
			wantName: "test",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := useCase.GetUser(ctx, tc.id)
			if err != nil {
				t.Errorf("GetUser() error = %v, want nil", err)
			}
			if user.Name != tc.wantName {
				t.Errorf("GetUser() = %v, want %v", user.Name, tc.wantName)
			}
		})
	}
}
