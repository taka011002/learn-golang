-- name: GetUser :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (id, name, project_v2) VALUES ($1, $2, $3) RETURNING *;