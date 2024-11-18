-- name: CreateFile :one
INSERT INTO files (name, slug, url, folder_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, slug, url, folder_id, created_at, updated_at;

-- name: GetFile :one
SELECT * FROM files
WHERE id = $1;

-- name: ListFiles :many
SELECT * FROM files
ORDER BY name;

-- name: GetFilesByFolder :many
SELECT * FROM files
WHERE folder_id = $1;

-- name: UpdateFile :one
UPDATE files
SET name = $2, slug = $3, url = $4, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1
RETURNING id, name, slug, url, folder_id, created_at, updated_at;

-- name: DeleteFile :exec
DELETE FROM files
WHERE id = $1;

-- name: GetFilesByFolderID :many
SELECT * FROM files
WHERE folder_id = $1;