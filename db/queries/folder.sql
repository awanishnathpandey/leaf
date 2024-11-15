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
  set name = $2, slug = $3, description = $4, updated_at = NOW()
WHERE id = $1;

-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1;