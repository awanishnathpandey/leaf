-- +goose Up
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        job_title VARCHAR(255),
        line_of_business VARCHAR(255),
        email_verified_at TIMESTAMPTZ,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        deleted_at TIMESTAMPTZ,
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
        deleted_by VARCHAR(255)
    );

-- +goose Down
DROP TABLE users;