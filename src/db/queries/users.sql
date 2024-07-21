-- name: CreateUser :one
INSERT INTO users (id, name, created_at) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUser :one
SELECT * FROM users LIMIT 1;