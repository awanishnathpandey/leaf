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
INSERT INTO permissions (name, description) VALUES ('create_role', 'Basic create permission for roles.');
INSERT INTO permissions (name, description) VALUES ('read_role', 'Basic read permission for roles.');
INSERT INTO permissions (name, description) VALUES ('update_role', 'Basic update permission for roles.');
INSERT INTO permissions (name, description) VALUES ('delete_role', 'Basic delete permission for roles.');
INSERT INTO permissions (name, description) VALUES ('read_permission', 'Basic read permission for permissions.');
INSERT INTO permissions (name, description) VALUES ('create_user', 'Basic create permission for users.');
INSERT INTO permissions (name, description) VALUES ('read_user', 'Basic read permission for users.');
INSERT INTO permissions (name, description) VALUES ('update_user', 'Basic update permission for users.');
INSERT INTO permissions (name, description) VALUES ('delete_user', 'Basic delete permission for users.');
INSERT INTO permissions (name, description) VALUES ('create_folder', 'Basic create permission for folders.');
INSERT INTO permissions (name, description) VALUES ('read_folder', 'Basic read permission for folders.');
INSERT INTO permissions (name, description) VALUES ('update_folder', 'Basic update permission for folders.');
INSERT INTO permissions (name, description) VALUES ('delete_folder', 'Basic delete permission for folders.');
INSERT INTO permissions (name, description) VALUES ('create_file', 'Basic create permission for files.');
INSERT INTO permissions (name, description) VALUES ('read_file', 'Basic read permission for files.');
INSERT INTO permissions (name, description) VALUES ('update_file', 'Basic update permission for files.');
INSERT INTO permissions (name, description) VALUES ('delete_file', 'Basic delete permission for files.');
INSERT INTO permissions (name, description) VALUES ('create_group', 'Basic create permission for groups.');
INSERT INTO permissions (name, description) VALUES ('read_group', 'Basic read permission for groups.');
INSERT INTO permissions (name, description) VALUES ('update_group', 'Basic update permission for groups.');
INSERT INTO permissions (name, description) VALUES ('delete_group', 'Basic delete permission for groups.');


-- +goose StatementEnd

-- +goose Down
DROP TABLE permissions;

-- +goose StatementBegin
-- +goose StatementEnd
