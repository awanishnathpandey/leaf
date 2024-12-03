-- name: CreatePasswordReset :one
INSERT INTO password_resets (user_id, reset_token, created_by, updated_by) 
VALUES ($1, $2, $3, $3) 
RETURNING *;

-- name: GetPasswordResetbyUserID :one
SELECT * FROM password_resets
WHERE user_id = $1 LIMIT 1;

-- name: DeletePasswordResetbyUserID :exec
DELETE FROM password_resets 
WHERE user_id = $1;