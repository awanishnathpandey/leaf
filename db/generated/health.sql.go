// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: health.sql

package generated

import (
	"context"
)

const checkHealth = `-- name: CheckHealth :exec
SELECT 1
`

func (q *Queries) CheckHealth(ctx context.Context) error {
	_, err := q.db.Exec(ctx, checkHealth)
	return err
}
