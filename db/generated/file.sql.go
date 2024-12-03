// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: file.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFile = `-- name: CreateFile :one
INSERT INTO files (name, slug, file_path, folder_id, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $5)
RETURNING id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by
`

type CreateFileParams struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	FilePath  string `json:"file_path"`
	FolderID  int64  `json:"folder_id"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) CreateFile(ctx context.Context, arg CreateFileParams) (File, error) {
	row := q.db.QueryRow(ctx, createFile,
		arg.Name,
		arg.Slug,
		arg.FilePath,
		arg.FolderID,
		arg.CreatedBy,
	)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.FilePath,
		&i.FileType,
		&i.FileBytes,
		&i.AutoDownload,
		&i.FolderID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteFile = `-- name: DeleteFile :exec
DELETE FROM files
WHERE id = $1
`

func (q *Queries) DeleteFile(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFile, id)
	return err
}

const deleteFilesByIDs = `-- name: DeleteFilesByIDs :exec
DELETE FROM files
WHERE id = ANY($1::bigint[])
`

func (q *Queries) DeleteFilesByIDs(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, deleteFilesByIDs, dollar_1)
	return err
}

const getFile = `-- name: GetFile :one
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files
WHERE id = $1
`

func (q *Queries) GetFile(ctx context.Context, id int64) (File, error) {
	row := q.db.QueryRow(ctx, getFile, id)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.FilePath,
		&i.FileType,
		&i.FileBytes,
		&i.AutoDownload,
		&i.FolderID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getFilesByFolder = `-- name: GetFilesByFolder :many
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files
WHERE folder_id = $1
`

func (q *Queries) GetFilesByFolder(ctx context.Context, folderID int64) ([]File, error) {
	rows, err := q.db.Query(ctx, getFilesByFolder, folderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.FilePath,
			&i.FileType,
			&i.FileBytes,
			&i.AutoDownload,
			&i.FolderID,
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

const getFilesByFolderID = `-- name: GetFilesByFolderID :many
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files
WHERE folder_id = $1
`

func (q *Queries) GetFilesByFolderID(ctx context.Context, folderID int64) ([]File, error) {
	rows, err := q.db.Query(ctx, getFilesByFolderID, folderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.FilePath,
			&i.FileType,
			&i.FileBytes,
			&i.AutoDownload,
			&i.FolderID,
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

const getFilesByIDs = `-- name: GetFilesByIDs :many
SELECT id FROM files
WHERE id = ANY($1::bigint[])
`

func (q *Queries) GetFilesByIDs(ctx context.Context, dollar_1 []int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getFilesByIDs, dollar_1)
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

const getPaginatedFilesByFolderID = `-- name: GetPaginatedFilesByFolderID :many
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files WHERE 
    folder_id = $3  -- Filter by folder_id
    AND (coalesce($4, '') = '' OR name ILIKE '%' || $4 || '%')
    AND (coalesce($5, '') = '' OR slug ILIKE '%' || $5 || '%')
ORDER BY 
    CASE 
        WHEN $6 = 'NAME' AND $7 = 'ASC' THEN name 
        WHEN $6 = 'SLUG' AND $7 = 'ASC' THEN slug 
    END ASC,
    CASE 
        WHEN $6 = 'NAME' AND $7 = 'DESC' THEN name 
        WHEN $6 = 'SLUG' AND $7 = 'DESC' THEN slug 
    END DESC
LIMIT $1
OFFSET $2
`

type GetPaginatedFilesByFolderIDParams struct {
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
	FolderID   pgtype.Int8 `json:"folder_id"`
	NameFilter interface{} `json:"name_filter"`
	SlugFilter interface{} `json:"slug_filter"`
	SortField  interface{} `json:"sort_field"`
	SortOrder  interface{} `json:"sort_order"`
}

func (q *Queries) GetPaginatedFilesByFolderID(ctx context.Context, arg GetPaginatedFilesByFolderIDParams) ([]File, error) {
	rows, err := q.db.Query(ctx, getPaginatedFilesByFolderID,
		arg.Limit,
		arg.Offset,
		arg.FolderID,
		arg.NameFilter,
		arg.SlugFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.FilePath,
			&i.FileType,
			&i.FileBytes,
			&i.AutoDownload,
			&i.FolderID,
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

const getPaginatedFilesByFolderIDCount = `-- name: GetPaginatedFilesByFolderIDCount :one
SELECT COUNT(*) FROM files WHERE 
    folder_id = $1  -- Filter by folder_id
    AND (coalesce($2, '') = '' OR name ILIKE '%' || $2 || '%')
    AND (coalesce($3, '') = '' OR slug ILIKE '%' || $3 || '%')
`

type GetPaginatedFilesByFolderIDCountParams struct {
	FolderID   pgtype.Int8 `json:"folder_id"`
	NameFilter interface{} `json:"name_filter"`
	SlugFilter interface{} `json:"slug_filter"`
}

func (q *Queries) GetPaginatedFilesByFolderIDCount(ctx context.Context, arg GetPaginatedFilesByFolderIDCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, getPaginatedFilesByFolderIDCount, arg.FolderID, arg.NameFilter, arg.SlugFilter)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listFiles = `-- name: ListFiles :many
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files
ORDER BY name
`

func (q *Queries) ListFiles(ctx context.Context) ([]File, error) {
	rows, err := q.db.Query(ctx, listFiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.FilePath,
			&i.FileType,
			&i.FileBytes,
			&i.AutoDownload,
			&i.FolderID,
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

const paginatedFiles = `-- name: PaginatedFiles :many
SELECT id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by FROM files
WHERE 
    (coalesce($3, '') = '' OR name ILIKE '%' || $3 || '%')
    AND (coalesce($4, '') = '' OR slug ILIKE '%' || $4 || '%')
ORDER BY 
    CASE 
        WHEN $5 = 'NAME' AND $6 = 'ASC' THEN name 
        WHEN $5 = 'SLUG' AND $6 = 'ASC' THEN slug 
    END ASC,
    CASE 
        WHEN $5 = 'NAME' AND $6 = 'DESC' THEN name 
        WHEN $5 = 'SLUG' AND $6 = 'DESC' THEN slug 
    END DESC
LIMIT $1
OFFSET $2
`

type PaginatedFilesParams struct {
	Limit      int32       `json:"limit"`
	Offset     int32       `json:"offset"`
	NameFilter interface{} `json:"name_filter"`
	SlugFilter interface{} `json:"slug_filter"`
	SortField  interface{} `json:"sort_field"`
	SortOrder  interface{} `json:"sort_order"`
}

func (q *Queries) PaginatedFiles(ctx context.Context, arg PaginatedFilesParams) ([]File, error) {
	rows, err := q.db.Query(ctx, paginatedFiles,
		arg.Limit,
		arg.Offset,
		arg.NameFilter,
		arg.SlugFilter,
		arg.SortField,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Slug,
			&i.FilePath,
			&i.FileType,
			&i.FileBytes,
			&i.AutoDownload,
			&i.FolderID,
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

const paginatedFilesCount = `-- name: PaginatedFilesCount :one
SELECT COUNT(*) FROM files
WHERE 
    (coalesce($1, '') = '' OR name ILIKE '%' || $1 || '%')
    AND (coalesce($2, '') = '' OR slug ILIKE '%' || $2 || '%')
`

type PaginatedFilesCountParams struct {
	NameFilter interface{} `json:"name_filter"`
	SlugFilter interface{} `json:"slug_filter"`
}

func (q *Queries) PaginatedFilesCount(ctx context.Context, arg PaginatedFilesCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, paginatedFilesCount, arg.NameFilter, arg.SlugFilter)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateFile = `-- name: UpdateFile :one
UPDATE files
SET name = $2, slug = $3, file_path = $4, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $5
WHERE id = $1
RETURNING id, name, slug, file_path, file_type, file_bytes, auto_download, folder_id, created_at, updated_at, created_by, updated_by
`

type UpdateFileParams struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	FilePath  string `json:"file_path"`
	UpdatedBy string `json:"updated_by"`
}

func (q *Queries) UpdateFile(ctx context.Context, arg UpdateFileParams) (File, error) {
	row := q.db.QueryRow(ctx, updateFile,
		arg.ID,
		arg.Name,
		arg.Slug,
		arg.FilePath,
		arg.UpdatedBy,
	)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Slug,
		&i.FilePath,
		&i.FileType,
		&i.FileBytes,
		&i.AutoDownload,
		&i.FolderID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
