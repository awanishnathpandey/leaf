-- +goose Up
CREATE TABLE
    group_folders (
        group_id BIGINT NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
        folder_id BIGINT NOT NULL REFERENCES folders(id) ON DELETE CASCADE,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
        PRIMARY KEY (group_id, folder_id)
    );

-- +goose StatementBegin
INSERT INTO group_folders (folder_id, group_id)
SELECT f.id, g.id
FROM folders f, groups g
WHERE f.slug = 'support' AND g.name = 'Default Group';

INSERT INTO group_folders (folder_id, group_id)
SELECT f.id, g.id
FROM folders f, groups g
WHERE f.slug = 'videos' AND g.name = 'Default Group';

INSERT INTO group_folders (folder_id, group_id)
SELECT f.id, g.id
FROM folders f, groups g
WHERE f.slug = 'support' AND g.name = 'Admin Group';

INSERT INTO group_folders (folder_id, group_id)
SELECT f.id, g.id
FROM folders f, groups g
WHERE f.slug = 'videos' AND g.name = 'Admin Group';
-- +goose StatementEnd

-- +goose Down
DROP TABLE group_folders;