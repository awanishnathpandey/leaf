-- +goose Up
CREATE TABLE notification_templates (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    description TEXT NOT NULL,
    response_options TEXT[] NULL, -- Stored as an array of strings
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL,
    updated_by VARCHAR(255) NOT NULL
);

CREATE TABLE notifications (
    id BIGSERIAL PRIMARY KEY,
    notification_type VARCHAR(50) NOT NULL, -- 'FILE' or 'TEMPLATE'
    record_key_id BIGINT NOT NULL, -- References `id` in either `file_notifications` or `template_notifications`
    payload JSONB NOT NULL,
    start_time_at BIGINT NOT NULL DEFAULT 0,
    end_time_at BIGINT NOT NULL DEFAULT 0,
    is_push_notification BOOLEAN NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'SCHEDULED',
    group_ids BIGINT[] NULL, -- Array of group IDs
    user_ids BIGINT[] NULL, -- Array of user IDs
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL
);

CREATE TABLE user_notification_responses (
    id BIGSERIAL PRIMARY KEY,
    notification_id BIGINT NOT NULL REFERENCES notifications(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,
    response TEXT NOT NULL, -- User's response
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL
);

-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
DROP TABLE user_notification_responses;
DROP TABLE notifications;
DROP TABLE notification_templates;
-- +goose StatementBegin

-- +goose StatementEnd
