package usecase

import (
	"context"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestUseCase_CreateUser(t *testing.T) {
	ctx := context.Background()
	useCase := setUpIT(t)

	cases := []struct {
		name     string
		input    string
		wantName string
	}{
		{
			name:     "Valid user",
			input:    "test",
			wantName: "test",
		},
		{
			name:     "Empty user",
			input:    "",
			wantName: "",
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
	useCase := setUpIT(t)

	testUsers := []string{"test", "test2"}

	for _, tu := range testUsers {
		_, err := useCase.CreateUser(ctx, tu)
		if err != nil {
			t.Fatalf("CreateUser() error = %v, want nil", err)
		}
	}

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
		{
			name:     "test2",
			input:    "test2",
			wantName: "test2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := useCase.GetUser(ctx, tc.input)
			if err != nil {
				t.Errorf("GetUser() error = %v, want nil", err)
			}
			if user.Name != tc.wantName {
				t.Errorf("GetUser() = %v, want %v", user.Name, tc.wantName)
			}
		})
	}
}
