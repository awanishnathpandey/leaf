-- +goose Up
CREATE TABLE
    group_files (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        file_id BIGINT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
        PRIMARY KEY (group_id, file_id)
    );

-- +goose Down
DROP TABLE group_files;