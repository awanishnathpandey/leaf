-- +goose Up
CREATE TABLE email_templates (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);
-- +goose StatementBegin
INSERT INTO email_templates (name, content) VALUES 
('welcome_email', '<html><body><h1>Hello {{.Name}},</h1><p>Welcome to our platform!</p><p>Your email: {{.Email}}</p></body></html>');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
