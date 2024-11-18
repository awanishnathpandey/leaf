-- name: CreateGroup :one
INSERT INTO groups (name, description) 
VALUES ($1, $2) 
RETURNING id, name, description, created_at, updated_at;

-- name: UpdateGroup :exec
UPDATE groups 
SET name = $1, description = $2, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $3;

-- name: DeleteGroup :exec
DELETE FROM groups 
WHERE id = $1;

-- name: AddUserToGroup :exec
INSERT INTO group_users (group_id, user_id, created_at, updated_at) 
VALUES ($1, $2, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())) 
ON CONFLICT DO NOTHING;

-- name: RemoveUserFromGroup :exec
DELETE FROM group_users 
WHERE group_id = $1 AND user_id = $2;

-- name: AddFolderToGroup :exec
INSERT INTO group_folders (group_id, folder_id, created_at, updated_at) 
VALUES ($1, $2, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())) 
ON CONFLICT DO NOTHING;

-- name: RemoveFolderFromGroup :exec
DELETE FROM group_folders 
WHERE group_id = $1 AND folder_id = $2;

-- name: AddFileToGroup :exec
INSERT INTO group_files (group_id, file_id, created_at, updated_at) 
VALUES ($1, $2, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())) 
ON CONFLICT DO NOTHING;

-- name: RemoveFileFromGroup :exec
DELETE FROM group_files 
WHERE group_id = $1 AND file_id = $2;

-- name: GetGroup :one
SELECT id, name, description, created_at, updated_at 
FROM groups 
WHERE id = $1;

-- name: ListGroups :many
SELECT * FROM groups
ORDER BY name;

-- name: GetUsersByGroupID :many
SELECT u.id, u.name, u.email, u.email_verified_at, u.last_seen_at, u.created_at, u.updated_at, u.deleted_at
FROM users u
JOIN group_users gu ON u.id = gu.user_id
WHERE gu.group_id = $1;

-- name: GetFoldersByGroupID :many
SELECT f.*
FROM folders f
JOIN group_folders gf ON f.id = gf.folder_id
WHERE gf.group_id = $1;

-- name: GetFilesByGroupID :many
SELECT f.*
FROM files f
JOIN group_files gf ON f.id = gf.file_id
WHERE gf.group_id = $1;
