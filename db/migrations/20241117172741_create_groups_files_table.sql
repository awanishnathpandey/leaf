-- +goose Up
CREATE TABLE
    group_files (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        file_id BIGINT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, file_id)
    );

-- +goose Down
DROP TABLE group_files;