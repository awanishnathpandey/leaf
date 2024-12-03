-- +goose Up
CREATE TABLE
    password_resets (
        id BIGSERIAL PRIMARY KEY,
        user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        reset_token VARCHAR(255) NOT NULL,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );
-- +goose StatementBegin

-- +goose StatementEnd

-- +goose Down
DROP TABLE password_resets;
-- +goose StatementBegin

-- +goose StatementEnd
