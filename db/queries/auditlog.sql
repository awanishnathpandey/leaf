-- name: CreateAuditLog :exec
INSERT INTO audit_logs (
  table_name,actor, action, ip_address, record_key, description, timestamp
) VALUES (
  $1, $2, $3, $4, $5, $6, EXTRACT(EPOCH FROM NOW())
);

-- name: GetAuditLog :one
SELECT * FROM audit_logs
WHERE id = $1 LIMIT 1;

-- name: DeleteAuditLog :exec
DELETE FROM audit_logs
WHERE id = $1;

-- name: GetAuditLogsByIDs :many
SELECT id FROM audit_logs
WHERE id = ANY($1::bigint[]);

-- name: DeleteAuditLogsByIDs :exec
DELETE FROM audit_logs
WHERE id = ANY($1::bigint[]);

-- name: GetPaginatedAuditLogsByUserEmail :many
SELECT * FROM audit_logs al
JOIN users u ON al.actor = u.email
WHERE 
    u.email = sqlc.narg(user_email)  -- Filter by user_id
    AND (coalesce(sqlc.narg(table_name_filter), '') = '' OR al.table_name ILIKE '%' || sqlc.narg(table_name_filter) || '%')
    AND (coalesce(sqlc.narg(actor_filter), '') = '' OR al.actor ILIKE '%' || sqlc.narg(actor_filter) || '%')
    AND (coalesce(sqlc.narg(action_filter), '') = '' OR al.action ILIKE '%' || sqlc.narg(action_filter) || '%')
    AND (coalesce(sqlc.narg(ip_address_filter), '') = '' OR al.ip_address ILIKE '%' || sqlc.narg(ip_address_filter) || '%')
    AND (coalesce(sqlc.narg(record_key_filter), '') = '' OR al.record_key ILIKE '%' || sqlc.narg(record_key_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR al.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'TABLENAME' AND sqlc.narg(sort_order) = 'ASC' THEN al.table_name 
        WHEN sqlc.narg(sort_field) = 'ACTOR' AND sqlc.narg(sort_order) = 'ASC' THEN al.actor 
        WHEN sqlc.narg(sort_field) = 'ACTION' AND sqlc.narg(sort_order) = 'ASC' THEN al.action 
        WHEN sqlc.narg(sort_field) = 'IPADDRESS' AND sqlc.narg(sort_order) = 'ASC' THEN al.ip_address 
        WHEN sqlc.narg(sort_field) = 'RECORDKEY' AND sqlc.narg(sort_order) = 'ASC' THEN al.record_key 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN al.description
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TIMESTAMP' AND sqlc.narg(sort_order) = 'ASC' THEN al.timestamp
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TABLENAME' AND sqlc.narg(sort_order) = 'DESC' THEN al.table_name 
        WHEN sqlc.narg(sort_field) = 'ACTOR' AND sqlc.narg(sort_order) = 'DESC' THEN al.actor 
        WHEN sqlc.narg(sort_field) = 'ACTION' AND sqlc.narg(sort_order) = 'DESC' THEN al.action 
        WHEN sqlc.narg(sort_field) = 'IPADDRESS' AND sqlc.narg(sort_order) = 'DESC' THEN al.ip_address 
        WHEN sqlc.narg(sort_field) = 'RECORDKEY' AND sqlc.narg(sort_order) = 'DESC' THEN al.record_key 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN al.description
    END DESC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TIMESTAMP' AND sqlc.narg(sort_order) = 'DESC' THEN al.timestamp 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedAuditLogsByUserEmailCount :one
SELECT COUNT(*) FROM audit_logs al
JOIN users u ON al.actor = u.email
WHERE 
    u.email = sqlc.narg(user_email)  -- Filter by user_id
    AND (coalesce(sqlc.narg(table_name_filter), '') = '' OR al.table_name ILIKE '%' || sqlc.narg(table_name_filter) || '%')
    AND (coalesce(sqlc.narg(actor_filter), '') = '' OR al.actor ILIKE '%' || sqlc.narg(actor_filter) || '%')
    AND (coalesce(sqlc.narg(action_filter), '') = '' OR al.action ILIKE '%' || sqlc.narg(action_filter) || '%')
    AND (coalesce(sqlc.narg(ip_address_filter), '') = '' OR al.ip_address ILIKE '%' || sqlc.narg(ip_address_filter) || '%')
    AND (coalesce(sqlc.narg(record_key_filter), '') = '' OR al.record_key ILIKE '%' || sqlc.narg(record_key_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR al.description ILIKE '%' || sqlc.narg(description_filter) || '%');

-- name: GetPaginatedAuditLogs :many
SELECT * FROM audit_logs al
WHERE 
    (coalesce(sqlc.narg(table_name_filter), '') = '' OR al.table_name ILIKE '%' || sqlc.narg(table_name_filter) || '%')
    AND (coalesce(sqlc.narg(actor_filter), '') = '' OR al.actor ILIKE '%' || sqlc.narg(actor_filter) || '%')
    AND (coalesce(sqlc.narg(action_filter), '') = '' OR al.action ILIKE '%' || sqlc.narg(action_filter) || '%')
    AND (coalesce(sqlc.narg(ip_address_filter), '') = '' OR al.ip_address ILIKE '%' || sqlc.narg(ip_address_filter) || '%')
    AND (coalesce(sqlc.narg(record_key_filter), '') = '' OR al.record_key ILIKE '%' || sqlc.narg(record_key_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR al.description ILIKE '%' || sqlc.narg(description_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'TABLENAME' AND sqlc.narg(sort_order) = 'ASC' THEN al.table_name 
        WHEN sqlc.narg(sort_field) = 'ACTOR' AND sqlc.narg(sort_order) = 'ASC' THEN al.actor 
        WHEN sqlc.narg(sort_field) = 'ACTION' AND sqlc.narg(sort_order) = 'ASC' THEN al.action 
        WHEN sqlc.narg(sort_field) = 'IPADDRESS' AND sqlc.narg(sort_order) = 'ASC' THEN al.ip_address 
        WHEN sqlc.narg(sort_field) = 'RECORDKEY' AND sqlc.narg(sort_order) = 'ASC' THEN al.record_key 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN al.description
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TIMESTAMP' AND sqlc.narg(sort_order) = 'ASC' THEN al.timestamp
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TABLENAME' AND sqlc.narg(sort_order) = 'DESC' THEN al.table_name 
        WHEN sqlc.narg(sort_field) = 'ACTOR' AND sqlc.narg(sort_order) = 'DESC' THEN al.actor 
        WHEN sqlc.narg(sort_field) = 'ACTION' AND sqlc.narg(sort_order) = 'DESC' THEN al.action 
        WHEN sqlc.narg(sort_field) = 'IPADDRESS' AND sqlc.narg(sort_order) = 'DESC' THEN al.ip_address 
        WHEN sqlc.narg(sort_field) = 'RECORDKEY' AND sqlc.narg(sort_order) = 'DESC' THEN al.record_key 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN al.description
    END DESC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'TIMESTAMP' AND sqlc.narg(sort_order) = 'DESC' THEN al.timestamp
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedAuditLogsCount :one
SELECT COUNT(*) FROM audit_logs al
WHERE 
    (coalesce(sqlc.narg(table_name_filter), '') = '' OR al.table_name ILIKE '%' || sqlc.narg(table_name_filter) || '%')
    AND (coalesce(sqlc.narg(actor_filter), '') = '' OR al.actor ILIKE '%' || sqlc.narg(actor_filter) || '%')
    AND (coalesce(sqlc.narg(action_filter), '') = '' OR al.action ILIKE '%' || sqlc.narg(action_filter) || '%')
    AND (coalesce(sqlc.narg(ip_address_filter), '') = '' OR al.ip_address ILIKE '%' || sqlc.narg(ip_address_filter) || '%')
    AND (coalesce(sqlc.narg(record_key_filter), '') = '' OR al.record_key ILIKE '%' || sqlc.narg(record_key_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR al.description ILIKE '%' || sqlc.narg(description_filter) || '%');