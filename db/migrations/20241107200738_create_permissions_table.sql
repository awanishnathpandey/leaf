-- +goose Up
CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);

-- +goose StatementBegin
INSERT INTO permissions (name, description) VALUES ('all', 'Explicit permission for administrators.');
INSERT INTO permissions (name, description) VALUES ('admin_access', 'Basic access permission for admin client.');
INSERT INTO permissions (name, description) VALUES ('read', 'Basic read permission for end users.');
INSERT INTO permissions (name, description) VALUES ('read_folder', 'Basic read permission for folders.');
-- +goose StatementEnd

-- +goose Down
DROP TABLE permissions;

-- +goose StatementBegin
-- +goose StatementEnd
