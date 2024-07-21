package repository

import (
	"context"
	"learn-golang/src/db/sqlc"
	"learn-golang/src/model"

	"github.com/jackc/pgx/v5/pgtype"
)

type PostRepository interface {
	GetPost(ctx context.Context, id string) (*model.Post, error)
	CreatePost(ctx context.Context, title string, content *string) (*model.Post, error)
}

type postRepository struct {
	queries       *sqlc.Queries
	idGenerator   IdGenerator
	timeGenerator TimeGenerator
}

func NewPostRepository(queries *sqlc.Queries, idGenerator IdGenerator, timeGenerator TimeGenerator) PostRepository {
	return &postRepository{
		queries:       queries,
		idGenerator:   idGenerator,
		timeGenerator: timeGenerator,
	}
}

func (p *postRepository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	dbId := pgtype.UUID{}
	err := dbId.Scan(id)
	if err != nil {
		return nil, err
	}
	post, err := p.queries.GetPost(ctx, dbId)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:        uuidToString(&post.ID),
		UserID:    uuidToString(&post.UserID),
		Title:     post.Title,
		Content:   &post.Content,
		CreatedAt: timestampToTime(&post.CreatedAt),
	}, nil
}

func (p *postRepository) CreatePost(ctx context.Context, title string, content *string) (*model.Post, error) {
	id, err := p.idGenerator.Generate()
	if err != nil {
		return nil, err
	}
	createdAt, err := p.timeGenerator.Now()
	if err != nil {
		return nil, err
	}
	// TODO: ユーザーを取得する処理を追加する
	user, err := p.queries.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	// TODO contentがnilの場合はNULLを保存するようにしたい
	contentValue := ""
	if content != nil {
		contentValue = *content
	}

	post, err := p.queries.CreatePost(ctx, sqlc.CreatePostParams{ID: *id, UserID: user.ID, Title: title, Content: contentValue, CreatedAt: *createdAt})
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:        uuidToString(&post.ID),
		UserID:    uuidToString(&post.UserID),
		Title:     post.Title,
		Content:   &post.Content,
		CreatedAt: timestampToTime(&post.CreatedAt),
	}, nil
}
