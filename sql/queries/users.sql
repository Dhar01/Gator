-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY name;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;