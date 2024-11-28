-- +goose Up
CREATE TABLE user_roles (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    role_id BIGINT REFERENCES roles(id) ON DELETE CASCADE,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL DEFAULT 'system',
    updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
    PRIMARY KEY (user_id, role_id)
);

-- +goose StatementBegin
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u, roles r
WHERE u.name = 'admin' AND LOWER(r.name) = 'admin';

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u, roles r
WHERE u.name = 'admin2' AND LOWER(r.name) = 'admin';

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u, roles r
WHERE u.name = 'user' AND LOWER(r.name) = 'user';
-- +goose StatementEnd

-- +goose Down
DROP TABLE user_roles;

-- +goose StatementBegin
-- +goose StatementEnd
