-- name: CreatePost :one
INSERT INTO posts (id, user_id, title, content, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;