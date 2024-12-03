-- name: RegisterUser :one
INSERT INTO users (
  first_name, last_name, email, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
  set password = $2, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE email = $1;