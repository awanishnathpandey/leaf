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
  description,
  created_by,
  updated_by
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateFolder :exec
UPDATE folders
  set name = $2, slug = $3, description = $4, created_by = $5, updated_by = $6
WHERE id = $1;

-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1;