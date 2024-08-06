-- name: CreateUser :one
INSERT INTO users (username, password, role)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
WHERE role = $1
ORDER BY id LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    password = $3,
    role     = $4
WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;