-- db/schema/files.sql
CREATE TABLE
    files (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        url VARCHAR(255) NOT NULL,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE, -- Foreign key
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );