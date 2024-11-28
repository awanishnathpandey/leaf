-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserID :one
SELECT id FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password, created_by, updated_by
) VALUES (
  $1, $2, $3, $4, $4
)
RETURNING id, name, email, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by;

-- name: UpdateUser :one
UPDATE users
  set name = $2, email = $3, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $4
WHERE id = $1
RETURNING id, name, email, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by;

-- name: UpdateUserEmailVerifiedAt :exec
UPDATE users
  set email_verified_at = EXTRACT(EPOCH FROM NOW()), updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: UpdateUserLastSeenAt :exec
UPDATE users
  set last_seen_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUsersByIDs :many
SELECT id FROM users
WHERE id = ANY($1::bigint[]);

-- name: DeleteUsersByIDs :exec
DELETE FROM users
WHERE id = ANY($1::bigint[]);

-- name: PaginatedUsers :many
SELECT 
    id, 
    name, 
    email, 
    password, 
    email_verified_at, 
    last_seen_at, 
    created_at, 
    updated_at, 
    deleted_at,
    created_by,
    updated_by
FROM users
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(email_filter), '') = '' OR email ILIKE '%' || sqlc.narg(email_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN name 
        WHEN sqlc.narg(sort_field) = 'EMAIL' AND sqlc.narg(sort_order) = 'ASC' THEN email 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN name 
        WHEN sqlc.narg(sort_field) = 'EMAIL' AND sqlc.narg(sort_order) = 'DESC' THEN email 
    END DESC
LIMIT $1
OFFSET $2;


-- name: PaginatedUsersCount :one
SELECT COUNT(*)
FROM users
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(email_filter), '') = '' OR email ILIKE '%' || sqlc.narg(email_filter) || '%');