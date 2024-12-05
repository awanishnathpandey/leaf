CREATE TABLE cron_jobs (
    id BIGSERIAL PRIMARY KEY,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    schedule TEXT NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    description TEXT,
    last_run_at BIGINT,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL DEFAULT 'system',
    updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
);


CREATE TABLE cron_job_logs (
    id BIGSERIAL PRIMARY KEY,
    cron_slug VARCHAR(255),
    status VARCHAR(50),
    message TEXT,
    start_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    end_time BIGINT,
    affected_records BIGINT
);