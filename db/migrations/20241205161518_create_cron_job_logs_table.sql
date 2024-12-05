-- +goose Up
CREATE TABLE cron_job_logs (
    id BIGSERIAL PRIMARY KEY,
    cron_slug VARCHAR(255),
    status VARCHAR(50),
    message TEXT,
    start_time BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    end_time BIGINT,
    affected_records BIGINT
);
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
DROP TABLE cron_job_logs;
-- +goose StatementBegin

-- +goose StatementEnd
