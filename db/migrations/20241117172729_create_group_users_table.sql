-- +goose Up
CREATE TABLE
    group_users (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, user_id)
    );

-- +goose Down
DROP TABLE group_users;