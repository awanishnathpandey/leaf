-- db/schema/files.sql
CREATE TABLE
    files (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        file_path VARCHAR(255) NOT NULL,
        file_type  VARCHAR(20) NOT NULL CHECK (file_type IN ('document', 'video', 'support')) DEFAULT 'document',
        file_bytes BIGINT NOT NULL DEFAULT 0,
        file_content_type VARCHAR(255) NOT NULL DEFAULT '-',
        auto_download BOOLEAN NOT NULL DEFAULT false,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE, -- Foreign key
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );