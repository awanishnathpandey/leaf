-- +goose Up
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
-- +goose StatementBegin
INSERT INTO email_templates (name, content, mail_to, mail_cc, mail_bcc, mail_data)
VALUES 
(
    'welcome_email', 
    '<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Welcome Email</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                margin: 0;
                padding: 0;
                background-color: #f8f9fa;
            }
            .container {
                width: 100%;
                max-width: 600px;
                margin: 20px auto;
                padding: 20px;
                background-color: #ffffff;
                border-radius: 8px;
                box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            }
            h1 {
                color: #343a40;
                font-size: 24px;
            }
            p {
                color: #495057;
                font-size: 16px;
                line-height: 1.5;
            }
            .footer {
                margin-top: 30px;
                font-size: 12px;
                color: #6c757d;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Hello {{.Name}},</h1>
            <p>Welcome to our platform!</p>
            <p>Your email: {{.Email}}</p>
            <div class="footer">
                <p>Best regards,<br/>The Team</p>
            </div>
        </div>
    </body>
    </html>',
    ARRAY['admin@example.com', 'admin2@example.com'], 
    ARRAY['admincc@example.com', 'admin2cc@example.com'], 
    ARRAY['adminbcc@example.com', 'admin2bcc@example.com'], 
    '{"Name": "John Doe", "Email": "YV5tD@example.com"}'::jsonb
);
INSERT INTO email_templates (name, content, mail_data)
VALUES 
(
    'password_reset', 
    '<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Welcome Email</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                margin: 0;
                padding: 0;
                background-color: #f8f9fa;
            }
            .container {
                width: 100%;
                max-width: 600px;
                margin: 20px auto;
                padding: 20px;
                background-color: #ffffff;
                border-radius: 8px;
                box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            }
            h1 {
                color: #343a40;
                font-size: 24px;
            }
            p {
                color: #495057;
                font-size: 16px;
                line-height: 1.5;
            }
            .footer {
                margin-top: 30px;
                font-size: 12px;
                color: #6c757d;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Hello {{.Name}},</h1>
            <p>Here are your password reset instructions!</p>
            <p>Your email: {{.Email}}</p>
            <p>Your reset token: {{.ResetToken}}</p>
            <div class="footer">
                <p>Best regards,<br/>The Team</p>
            </div>
        </div>
    </body>
    </html>', 
    '{"Name": "John Doe", "Email": "YV5tD@example.com", "ResetToken": "1234567890"}'::jsonb
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE email_templates;
-- +goose StatementBegin
-- +goose StatementEnd
