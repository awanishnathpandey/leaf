-- +goose Up
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
-- +goose StatementBegin
INSERT INTO cron_jobs (slug, name, schedule, active, description) VALUES
('sync_users', 'Sync Users', '0 0 * * *', false, 'Runs every day at midnight'),
('clean_audit_logs', 'Clean Audit Logs', '0 0 * * *', false, 'Runs every day at midnight'),
('push_notifications', 'Push Notifications', '*/5 * * * *', false, 'Sends notifications every 5 minutes'); -- can support '@every 10s'
-- +goose StatementEnd

-- +goose Down
DROP TABLE cron_jobs;
-- +goose StatementBegin

-- +goose StatementEnd
