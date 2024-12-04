// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const CreateUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email, password, created_by, updated_by
) VALUES (
  $1, $2, $3, $4, $5, $5
)
RETURNING id, first_name, last_name, email, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by
`

type CreateUserParams struct {
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedBy string `db:"created_by" json:"created_by"`
}

type CreateUserRow struct {
	ID              int64       `db:"id" json:"id"`
	FirstName       string      `db:"first_name" json:"first_name"`
	LastName        string      `db:"last_name" json:"last_name"`
	Email           string      `db:"email" json:"email"`
	JobTitle        pgtype.Text `db:"job_title" json:"job_title"`
	LineOfBusiness  pgtype.Text `db:"line_of_business" json:"line_of_business"`
	LineManager     pgtype.Text `db:"line_manager" json:"line_manager"`
	EmailVerifiedAt pgtype.Int8 `db:"email_verified_at" json:"email_verified_at"`
	LastSeenAt      int64       `db:"last_seen_at" json:"last_seen_at"`
	CreatedAt       int64       `db:"created_at" json:"created_at"`
	UpdatedAt       int64       `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Int8 `db:"deleted_at" json:"deleted_at"`
	CreatedBy       string      `db:"created_by" json:"created_by"`
	UpdatedBy       string      `db:"updated_by" json:"updated_by"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, CreateUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.CreatedBy,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.JobTitle,
		&i.LineOfBusiness,
		&i.LineManager,
		&i.EmailVerifiedAt,
		&i.LastSeenAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const DeleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, DeleteUser, id)
	return err
}

const DeleteUsersByIDs = `-- name: DeleteUsersByIDs :exec
DELETE FROM users
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteUsersByIDs(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, DeleteUsersByIDs, dollar_1)
	return err
}

const GetUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, password, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, GetUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.JobTitle,
		&i.LineOfBusiness,
		&i.LineManager,
		&i.EmailVerifiedAt,
		&i.LastSeenAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const GetUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, GetUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.JobTitle,
		&i.LineOfBusiness,
		&i.LineManager,
		&i.EmailVerifiedAt,
		&i.LastSeenAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const GetUserID = `-- name: GetUserID :one
SELECT id FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserID(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, GetUserID, id)
	err := row.Scan(&id)
	return id, err
}

const GetUsersByIDs = `-- name: GetUsersByIDs :many
SELECT id FROM users
WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetUsersByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, GetUsersByIDs, dollar_1)
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

const ListUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email, password, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by FROM users
ORDER BY email
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, ListUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.JobTitle,
			&i.LineOfBusiness,
			&i.LineManager,
			&i.EmailVerifiedAt,
			&i.LastSeenAt,
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

const PaginatedUsers = `-- name: PaginatedUsers :many
SELECT 
    id, 
    first_name, 
    last_name,
    email, 
    password,
    job_title,
    line_of_business,
    line_manager, 
    email_verified_at, 
    last_seen_at, 
    created_at, 
    updated_at, 
    deleted_at,
    created_by,
    updated_by
FROM users
WHERE 
    (coalesce($3, '') = '' OR first_name ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR email ILIKE '%' || $4 || '%')
ORDER BY 
    CASE 
        WHEN $5 = 'NAME' AND $6 = 'ASC' THEN first_name 
        WHEN $5 = 'EMAIL' AND $6 = 'ASC' THEN email 
    END ASC,
    CASE 
        WHEN $5 = 'NAME' AND $6 = 'DESC' THEN first_name 
        WHEN $5 = 'EMAIL' AND $6 = 'DESC' THEN email 
    END DESC
LIMIT $1
OFFSET $2
`

type PaginatedUsersParams struct {
	Limit       int32       `db:"limit" json:"limit"`
	Offset      int32       `db:"offset" json:"offset"`
	NameFilter  interface{} `db:"name_filter" json:"name_filter"`
	EmailFilter interface{} `db:"email_filter" json:"email_filter"`
	SortField   interface{} `db:"sort_field" json:"sort_field"`
	SortOrder   interface{} `db:"sort_order" json:"sort_order"`
}

func (q *Queries) PaginatedUsers(ctx context.Context, arg PaginatedUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, PaginatedUsers,
		arg.Limit,
		arg.Offset,
		arg.NameFilter,
		arg.EmailFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.JobTitle,
			&i.LineOfBusiness,
			&i.LineManager,
			&i.EmailVerifiedAt,
			&i.LastSeenAt,
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

const PaginatedUsersCount = `-- name: PaginatedUsersCount :one
SELECT COUNT(*)
FROM users
WHERE 
    (coalesce($1, '') = '' OR first_name ILIKE '%' || $1 || '%')
    AND (coalesce($2, '') = '' OR email ILIKE '%' || $2 || '%')
`

type PaginatedUsersCountParams struct {
	NameFilter  interface{} `db:"name_filter" json:"name_filter"`
	EmailFilter interface{} `db:"email_filter" json:"email_filter"`
}

func (q *Queries) PaginatedUsersCount(ctx context.Context, arg PaginatedUsersCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, PaginatedUsersCount, arg.NameFilter, arg.EmailFilter)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const UpdateUser = `-- name: UpdateUser :one
UPDATE users
  set first_name = $2, last_name = $3, email = $4, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $5
WHERE id = $1
RETURNING id, first_name, last_name, email, job_title, line_of_business, line_manager, email_verified_at, last_seen_at, created_at, updated_at, deleted_at, created_by, updated_by
`

type UpdateUserParams struct {
	ID        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type UpdateUserRow struct {
	ID              int64       `db:"id" json:"id"`
	FirstName       string      `db:"first_name" json:"first_name"`
	LastName        string      `db:"last_name" json:"last_name"`
	Email           string      `db:"email" json:"email"`
	JobTitle        pgtype.Text `db:"job_title" json:"job_title"`
	LineOfBusiness  pgtype.Text `db:"line_of_business" json:"line_of_business"`
	LineManager     pgtype.Text `db:"line_manager" json:"line_manager"`
	EmailVerifiedAt pgtype.Int8 `db:"email_verified_at" json:"email_verified_at"`
	LastSeenAt      int64       `db:"last_seen_at" json:"last_seen_at"`
	CreatedAt       int64       `db:"created_at" json:"created_at"`
	UpdatedAt       int64       `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Int8 `db:"deleted_at" json:"deleted_at"`
	CreatedBy       string      `db:"created_by" json:"created_by"`
	UpdatedBy       string      `db:"updated_by" json:"updated_by"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, UpdateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.UpdatedBy,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.JobTitle,
		&i.LineOfBusiness,
		&i.LineManager,
		&i.EmailVerifiedAt,
		&i.LastSeenAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const UpdateUserEmailVerifiedAt = `-- name: UpdateUserEmailVerifiedAt :exec
UPDATE users
  set email_verified_at = EXTRACT(EPOCH FROM NOW()), updated_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1
`

func (q *Queries) UpdateUserEmailVerifiedAt(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, UpdateUserEmailVerifiedAt, id)
	return err
}

const UpdateUserLastSeenAt = `-- name: UpdateUserLastSeenAt :exec
UPDATE users
  set last_seen_at = EXTRACT(EPOCH FROM NOW())
WHERE id = $1
`

func (q *Queries) UpdateUserLastSeenAt(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, UpdateUserLastSeenAt, id)
	return err
}

const UpdateUserLastSeenAtByEmail = `-- name: UpdateUserLastSeenAtByEmail :exec
UPDATE users
  set last_seen_at = EXTRACT(EPOCH FROM NOW())
WHERE email = $1
`

func (q *Queries) UpdateUserLastSeenAtByEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, UpdateUserLastSeenAtByEmail, email)
	return err
}
