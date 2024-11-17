-- db/schema/groups.sql
CREATE TABLE
    groups (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
    );

CREATE TABLE
    group_users (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, user_id)
    );

CREATE TABLE
    group_folders (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, folder_id)
    );

CREATE TABLE
    group_files (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        file_id BIGINT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        PRIMARY KEY (group_id, file_id)
    );