-- +goose Up
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL, -- Prevent deletion of default roles
    description VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);

-- +goose StatementBegin
INSERT INTO roles (name, description) VALUES ('Admin', 'Role for administrators.');
INSERT INTO roles (name, description) VALUES ('User', 'Role for end users.');
-- +goose StatementEnd

-- +goose Down
DROP TABLE roles;

-- +goose StatementBegin
-- +goose StatementEnd