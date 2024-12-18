// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auditlog.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const CreateAuditLog = `-- name: CreateAuditLog :exec
INSERT INTO audit_logs (
  table_name,actor, action, ip_address, record_key, description, timestamp
) VALUES (
  $1, $2, $3, $4, $5, $6, EXTRACT(EPOCH FROM NOW())
)
`

type CreateAuditLogParams struct {
	TableName   string `db:"table_name" json:"table_name"`
	Actor       string `db:"actor" json:"actor"`
	Action      string `db:"action" json:"action"`
	IpAddress   string `db:"ip_address" json:"ip_address"`
	RecordKey   string `db:"record_key" json:"record_key"`
	Description string `db:"description" json:"description"`
}

func (q *Queries) CreateAuditLog(ctx context.Context, arg CreateAuditLogParams) error {
	_, err := q.db.Exec(ctx, CreateAuditLog,
		arg.TableName,
		arg.Actor,
		arg.Action,
		arg.IpAddress,
		arg.RecordKey,
		arg.Description,
	)
	return err
}

const DeleteAuditLog = `-- name: DeleteAuditLog :exec
DELETE FROM audit_logs
WHERE id = $1
`

func (q *Queries) DeleteAuditLog(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, DeleteAuditLog, id)
	return err
}

const DeleteAuditLogsByIDs = `-- name: DeleteAuditLogsByIDs :exec
DELETE FROM audit_logs
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteAuditLogsByIDs(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, DeleteAuditLogsByIDs, dollar_1)
	return err
}

const GetAuditLog = `-- name: GetAuditLog :one
SELECT id, table_name, actor, action, ip_address, record_key, description, timestamp FROM audit_logs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuditLog(ctx context.Context, id int64) (AuditLog, error) {
	row := q.db.QueryRow(ctx, GetAuditLog, id)
	var i AuditLog
	err := row.Scan(
		&i.ID,
		&i.TableName,
		&i.Actor,
		&i.Action,
		&i.IpAddress,
		&i.RecordKey,
		&i.Description,
		&i.Timestamp,
	)
	return i, err
}

const GetAuditLogsByIDs = `-- name: GetAuditLogsByIDs :many
SELECT id FROM audit_logs
WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetAuditLogsByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, GetAuditLogsByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetPaginatedAuditLogs = `-- name: GetPaginatedAuditLogs :many
SELECT id, table_name, actor, action, ip_address, record_key, description, timestamp FROM audit_logs al
WHERE 
    (coalesce($3, '') = '' OR al.table_name ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR al.actor ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR al.action ILIKE '%' || $5 || '%')
    AND (coalesce($6, '') = '' OR al.ip_address ILIKE '%' || $6 || '%')
    AND (coalesce($7, '') = '' OR al.record_key ILIKE '%' || $7 || '%')
    AND (coalesce($8, '') = '' OR al.description ILIKE '%' || $8 || '%')
ORDER BY 
    CASE 
        WHEN $9 = 'TABLENAME' AND $10 = 'ASC' THEN al.table_name 
        WHEN $9 = 'ACTOR' AND $10 = 'ASC' THEN al.actor 
        WHEN $9 = 'ACTION' AND $10 = 'ASC' THEN al.action 
        WHEN $9 = 'IPADDRESS' AND $10 = 'ASC' THEN al.ip_address 
        WHEN $9 = 'RECORDKEY' AND $10 = 'ASC' THEN al.record_key 
        WHEN $9 = 'DESCRIPTION' AND $10 = 'ASC' THEN al.description
    END ASC,
    CASE 
        WHEN $9 = 'TIMESTAMP' AND $10 = 'ASC' THEN al.timestamp
    END ASC,
    CASE 
        WHEN $9 = 'TABLENAME' AND $10 = 'DESC' THEN al.table_name 
        WHEN $9 = 'ACTOR' AND $10 = 'DESC' THEN al.actor 
        WHEN $9 = 'ACTION' AND $10 = 'DESC' THEN al.action 
        WHEN $9 = 'IPADDRESS' AND $10 = 'DESC' THEN al.ip_address 
        WHEN $9 = 'RECORDKEY' AND $10 = 'DESC' THEN al.record_key 
        WHEN $9 = 'DESCRIPTION' AND $10 = 'DESC' THEN al.description
    END DESC,
    CASE 
        WHEN $9 = 'TIMESTAMP' AND $10 = 'DESC' THEN al.timestamp
    END DESC
LIMIT $1
OFFSET $2
`

type GetPaginatedAuditLogsParams struct {
	Limit             int32       `db:"limit" json:"limit"`
	Offset            int32       `db:"offset" json:"offset"`
	TableNameFilter   interface{} `db:"table_name_filter" json:"table_name_filter"`
	ActorFilter       interface{} `db:"actor_filter" json:"actor_filter"`
	ActionFilter      interface{} `db:"action_filter" json:"action_filter"`
	IpAddressFilter   interface{} `db:"ip_address_filter" json:"ip_address_filter"`
	RecordKeyFilter   interface{} `db:"record_key_filter" json:"record_key_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
	SortField         interface{} `db:"sort_field" json:"sort_field"`
	SortOrder         interface{} `db:"sort_order" json:"sort_order"`
}

func (q *Queries) GetPaginatedAuditLogs(ctx context.Context, arg GetPaginatedAuditLogsParams) ([]AuditLog, error) {
	rows, err := q.db.Query(ctx, GetPaginatedAuditLogs,
		arg.Limit,
		arg.Offset,
		arg.TableNameFilter,
		arg.ActorFilter,
		arg.ActionFilter,
		arg.IpAddressFilter,
		arg.RecordKeyFilter,
		arg.DescriptionFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AuditLog{}
	for rows.Next() {
		var i AuditLog
		if err := rows.Scan(
			&i.ID,
			&i.TableName,
			&i.Actor,
			&i.Action,
			&i.IpAddress,
			&i.RecordKey,
			&i.Description,
			&i.Timestamp,
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

const GetPaginatedAuditLogsByUserEmail = `-- name: GetPaginatedAuditLogsByUserEmail :many
SELECT al.id, table_name, actor, action, ip_address, record_key, description, timestamp, u.id, first_name, last_name, email, password, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, last_notification_read_at, created_at, updated_at, deleted_at, created_by, updated_by FROM audit_logs al
JOIN users u ON al.actor = u.email
WHERE 
    u.email = $3  -- Filter by user_id
    AND (coalesce($4, '') = '' OR al.table_name ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR al.actor ILIKE '%' || $5 || '%')
    AND (coalesce($6, '') = '' OR al.action ILIKE '%' || $6 || '%')
    AND (coalesce($7, '') = '' OR al.ip_address ILIKE '%' || $7 || '%')
    AND (coalesce($8, '') = '' OR al.record_key ILIKE '%' || $8 || '%')
    AND (coalesce($9, '') = '' OR al.description ILIKE '%' || $9 || '%')
ORDER BY 
    CASE 
        WHEN $10 = 'TABLENAME' AND $11 = 'ASC' THEN al.table_name 
        WHEN $10 = 'ACTOR' AND $11 = 'ASC' THEN al.actor 
        WHEN $10 = 'ACTION' AND $11 = 'ASC' THEN al.action 
        WHEN $10 = 'IPADDRESS' AND $11 = 'ASC' THEN al.ip_address 
        WHEN $10 = 'RECORDKEY' AND $11 = 'ASC' THEN al.record_key 
        WHEN $10 = 'DESCRIPTION' AND $11 = 'ASC' THEN al.description
    END ASC,
    CASE 
        WHEN $10 = 'TIMESTAMP' AND $11 = 'ASC' THEN al.timestamp
    END ASC,
    CASE 
        WHEN $10 = 'TABLENAME' AND $11 = 'DESC' THEN al.table_name 
        WHEN $10 = 'ACTOR' AND $11 = 'DESC' THEN al.actor 
        WHEN $10 = 'ACTION' AND $11 = 'DESC' THEN al.action 
        WHEN $10 = 'IPADDRESS' AND $11 = 'DESC' THEN al.ip_address 
        WHEN $10 = 'RECORDKEY' AND $11 = 'DESC' THEN al.record_key 
        WHEN $10 = 'DESCRIPTION' AND $11 = 'DESC' THEN al.description
    END DESC,
    CASE 
        WHEN $10 = 'TIMESTAMP' AND $11 = 'DESC' THEN al.timestamp 
    END DESC
LIMIT $1
OFFSET $2
`

type GetPaginatedAuditLogsByUserEmailParams struct {
	Limit             int32       `db:"limit" json:"limit"`
	Offset            int32       `db:"offset" json:"offset"`
	UserEmail         pgtype.Text `db:"user_email" json:"user_email"`
	TableNameFilter   interface{} `db:"table_name_filter" json:"table_name_filter"`
	ActorFilter       interface{} `db:"actor_filter" json:"actor_filter"`
	ActionFilter      interface{} `db:"action_filter" json:"action_filter"`
	IpAddressFilter   interface{} `db:"ip_address_filter" json:"ip_address_filter"`
	RecordKeyFilter   interface{} `db:"record_key_filter" json:"record_key_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
	SortField         interface{} `db:"sort_field" json:"sort_field"`
	SortOrder         interface{} `db:"sort_order" json:"sort_order"`
}

type GetPaginatedAuditLogsByUserEmailRow struct {
	ID                     int64       `db:"id" json:"id"`
	TableName              string      `db:"table_name" json:"table_name"`
	Actor                  string      `db:"actor" json:"actor"`
	Action                 string      `db:"action" json:"action"`
	IpAddress              string      `db:"ip_address" json:"ip_address"`
	RecordKey              string      `db:"record_key" json:"record_key"`
	Description            string      `db:"description" json:"description"`
	Timestamp              int64       `db:"timestamp" json:"timestamp"`
	ID_2                   int64       `db:"id_2" json:"id_2"`
	FirstName              string      `db:"first_name" json:"first_name"`
	LastName               string      `db:"last_name" json:"last_name"`
	Email                  string      `db:"email" json:"email"`
	Password               string      `db:"password" json:"password"`
	JobTitle               pgtype.Text `db:"job_title" json:"job_title"`
	LineOfBusiness         pgtype.Text `db:"line_of_business" json:"line_of_business"`
	LineManager            pgtype.Text `db:"line_manager" json:"line_manager"`
	EmailVerifiedAt        pgtype.Int8 `db:"email_verified_at" json:"email_verified_at"`
	LastSeenAt             int64       `db:"last_seen_at" json:"last_seen_at"`
	LastNotificationReadAt int64       `db:"last_notification_read_at" json:"last_notification_read_at"`
	CreatedAt              int64       `db:"created_at" json:"created_at"`
	UpdatedAt              int64       `db:"updated_at" json:"updated_at"`
	DeletedAt              pgtype.Int8 `db:"deleted_at" json:"deleted_at"`
	CreatedBy              string      `db:"created_by" json:"created_by"`
	UpdatedBy              string      `db:"updated_by" json:"updated_by"`
}

func (q *Queries) GetPaginatedAuditLogsByUserEmail(ctx context.Context, arg GetPaginatedAuditLogsByUserEmailParams) ([]GetPaginatedAuditLogsByUserEmailRow, error) {
	rows, err := q.db.Query(ctx, GetPaginatedAuditLogsByUserEmail,
		arg.Limit,
		arg.Offset,
		arg.UserEmail,
		arg.TableNameFilter,
		arg.ActorFilter,
		arg.ActionFilter,
		arg.IpAddressFilter,
		arg.RecordKeyFilter,
		arg.DescriptionFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPaginatedAuditLogsByUserEmailRow{}
	for rows.Next() {
		var i GetPaginatedAuditLogsByUserEmailRow
		if err := rows.Scan(
			&i.ID,
			&i.TableName,
			&i.Actor,
			&i.Action,
			&i.IpAddress,
			&i.RecordKey,
			&i.Description,
			&i.Timestamp,
			&i.ID_2,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.JobTitle,
			&i.LineOfBusiness,
			&i.LineManager,
			&i.EmailVerifiedAt,
			&i.LastSeenAt,
			&i.LastNotificationReadAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const GetPaginatedAuditLogsByUserEmailCount = `-- name: GetPaginatedAuditLogsByUserEmailCount :one
SELECT COUNT(*) FROM audit_logs al
JOIN users u ON al.actor = u.email
WHERE 
    u.email = $1  -- Filter by user_id
    AND (coalesce($2, '') = '' OR al.table_name ILIKE '%' || $2 || '%')
    AND (coalesce($3, '') = '' OR al.actor ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR al.action ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR al.ip_address ILIKE '%' || $5 || '%')
    AND (coalesce($6, '') = '' OR al.record_key ILIKE '%' || $6 || '%')
    AND (coalesce($7, '') = '' OR al.description ILIKE '%' || $7 || '%')
`

type GetPaginatedAuditLogsByUserEmailCountParams struct {
	UserEmail         pgtype.Text `db:"user_email" json:"user_email"`
	TableNameFilter   interface{} `db:"table_name_filter" json:"table_name_filter"`
	ActorFilter       interface{} `db:"actor_filter" json:"actor_filter"`
	ActionFilter      interface{} `db:"action_filter" json:"action_filter"`
	IpAddressFilter   interface{} `db:"ip_address_filter" json:"ip_address_filter"`
	RecordKeyFilter   interface{} `db:"record_key_filter" json:"record_key_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
}

func (q *Queries) GetPaginatedAuditLogsByUserEmailCount(ctx context.Context, arg GetPaginatedAuditLogsByUserEmailCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, GetPaginatedAuditLogsByUserEmailCount,
		arg.UserEmail,
		arg.TableNameFilter,
		arg.ActorFilter,
		arg.ActionFilter,
		arg.IpAddressFilter,
		arg.RecordKeyFilter,
		arg.DescriptionFilter,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const GetPaginatedAuditLogsCount = `-- name: GetPaginatedAuditLogsCount :one
SELECT COUNT(*) FROM audit_logs al
WHERE 
    (coalesce($1, '') = '' OR al.table_name ILIKE '%' || $1 || '%')
    AND (coalesce($2, '') = '' OR al.actor ILIKE '%' || $2 || '%')
    AND (coalesce($3, '') = '' OR al.action ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR al.ip_address ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR al.record_key ILIKE '%' || $5 || '%')
    AND (coalesce($6, '') = '' OR al.description ILIKE '%' || $6 || '%')
`

type GetPaginatedAuditLogsCountParams struct {
	TableNameFilter   interface{} `db:"table_name_filter" json:"table_name_filter"`
	ActorFilter       interface{} `db:"actor_filter" json:"actor_filter"`
	ActionFilter      interface{} `db:"action_filter" json:"action_filter"`
	IpAddressFilter   interface{} `db:"ip_address_filter" json:"ip_address_filter"`
	RecordKeyFilter   interface{} `db:"record_key_filter" json:"record_key_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
}

func (q *Queries) GetPaginatedAuditLogsCount(ctx context.Context, arg GetPaginatedAuditLogsCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, GetPaginatedAuditLogsCount,
		arg.TableNameFilter,
		arg.ActorFilter,
		arg.ActionFilter,
		arg.IpAddressFilter,
		arg.RecordKeyFilter,
		arg.DescriptionFilter,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}
