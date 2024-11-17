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
INSERT INTO group_users (group_id, user_id) 
VALUES ($1, $2) 
ON CONFLICT DO NOTHING;

-- name: RemoveUserFromGroup :exec
DELETE FROM group_users 
WHERE group_id = $1 AND user_id = $2;

-- name: AddFolderToGroup :exec
INSERT INTO group_folders (group_id, folder_id) 
VALUES ($1, $2) 
ON CONFLICT DO NOTHING;

-- name: RemoveFolderFromGroup :exec
DELETE FROM group_folders 
WHERE group_id = $1 AND folder_id = $2;

-- name: AddFileToGroup :exec
INSERT INTO group_files (group_id, file_id) 
VALUES ($1, $2) 
ON CONFLICT DO NOTHING;

-- name: RemoveFileFromGroup :exec
DELETE FROM group_files 
WHERE group_id = $1 AND file_id = $2;

-- name: GetGroupByID :one
SELECT id, name, description, created_at, updated_at 
FROM groups 
WHERE id = $1;

-- name: GetGroups :many
SELECT id, name, description, created_at, updated_at 
FROM groups;

-- name: GetGroupUsers :many
SELECT users.id, users.name, users.email 
FROM users 
JOIN group_users ON group_users.user_id = users.id 
WHERE group_users.group_id = $1;

-- name: GetGroupFolders :many
SELECT folders.id, folders.name 
FROM folders 
JOIN group_folders ON group_folders.folder_id = folders.id 
WHERE group_folders.group_id = $1;

-- name: GetGroupFiles :many
SELECT files.id, files.name 
FROM files 
JOIN group_files ON group_files.file_id = files.id 
WHERE group_files.group_id = $1;
