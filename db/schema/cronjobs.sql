CREATE TABLE cron_jobs (
    id BIGSERIAL PRIMARY KEY,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    schedule TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    description TEXT NOT NULL,
    last_run_at BIGINT NOT NULL DEFAULT 0,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL DEFAULT 'system',
    updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
);


CREATE TABLE cron_job_logs (
    id BIGSERIAL PRIMARY KEY,
    cron_slug TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    message TEXT NOT NULL DEFAULT '-',
    start_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    end_time BIGINT NOT NULL DEFAULT 0,
    affected_records BIGINT NOT NULL DEFAULT 0
);