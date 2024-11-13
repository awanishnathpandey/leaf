// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: folders.sql

package generated

import (
	"context"
)

const createFolder = `-- name: CreateFolder :one
INSERT INTO folders (
  name,
  slug,
  description,
  created_by,
  updated_by
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, slug, description, created_at, updated_at, created_by, updated_by
`

type CreateFolderParams struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

func (q *Queries) CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error) {
	row := q.db.QueryRow(ctx, createFolder,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.CreatedBy,
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

const deleteFolder = `-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = $1
`

func (q *Queries) DeleteFolder(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFolder, id)
	return err
}

const getFolder = `-- name: GetFolder :one
SELECT id, name, slug, description, created_at, updated_at, created_by, updated_by FROM folders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFolder(ctx context.Context, id int32) (Folder, error) {
	row := q.db.QueryRow(ctx, getFolder, id)
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

const listFolders = `-- name: ListFolders :many
SELECT id, name, slug, description, created_at, updated_at, created_by, updated_by FROM folders
ORDER BY name
`

func (q *Queries) ListFolders(ctx context.Context) ([]Folder, error) {
	rows, err := q.db.Query(ctx, listFolders)
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

const updateFolder = `-- name: UpdateFolder :exec
UPDATE folders
  set name = $2, slug = $3, description = $4, created_by = $5, updated_by = $6
WHERE id = $1
`

type UpdateFolderParams struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}

func (q *Queries) UpdateFolder(ctx context.Context, arg UpdateFolderParams) error {
	_, err := q.db.Exec(ctx, updateFolder,
		arg.ID,
		arg.Name,
		arg.Slug,
		arg.Description,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	return err
}
