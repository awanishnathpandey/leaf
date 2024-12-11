-- db/schema/users.sql
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        job_title VARCHAR(255),
        line_of_business VARCHAR(255),
        line_manager VARCHAR(255),
        email_verified_at BIGINT,
        last_seen_at BIGINT NOT NULL DEFAULT 0,
        last_notification_read_at BIGINT NOT NULL DEFAULT 0,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        deleted_at BIGINT,
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );