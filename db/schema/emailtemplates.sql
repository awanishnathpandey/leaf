CREATE TABLE email_templates (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    mail_to TEXT[] DEFAULT NULL,
    mail_cc TEXT[] DEFAULT NULL,
    mail_bcc TEXT[] DEFAULT NULL,
    mail_data JSONB DEFAULT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);