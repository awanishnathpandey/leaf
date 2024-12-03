-- name: CreateFile :one
INSERT INTO files (name, slug, file_path, folder_id, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $5)
RETURNING *;

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
SET name = $2, slug = $3, file_path = $4, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $5
WHERE id = $1
RETURNING *;

-- name: DeleteFile :exec
DELETE FROM files
WHERE id = $1;

-- name: GetFilesByIDs :many
SELECT id FROM files
WHERE id = ANY($1::bigint[]);

-- name: DeleteFilesByIDs :exec
DELETE FROM files
WHERE id = ANY($1::bigint[]);

-- name: GetFilesByFolderID :many
SELECT * FROM files
WHERE folder_id = $1;

-- name: PaginatedFiles :many
SELECT * FROM files
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN slug 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN slug 
    END DESC
LIMIT $1
OFFSET $2;

-- name: PaginatedFilesCount :one
SELECT COUNT(*) FROM files
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%');

-- name: GetPaginatedFilesByFolderID :many
SELECT * FROM files WHERE 
    folder_id = sqlc.narg(folder_id)  -- Filter by folder_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN slug 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN slug 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedFilesByFolderIDCount :one
SELECT COUNT(*) FROM files WHERE 
    folder_id = sqlc.narg(folder_id)  -- Filter by folder_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%');