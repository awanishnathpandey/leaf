-- name: GetNotification :one
SELECT * FROM notifications
WHERE id = $1 LIMIT 1;

-- name: ListNotifications :many
SELECT * FROM notifications;

-- name: GetNotificationTemplate :one
SELECT * FROM notification_templates
WHERE id = $1 LIMIT 1;

-- name: ListNotificationTemplates :many
SELECT * FROM notification_templates;

-- name: CreateNotificationTemplate :one
INSERT INTO notification_templates (title, body, description, response_options, created_by, updated_by) 
VALUES ($1, $2, $3, $4, $5, $5) 
RETURNING *;

-- name: CreateNotification :one
INSERT INTO notifications (notification_type, record_key_id, payload, start_time_at, end_time_at, is_push_notification, status, group_ids, user_ids, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;