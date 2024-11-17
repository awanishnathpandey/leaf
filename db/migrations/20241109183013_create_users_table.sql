-- +goose Up
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        email_verified_at BIGINT,
        last_seen_at BIGINT NOT NULL DEFAULT 0,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        deleted_at BIGINT
    );

-- +goose StatementBegin
INSERT INTO users (name, email, password, email_verified_at) VALUES ('admin', 'admin@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
-- DELETE FROM users WHERE email = 'admin@example.com';
-- +goose StatementEnd