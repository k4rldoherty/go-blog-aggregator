-- sqlfluff: disable=L006,L009,L014
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE users.name = $1;
-- sqlfluff: enable=L006,L009,L014

-- name: DeleteAllUsers :exec
DELETE FROM users;


-- name: GetUsers :many
SELECT *
  FROM users;
