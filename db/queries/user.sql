-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set name = $2, email = $3, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: UpdateUserEmailVerifiedAt :exec
UPDATE users
  set email_verified_at = EXTRACT(EPOCH FROM NOW()), updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;