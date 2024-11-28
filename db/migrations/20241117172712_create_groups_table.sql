-- +goose Up
CREATE TABLE
    groups (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );

-- +goose StatementBegin
INSERT INTO groups (name, description) VALUES ('Default Group', 'Default group for end users');
INSERT INTO groups (name, description) VALUES ('Admin Group', 'Group for admin for all content access');
-- +goose StatementEnd

-- +goose Down
DROP TABLE groups;