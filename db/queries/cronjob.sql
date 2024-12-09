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

-- name: GetCronJob :one
SELECT * FROM cron_jobs
WHERE slug = $1 LIMIT 1;

-- name: UpdateCronJob :one
UPDATE cron_jobs
SET active = $2, name=$3, description = $4, schedule = $5, updated_at = EXTRACT(EPOCH FROM NOW()), updated_by = $6
WHERE slug = $1
RETURNING *;

-- name: GetCronJobLog :one
SELECT * FROM cron_job_logs
WHERE id = $1 LIMIT 1;

-- name: DeleteCronJobLog :exec
DELETE FROM cron_job_logs
WHERE id = $1;

-- name: GetCronJobLogsByIDs :many
SELECT id FROM cron_job_logs
WHERE id = ANY($1::bigint[]);

-- name: DeleteCronJobLogsByIDs :exec
DELETE FROM cron_job_logs
WHERE id = ANY($1::bigint[]);

-- name: GetPaginatedCronJobs :many
SELECT * FROM cron_jobs cj
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR cj.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR cj.description ILIKE '%' || sqlc.narg(description_filter) || '%')
    AND (coalesce(sqlc.narg(schedule_filter), '') = '' OR cj.schedule ILIKE '%' || sqlc.narg(schedule_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'ASC' THEN cj.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'ASC' THEN cj.description 
        WHEN sqlc.narg(sort_field) = 'SCHEDULE' AND sqlc.narg(sort_order) = 'ASC' THEN cj.schedule 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'NAME' AND sqlc.narg(sort_order) = 'DESC' THEN cj.name 
        WHEN sqlc.narg(sort_field) = 'DESCRIPTION' AND sqlc.narg(sort_order) = 'DESC' THEN cj.description 
        WHEN sqlc.narg(sort_field) = 'SCHEDULE' AND sqlc.narg(sort_order) = 'DESC' THEN cj.schedule 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedCronJobsCount :one
SELECT COUNT(*) FROM cron_jobs cj
WHERE 
    (coalesce(sqlc.narg(name_filter), '') = '' OR cj.name ILIKE '%' || sqlc.narg(name_filter) || '%')
    AND (coalesce(sqlc.narg(description_filter), '') = '' OR cj.description ILIKE '%' || sqlc.narg(description_filter) || '%')
    AND (coalesce(sqlc.narg(schedule_filter), '') = '' OR cj.schedule ILIKE '%' || sqlc.narg(schedule_filter) || '%');


-- name: GetPaginatedCronJobLogsByCronSlug :many
SELECT * FROM cron_job_logs cjl
JOIN cron_jobs cj ON cjl.cron_slug = cj.slug
WHERE 
    cj.slug = sqlc.narg(cron_job_slug)  -- Filter by cron job slug
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR cjl.cron_slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(message_filter), '') = '' OR cjl.message ILIKE '%' || sqlc.narg(message_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN cjl.cron_slug 
        WHEN sqlc.narg(sort_field) = 'MESSAGE' AND sqlc.narg(sort_order) = 'ASC' THEN cjl.message 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN cjl.cron_slug 
        WHEN sqlc.narg(sort_field) = 'MESSAGE' AND sqlc.narg(sort_order) = 'DESC' THEN cjl.message 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedCronJobLogsByCronSlugCount :one
SELECT COUNT(*) FROM cron_job_logs cjl
JOIN cron_jobs cj ON cjl.cron_slug = cj.slug
WHERE 
    cj.slug = sqlc.narg(cron_job_slug)  -- Filter by cron job slug
    AND (coalesce(sqlc.narg(slug_filter), '') = '' OR cjl.cron_slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(message_filter), '') = '' OR cjl.message ILIKE '%' || sqlc.narg(message_filter) || '%');

-- name: GetPaginatedCronJobLogs :many
SELECT * FROM cron_job_logs cjl
WHERE 
    (coalesce(sqlc.narg(slug_filter), '') = '' OR cjl.cron_slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(message_filter), '') = '' OR cjl.message ILIKE '%' || sqlc.narg(message_filter) || '%')
ORDER BY 
    CASE 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'ASC' THEN cjl.cron_slug 
        WHEN sqlc.narg(sort_field) = 'MESSAGE' AND sqlc.narg(sort_order) = 'ASC' THEN cjl.message 
    END ASC,
    CASE 
        WHEN sqlc.narg(sort_field) = 'SLUG' AND sqlc.narg(sort_order) = 'DESC' THEN cjl.cron_slug 
        WHEN sqlc.narg(sort_field) = 'MESSAGE' AND sqlc.narg(sort_order) = 'DESC' THEN cjl.message 
    END DESC
LIMIT $1
OFFSET $2;

-- name: GetPaginatedCronJobLogsCount :one
SELECT COUNT(*) FROM cron_job_logs cjl
WHERE 
    (coalesce(sqlc.narg(slug_filter), '') = '' OR cjl.slug ILIKE '%' || sqlc.narg(slug_filter) || '%')
    AND (coalesce(sqlc.narg(message_filter), '') = '' OR cjl.message ILIKE '%' || sqlc.narg(message_filter) || '%');