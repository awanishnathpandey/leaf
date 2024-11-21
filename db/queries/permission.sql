-- name: GetUserPermissions :many
SELECT p.name
        FROM permissions p
        JOIN role_permissions rp ON rp.permission_id = p.id
        JOIN user_roles ur ON ur.role_id = rp.role_id
        WHERE ur.user_id = $1;

-- name: CreateRole :one
INSERT INTO roles (name, description) 
VALUES ($1, $2) 
RETURNING id, name, description, created_at, updated_at;

-- name: CreatePermission :one
INSERT INTO roles (name, description) 
VALUES ($1, $2) 
RETURNING id, name, description, created_at, updated_at;

-- name: UpdateRole :exec
UPDATE roles 
SET name = $1, description = $2, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $3;

-- name: UpdatePermission :exec
UPDATE permissions 
SET name = $1, description = $2, updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $3;

-- name: DeleteRole :exec
DELETE FROM roles 
WHERE id = $1;

-- name: DeletePermission :exec
DELETE FROM permissions 
WHERE id = $1;

-- name: AddRoleToUser :exec
INSERT INTO user_roles (role_id, user_id, created_at, updated_at) 
VALUES ($1, $2, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())) 
ON CONFLICT DO NOTHING;

-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles 
WHERE role_id = $1 AND user_id = $2;

-- name: AddPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at) 
VALUES ($1, $2, EXTRACT(EPOCH FROM NOW()), EXTRACT(EPOCH FROM NOW())) 
ON CONFLICT DO NOTHING;

-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions 
WHERE role_id = $1 AND permission_id = $2;

-- name: GetRole :one
SELECT id, name, description, created_at, updated_at 
FROM roles 
WHERE id = $1;

-- name: GetPermission :one
SELECT id, name, description, created_at, updated_at 
FROM permissions 
WHERE id = $1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY name;

-- name: ListPermissions :many
SELECT * FROM permissions
ORDER BY name;

-- name: GetUsersByRoleID :many
SELECT u.id, u.name, u.email, u.email_verified_at, u.last_seen_at, u.created_at, u.updated_at, u.deleted_at
FROM users u
JOIN user_roles ur ON u.id = ur.user_id
WHERE ur.role_id = $1;

-- name: GetPermissionsByRoleID :many
SELECT p.*
FROM permissions p
JOIN role_permissions rp ON p.id = rp.permission_id
WHERE rp.role_id = $1;


-- name: GetRolesByUserID :many
SELECT r.*
FROM roles r
JOIN user_roles ur ON r.id = ur.role_id
WHERE ur.user_id = $1;

-- name: GetRolesByPermissionID :many
SELECT r.*
FROM roles r
JOIN role_permissions rp ON r.id = rp.role_id
WHERE rp.permission_id = $1;

-- name: PaginatedRoles :many
SELECT * FROM roles
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

-- name: PaginatedPermissions :many
SELECT * FROM permissions
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


-- name: GetPaginatedRolesByPermissionID :many
SELECT * FROM roles r
JOIN role_permissions rp ON r.id = rp.role_id
WHERE 
    rp.permission_id = sqlc.narg(permission_id)  -- Filter by permission_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR r.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR r.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN r.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN r.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN r.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN r.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedPermissionsByRoleID :many
SELECT * FROM permissions p
JOIN role_permissions rp ON p.id = rp.permission_id
WHERE 
    rp.role_id = sqlc.narg(role_id)  -- Filter by permission_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR p.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR p.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN p.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN p.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN p.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN p.description 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedUsersByRoleID :many
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
JOIN user_roles ur ON u.id = ur.user_id
WHERE 
    ur.role_id = sqlc.narg(role_id)  -- Filter by group_id
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

-- name: GetPaginatedRolesByUserID :many
SELECT * FROM roles r
JOIN user_roles ur ON r.id = ur.role_id
WHERE 
    ur.user_id = sqlc.narg(user_id)  -- Filter by user_id
    AND (coalesce(sqlc.narg(name_filter), '') = '' OR r.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR r.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN r.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN r.description 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN r.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN r.description 
    END DESC
LIMIT $1
OFFSET $2;