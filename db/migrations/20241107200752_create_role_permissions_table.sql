-- +goose Up
CREATE TABLE role_permissions (
    role_id BIGINT REFERENCES roles(id) ON DELETE CASCADE,
    permission_id BIGINT REFERENCES permissions(id) ON DELETE CASCADE,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL DEFAULT 'system',
    updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
    PRIMARY KEY (role_id, permission_id)
);

-- +goose StatementBegin
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE LOWER(r.name) = 'admin' AND p.name IN ('all');

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE LOWER(r.name) = 'user' AND p.name IN ('admin_access', 'read');
-- +goose StatementEnd

-- +goose Down
DROP TABLE role_permissions;

-- +goose StatementBegin
-- +goose StatementEnd
