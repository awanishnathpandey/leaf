-- name: CreateAuditLog :exec
INSERT INTO audit_logs (
  table_name,actor, action, ip_address, record_key, description, timestamp
) VALUES (
  $1, $2, $3, $4, $5, $6, EXTRACT(EPOCH FROM NOW())
);