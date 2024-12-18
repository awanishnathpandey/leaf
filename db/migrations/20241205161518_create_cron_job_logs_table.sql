-- +goose Up
CREATE TABLE cron_job_logs (
    id BIGSERIAL PRIMARY KEY,
    cron_slug TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    message TEXT NOT NULL DEFAULT '-',
    start_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    end_time BIGINT NOT NULL DEFAULT 0,
    affected_records BIGINT NOT NULL DEFAULT 0
);
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
DROP TABLE cron_job_logs;
-- +goose StatementBegin

-- +goose StatementEnd
