-- +goose Up
CREATE TABLE
    files (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        url VARCHAR(255) NOT NULL,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE, -- Foreign key
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
    );

-- +goose Down
DROP TABLE files;