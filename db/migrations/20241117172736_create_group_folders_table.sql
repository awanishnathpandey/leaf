-- +goose Up
CREATE TABLE
    group_folders (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, folder_id)
    );

-- +goose Down
DROP TABLE group_folders;