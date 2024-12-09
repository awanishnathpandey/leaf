-- name: GetAppConfigByKey :one
SELECT * FROM app_config
WHERE config_key = $1
LIMIT 1;

-- name: UpdateAppConfigByKey :exec
UPDATE app_config
SET config_data = $2, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $3
WHERE config_key = $1;