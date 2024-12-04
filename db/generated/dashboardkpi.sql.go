// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: dashboardkpi.sql

package generated

import (
	"context"
)

const GetDashboardKPICount = `-- name: GetDashboardKPICount :one
SELECT
  (SELECT COUNT(*) FROM users) AS users,
  (SELECT COUNT(*) FROM roles) AS roles,
  (SELECT COUNT(*) FROM permissions) AS permissions,
  (SELECT COUNT(*) FROM groups) AS groups,
  (SELECT COUNT(*) FROM folders) AS folders,
  (SELECT COUNT(*) FROM files) AS files
`

type GetDashboardKPICountRow struct {
	Users       int64 `db:"users" json:"users"`
	Roles       int64 `db:"roles" json:"roles"`
	Permissions int64 `db:"permissions" json:"permissions"`
	Groups      int64 `db:"groups" json:"groups"`
	Folders     int64 `db:"folders" json:"folders"`
	Files       int64 `db:"files" json:"files"`
}

func (q *Queries) GetDashboardKPICount(ctx context.Context) (GetDashboardKPICountRow, error) {
	row := q.db.QueryRow(ctx, GetDashboardKPICount)
	var i GetDashboardKPICountRow
	err := row.Scan(
		&i.Users,
		&i.Roles,
		&i.Permissions,
		&i.Groups,
		&i.Folders,
		&i.Files,
	)
	return i, err
}
