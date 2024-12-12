// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: notification.sql

package generated

import (
	"context"
)

const CreateNotification = `-- name: CreateNotification :one
INSERT INTO notifications (notification_type, record_key_id, payload, start_time_at, end_time_at, is_push_notification, status, group_ids, user_ids, created_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, notification_type, record_key_id, payload, start_time_at, end_time_at, is_push_notification, status, group_ids, user_ids, created_at, created_by
`

type CreateNotificationParams struct {
	NotificationType   string  `db:"notification_type" json:"notification_type"`
	RecordKeyID        int64   `db:"record_key_id" json:"record_key_id"`
	Payload            []byte  `db:"payload" json:"payload"`
	StartTimeAt        int64   `db:"start_time_at" json:"start_time_at"`
	EndTimeAt          int64   `db:"end_time_at" json:"end_time_at"`
	IsPushNotification bool    `db:"is_push_notification" json:"is_push_notification"`
	Status             string  `db:"status" json:"status"`
	GroupIds           []int64 `db:"group_ids" json:"group_ids"`
	UserIds            []int64 `db:"user_ids" json:"user_ids"`
	CreatedBy          string  `db:"created_by" json:"created_by"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error) {
	row := q.db.QueryRow(ctx, CreateNotification,
		arg.NotificationType,
		arg.RecordKeyID,
		arg.Payload,
		arg.StartTimeAt,
		arg.EndTimeAt,
		arg.IsPushNotification,
		arg.Status,
		arg.GroupIds,
		arg.UserIds,
		arg.CreatedBy,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.NotificationType,
		&i.RecordKeyID,
		&i.Payload,
		&i.StartTimeAt,
		&i.EndTimeAt,
		&i.IsPushNotification,
		&i.Status,
		&i.GroupIds,
		&i.UserIds,
		&i.CreatedAt,
		&i.CreatedBy,
	)
	return i, err
}

const CreateNotificationTemplate = `-- name: CreateNotificationTemplate :one
INSERT INTO notification_templates (title, body, description, response_options, created_by, updated_by) 
VALUES ($1, $2, $3, $4, $5, $5) 
RETURNING id, title, body, description, response_options, created_at, updated_at, created_by, updated_by
`

type CreateNotificationTemplateParams struct {
	Title           string   `db:"title" json:"title"`
	Body            string   `db:"body" json:"body"`
	Description     string   `db:"description" json:"description"`
	ResponseOptions []string `db:"response_options" json:"response_options"`
	CreatedBy       string   `db:"created_by" json:"created_by"`
}

func (q *Queries) CreateNotificationTemplate(ctx context.Context, arg CreateNotificationTemplateParams) (NotificationTemplate, error) {
	row := q.db.QueryRow(ctx, CreateNotificationTemplate,
		arg.Title,
		arg.Body,
		arg.Description,
		arg.ResponseOptions,
		arg.CreatedBy,
	)
	var i NotificationTemplate
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Description,
		&i.ResponseOptions,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const GetNotification = `-- name: GetNotification :one
SELECT id, notification_type, record_key_id, payload, start_time_at, end_time_at, is_push_notification, status, group_ids, user_ids, created_at, created_by FROM notifications
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetNotification(ctx context.Context, id int64) (Notification, error) {
	row := q.db.QueryRow(ctx, GetNotification, id)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.NotificationType,
		&i.RecordKeyID,
		&i.Payload,
		&i.StartTimeAt,
		&i.EndTimeAt,
		&i.IsPushNotification,
		&i.Status,
		&i.GroupIds,
		&i.UserIds,
		&i.CreatedAt,
		&i.CreatedBy,
	)
	return i, err
}

const GetNotificationTemplate = `-- name: GetNotificationTemplate :one
SELECT id, title, body, description, response_options, created_at, updated_at, created_by, updated_by FROM notification_templates
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetNotificationTemplate(ctx context.Context, id int64) (NotificationTemplate, error) {
	row := q.db.QueryRow(ctx, GetNotificationTemplate, id)
	var i NotificationTemplate
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Description,
		&i.ResponseOptions,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const ListNotificationTemplates = `-- name: ListNotificationTemplates :many
SELECT id, title, body, description, response_options, created_at, updated_at, created_by, updated_by FROM notification_templates
`

func (q *Queries) ListNotificationTemplates(ctx context.Context) ([]NotificationTemplate, error) {
	rows, err := q.db.Query(ctx, ListNotificationTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NotificationTemplate{}
	for rows.Next() {
		var i NotificationTemplate
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.Description,
			&i.ResponseOptions,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListNotifications = `-- name: ListNotifications :many
SELECT id, notification_type, record_key_id, payload, start_time_at, end_time_at, is_push_notification, status, group_ids, user_ids, created_at, created_by FROM notifications
`

func (q *Queries) ListNotifications(ctx context.Context) ([]Notification, error) {
	rows, err := q.db.Query(ctx, ListNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.NotificationType,
			&i.RecordKeyID,
			&i.Payload,
			&i.StartTimeAt,
			&i.EndTimeAt,
			&i.IsPushNotification,
			&i.Status,
			&i.GroupIds,
			&i.UserIds,
			&i.CreatedAt,
			&i.CreatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}