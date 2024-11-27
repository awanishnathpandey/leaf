-- +goose Up
CREATE TABLE
    folders (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        slug VARCHAR(255) NOT NULL UNIQUE,
        description VARCHAR(255) NOT NULL,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
    );

-- +goose StatementBegin
INSERT INTO folders (name, slug, description) VALUES ('Support', 'support', 'support documents folder');
INSERT INTO folders (name, slug, description) VALUES ('Videos', 'videos', 'videos folder');
-- +goose StatementEnd

-- +goose Down
DROP TABLE folders;

-- +goose StatementBegin
-- DELETE FROM folders WHERE name = 'sample';
-- +goose StatementEnd