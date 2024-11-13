-- +goose Up
CREATE TABLE
    folders (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        description VARCHAR(255) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );

-- +goose Down
DROP TABLE folders;