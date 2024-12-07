-- name: ListCronJobs :many
SELECT * FROM cron_jobs;


-- name: CreateCronJobLog :one
INSERT INTO cron_job_logs (
    cron_slug,
    status,
    start_time
) VALUES (
    $1, -- cron_slug
    'started', -- status
    EXTRACT(EPOCH FROM NOW()) -- start_time
) RETURNING id;

-- name: UpdateCronJobLogSuccess :exec
UPDATE cron_job_logs
SET 
    status = 'success',
    end_time = EXTRACT(EPOCH FROM NOW()),
    affected_records = $2 -- affected_records
WHERE id = $1;

-- name: UpdateCronJobLogFailed :exec
UPDATE cron_job_logs
SET 
    status = 'failed',
    message = $2, -- error message
    end_time = EXTRACT(EPOCH FROM NOW())
WHERE id = $1;