// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cronjob.sql

package generated

import (
	"context"
)

const ListActiveCronJobs = `-- name: ListActiveCronJobs :many
SELECT id, slug, name, schedule, active, description, last_run_at, created_at, updated_at, created_by, updated_by FROM cron_jobs
WHERE active = true
`

func (q *Queries) ListActiveCronJobs(ctx context.Context) ([]CronJob, error) {
	rows, err := q.db.Query(ctx, ListActiveCronJobs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CronJob{}
	for rows.Next() {
		var i CronJob
		if err := rows.Scan(
			&i.ID,
			&i.Slug,
			&i.Name,
			&i.Schedule,
			&i.Active,
			&i.Description,
			&i.LastRunAt,
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