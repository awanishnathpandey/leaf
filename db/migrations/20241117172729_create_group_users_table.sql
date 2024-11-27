-- +goose Up
CREATE TABLE
    group_users (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, user_id)
    );

-- +goose StatementBegin
INSERT INTO group_users (user_id, group_id)
SELECT u.id, g.id
FROM users u, groups g
WHERE LOWER(u.name) = 'admin' AND g.name = 'Admin Group';

INSERT INTO group_users (user_id, group_id)
SELECT u.id, g.id
FROM users u, groups g
WHERE LOWER(u.name) = 'admin' AND g.name = 'Default Group';
-- +goose StatementEnd

-- +goose Down
DROP TABLE group_users;