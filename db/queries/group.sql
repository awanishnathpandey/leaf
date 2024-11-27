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

-- name: GetGroupsByFolderID :many
SELECT g.*
FROM groups g
JOIN group_folders gf ON g.id = gf.group_id
WHERE gf.folder_id = $1;

-- name: GetGroupsByUserID :many
SELECT g.*
FROM groups g
JOIN group_users gu ON g.id = gu.group_id
WHERE gu.user_id = $1;

-- name: GetGroupsByFileID :many
SELECT g.*
FROM groups g
JOIN group_files gf ON g.id = gf.group_id
WHERE gf.file_id = $1;

-- name: PaginatedGroups :many
SELECT * FROM groups
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: PaginatedGroupsCount :one
SELECT COUNT(*) FROM groups
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR description ILIKE '%' || sqlc.narg(description_filter) || '%');

-- name: GetPaginatedUsersByGroupID :many
SELECT 
    u.id, 
    u.name, 
    u.email, 
    u.email_verified_at, 
    u.last_seen_at, 
    u.created_at, 
    u.updated_at, 
    u.deleted_at
FROM users u
JOIN group_users gu ON u.id = gu.user_id
WHERE 
    gu.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR u.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(email_filter), '') = '' OR u.email ILIKE '%' || sqlc.narg(email_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN u.name 
        WHEN sqlc.narg(sort_field) = 'EMAIL' AND sqlc.narg(sort_order) = 'ASC' THEN u.email 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN u.name 
        WHEN sqlc.narg(sort_field) = 'EMAIL' AND sqlc.narg(sort_order) = 'DESC' THEN u.email 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedUsersByGroupIDCount :one
SELECT COUNT(*) FROM users u
JOIN group_users gu ON u.id = gu.user_id
WHERE 
    gu.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR u.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(email_filter), '') = '' OR u.email ILIKE '%' || sqlc.narg(email_filter) || '%');


-- name: GetPaginatedFilesByGroupID :many
SELECT * FROM files f
JOIN group_files gf ON f.id = gf.file_id
WHERE 
    gf.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR f.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR f.slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN f.name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN f.slug 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN f.name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN f.slug 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedFilesByGroupIDCount :one
SELECT COUNT(*) FROM files f
JOIN group_files gf ON f.id = gf.file_id
WHERE 
    gf.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR f.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR f.slug ILIKE '%' || sqlc.narg(slug_filter) || '%');

-- name: GetPaginatedFoldersByGroupID :many
SELECT * FROM folders f
JOIN group_folders gf ON f.id = gf.folder_id
WHERE 
    gf.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR f.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR f.slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR f.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN f.name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN f.slug 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN f.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN f.name 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN f.slug 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN f.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedFoldersByGroupIDCount :one
SELECT COUNT(*) FROM folders f
JOIN group_folders gf ON f.id = gf.folder_id
WHERE 
    gf.group_id = sqlc.narg(group_id)  -- Filter by group_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR f.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR f.slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR f.description ILIKE '%' || sqlc.narg(description_filter) || '%');

-- name: GetPaginatedGroupsByUserID :many
SELECT * FROM groups g
JOIN group_users gu ON g.id = gu.group_id
WHERE 
    gu.user_id = sqlc.narg(user_id)  -- Filter by user_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN g.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN g.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedGroupsByUserIDCount :one
SELECT COUNT(*) FROM groups g
JOIN group_users gu ON g.id = gu.group_id
WHERE 
    gu.user_id = sqlc.narg(user_id)  -- Filter by user_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%');

-- name: GetPaginatedGroupsByFolderID :many
SELECT * FROM groups g
JOIN group_folders gf ON g.id = gf.group_id
WHERE 
    gf.folder_id = sqlc.narg(folder_id)  -- Filter by folder_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN g.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN g.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedGroupsByFolderIDCount :one
SELECT COUNT(*) FROM groups g
JOIN group_folders gf ON g.id = gf.group_id
WHERE 
    gf.folder_id = sqlc.narg(folder_id)  -- Filter by folder_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%');

-- name: GetPaginatedGroupsByFileID :many
SELECT * FROM groups g
JOIN group_files gf ON g.id = gf.group_id
WHERE 
    gf.file_id = sqlc.narg(file_id)  -- Filter by user_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN g.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN g.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN g.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedGroupsByFileIDCount :one
SELECT COUNT(*) FROM groups g
JOIN group_files gf ON g.id = gf.group_id
WHERE 
    gf.file_id = sqlc.narg(file_id)  -- Filter by user_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR g.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR g.description ILIKE '%' || sqlc.narg(description_filter) || '%');