-- name: ListActiveCronJobs :many
SELECT * FROM cron_jobs
WHERE active = true;