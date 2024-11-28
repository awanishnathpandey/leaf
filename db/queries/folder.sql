-- name: GetFolder :one
SELECT * FROM folders
WHERE id = $1 LIMIT 1;

-- name: ListFolders :many
SELECT * FROM folders
ORDER BY name;

-- name: CreateFolder :one
INSERT INTO folders (
  name,
  slug,
  description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateFolder :exec
UPDATE folders
  set name = $2, slug = $3, description = $4, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;

-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1;

-- name: GetFoldersByIDs :many
SELECT id FROM folders
WHERE id = ANY($1::bigint[]);

-- name: DeleteFoldersByIDs :exec
DELETE FROM folders
WHERE id = ANY($1::bigint[]);

-- name: PaginatedFolders :many
SELECT * FROM folders
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN slug 
        WHEN sqlc.narg(sort_field) = 'DESCRPITION' AND sqlc.narg(sort_order) = 'ASC' THEN description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN slug 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: PaginatedFoldersCount :one
SELECT COUNT(*) FROM folders
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR description ILIKE '%' || sqlc.narg(description_filter) || '%');