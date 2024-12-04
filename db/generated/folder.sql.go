// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: folder.sql

package generated

import (
	"context"
)

const CreateFolder = `-- name: CreateFolder :one
INSERT INTO folders (
  name,
  slug,
  description,
  created_by,
  updated_by
) VALUES (
  $1, $2, $3, $4, $4
)
RETURNING id, name, slug, description, created_at, updated_at, created_by, updated_by
`

type CreateFolderParams struct {
	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	Description string `db:"description" json:"description"`
	CreatedBy   string `db:"created_by" json:"created_by"`
}

func (q *Queries) CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, CreateFolder,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.CreatedBy,
	)
	var i Folder
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const DeleteFolder = `-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1
`

func (q *Queries) DeleteFolder(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, DeleteFolder, id)
	return err
}

const DeleteFoldersByIDs = `-- name: DeleteFoldersByIDs :exec
DELETE FROM folders
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteFoldersByIDs(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, DeleteFoldersByIDs, dollar_1)
	return err
}

const GetFolder = `-- name: GetFolder :one
SELECT id, name, slug, description, created_at, updated_at, created_by, updated_by FROM folders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFolder(ctx context.Context, id int64) (Folder, error) {
	row := q.db.QueryRow(ctx, GetFolder, id)
	var i Folder
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const GetFoldersByIDs = `-- name: GetFoldersByIDs :many
SELECT id FROM folders
WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetFoldersByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, GetFoldersByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
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

const ListFolders = `-- name: ListFolders :many
SELECT id, name, slug, description, created_at, updated_at, created_by, updated_by FROM folders
ORDER BY name
`

func (q *Queries) ListFolders(ctx context.Context) ([]Folder, error) {
	rows, err := q.db.Query(ctx, ListFolders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Folder
	for rows.Next() {
		var i Folder
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.Description,
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

const PaginatedFolders = `-- name: PaginatedFolders :many
SELECT id, name, slug, description, created_at, updated_at, created_by, updated_by FROM folders
WHERE 
    (coalesce($3, '') = '' OR name ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR slug ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR description ILIKE '%' || $5 || '%')
ORDER BY 
    CASE 
        WHEN $6 = 'NAME' AND $7 = 'ASC' THEN name 
        WHEN $6 = 'SLUG' AND $7 = 'ASC' THEN slug 
        WHEN $6 = 'DESCRPITION' AND $7 = 'ASC' THEN description 
    END ASC,
    CASE 
        WHEN $6 = 'NAME' AND $7 = 'DESC' THEN name 
        WHEN $6 = 'SLUG' AND $7 = 'DESC' THEN slug 
        WHEN $6 = 'DESCRIPTION' AND $7 = 'DESC' THEN description 
    END DESC
LIMIT $1
OFFSET $2
`

type PaginatedFoldersParams struct {
	Limit             int32       `db:"limit" json:"limit"`
	Offset            int32       `db:"offset" json:"offset"`
	NameFilter        interface{} `db:"name_filter" json:"name_filter"`
	SlugFilter        interface{} `db:"slug_filter" json:"slug_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
	SortField         interface{} `db:"sort_field" json:"sort_field"`
	SortOrder         interface{} `db:"sort_order" json:"sort_order"`
}

func (q *Queries) PaginatedFolders(ctx context.Context, arg PaginatedFoldersParams) ([]Folder, error) {
	rows, err := q.db.Query(ctx, PaginatedFolders,
		arg.Limit,
		arg.Offset,
		arg.NameFilter,
		arg.SlugFilter,
		arg.DescriptionFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Folder
	for rows.Next() {
		var i Folder
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.Description,
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

const PaginatedFoldersCount = `-- name: PaginatedFoldersCount :one
SELECT COUNT(*) FROM folders
WHERE 
    (coalesce($1, '') = '' OR name ILIKE '%' || $1 || '%')
    AND (coalesce($2, '') = '' OR slug ILIKE '%' || $2 || '%')
    AND (coalesce($3, '') = '' OR description ILIKE '%' || $3 || '%')
`

type PaginatedFoldersCountParams struct {
	NameFilter        interface{} `db:"name_filter" json:"name_filter"`
	SlugFilter        interface{} `db:"slug_filter" json:"slug_filter"`
	DescriptionFilter interface{} `db:"description_filter" json:"description_filter"`
}

func (q *Queries) PaginatedFoldersCount(ctx context.Context, arg PaginatedFoldersCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, PaginatedFoldersCount, arg.NameFilter, arg.SlugFilter, arg.DescriptionFilter)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const UpdateFolder = `-- name: UpdateFolder :one
UPDATE folders
SET name = $2, slug = $3, description = $4, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $5
WHERE id = $1
RETURNING id, name, slug, description, created_at, updated_at, created_by, updated_by
`

type UpdateFolderParams struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Slug        string `db:"slug" json:"slug"`
	Description string `db:"description" json:"description"`
	UpdatedBy   string `db:"updated_by" json:"updated_by"`
}

func (q *Queries) UpdateFolder(ctx context.Context, arg UpdateFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, UpdateFolder,
		arg.ID,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.UpdatedBy,
	)
	var i Folder
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
