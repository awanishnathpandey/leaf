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

-- name: GetFilesAndFoldersByUserBB :many
SELECT
    folder.id AS folder_id,
    folder.name AS folder_name,
    folder.slug AS folder_slug,
    folder.description AS folder_description,
    folder.created_at AS folder_created_at,
    folder.updated_at AS folder_updated_at,
    f.id AS file_id,
    f.name AS file_name,
    f.slug AS file_slug,
    f.file_path AS file_path,
    f.file_type AS file_type,
    f.file_bytes AS file_bytes,
    f.auto_download AS file_auto_download,
    f.folder_id AS file_folder_id,
    f.created_at AS file_created_at,
    f.updated_at AS file_updated_at
FROM files f
-- Join the folders table to get folder details
JOIN folders folder ON folder.id = f.folder_id
-- Join group_users to get the groups the user belongs to
JOIN group_users gu ON gu.group_id IN (
    SELECT group_id FROM group_users WHERE group_users.user_id = $1
)
-- Join group_files to get the files directly associated with the user's groups
LEFT JOIN group_files gf ON gf.file_id = f.id
-- Join group_folders to get the folders associated with the user's groups via the pivot table
LEFT JOIN group_folders gfo ON gfo.folder_id = folder.id
WHERE gu.user_id = $1  -- The user ID is passed as the first parameter
AND f.file_type = $2  -- The file type (e.g., 'document') is passed as the second parameter
ORDER BY folder.name, f.name;



-- name: GetFilesAndFoldersByUser :many
WITH epoch_7_days_ago AS (
    SELECT NOW() - INTERVAL '7 days' AS threshold
)
SELECT
    folder.id AS folder_id,
    folder.name AS folder_name,
    folder.slug AS folder_slug,
    folder.description AS folder_description,
    folder.created_at AS folder_created_at,
    folder.updated_at AS folder_updated_at,
    -- Adding "hasNewFile" based on files updated in the last 7 days
    CASE 
        WHEN EXISTS (
            SELECT 1
            FROM files f, epoch_7_days_ago e
            WHERE f.folder_id = folder.id
            AND f.file_type = $2  -- File type condition
            AND to_timestamp(f.updated_at) > e.threshold  -- Convert bigint to timestamp
        ) THEN true
        ELSE false
    END AS hasNewFile,
    f.id AS file_id,
    f.name AS file_name,
    f.slug AS file_slug,
    f.file_path AS file_path,
    f.file_type AS file_type,
    f.file_bytes AS file_bytes,
    f.auto_download AS file_auto_download,
    f.folder_id AS file_folder_id,
    f.created_at AS file_created_at,
    f.updated_at AS file_updated_at,
    -- Adding "isNew" based on file update within the last 7 days
    CASE 
        WHEN to_timestamp(f.updated_at) > (SELECT threshold FROM epoch_7_days_ago) THEN true
        ELSE false
    END AS isNew
FROM files f
-- Join the folders table to get folder details
JOIN folders folder ON folder.id = f.folder_id
-- Join group_users to get the groups the user belongs to
JOIN group_users gu ON gu.group_id IN (
    SELECT group_id FROM group_users WHERE group_users.user_id = $1
)
-- Join group_files to get the files directly associated with the user's groups
LEFT JOIN group_files gf ON gf.file_id = f.id
-- Join group_folders to get the folders associated with the user's groups via the pivot table
LEFT JOIN group_folders gfo ON gfo.folder_id = folder.id
WHERE gu.user_id = $1  -- The user ID is passed as the first parameter
AND f.file_type = $2  -- The file type (e.g., 'document') is passed as the second parameter
ORDER BY folder.name, f.name;
