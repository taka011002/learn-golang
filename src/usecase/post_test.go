package usecase

import (
	"context"
	"learn-golang/src/model"
	"learn-golang/src/repository"
	"testing"
)

func setup(t *testing.T, ctx context.Context) PostUseCase {
	queries := setUpIT(t)
	useCase := &postUseCase{
		repo: repository.NewPostRepository(queries, repository.NewUuidGenerator(), repository.NewTimeGenerator()),
	}

	// Postの作成でUserが必要なため、Userを作成する
	// TODO: もう少し綺麗に書きたい
	userRepository := repository.NewUserRepository(queries, repository.NewUuidGenerator(), repository.NewTimeGenerator())
	userRepository.CreateUser(ctx, "test")

	return useCase
}

func stringPtr(s string) *string {
	return &s
}

func TestPostUseCase_CreatePost(t *testing.T) {
	ctx := context.Background()
	useCase := setup(t, ctx)

	cases := []struct {
		name        string
		title       string
		content     *string
		wantTitle   string
		wantContent *string
	}{
		{
			name:        "Success",
			title:       "test",
			content:     stringPtr("test"),
			wantTitle:   "test",
			wantContent: stringPtr("test"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			post, err := useCase.CreatePost(ctx, tc.title, tc.content)
			if err != nil {
				t.Errorf("CreatePost() error = %v, want nil", err)
			}
			if post.Title != tc.wantTitle {
				t.Errorf("CreatePost() = %v, want %v", post.Title, tc.wantTitle)
			}
			if *post.Content != *tc.wantContent {
				t.Errorf("CreatePost() = %v, want %v", *post.Content, *tc.wantContent)
			}
		})
	}
}

func TestPostUseCase_GetPost(t *testing.T) {
	ctx := context.Background()
	useCase := setup(t, ctx)

	// TODO: IDを動的に決定せず、固定値にしてテストしたい
	testPosts := []model.Post{
		{Title: "test", Content: stringPtr("test")},
	}
	for i, tu := range testPosts {
		p, err := useCase.CreatePost(ctx, tu.Title, tu.Content)
		if err != nil {
			t.Fatalf("CreatePost() error = %v, want nil", err)
		}
		testPosts[i].ID = p.ID
		testPosts[i].UserID = p.UserID
		testPosts[i].CreatedAt = p.CreatedAt
	}

	cases := []struct {
		name        string
		title       string
		content     *string
		wantTitle   string
		wantContent *string
	}{
		{
			name:        "Success",
			title:       "test",
			content:     stringPtr("test"),
			wantTitle:   "test",
			wantContent: stringPtr("test"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			post, err := useCase.GetPost(ctx, testPosts[0].ID)
			if err != nil {
				t.Errorf("GetPost() error = %v, want nil", err)
			}
			if post.Title != tc.wantTitle {
				t.Errorf("GetPost() = %v, want %v", post.Title, tc.wantTitle)
			}
			if *post.Content != *tc.wantContent {
				t.Errorf("GetPost() = %v, want %v", *post.Content, *tc.wantContent)
			}
		})
	}

}
