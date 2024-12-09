-- +goose Up
CREATE TABLE app_config (
    id BIGSERIAL PRIMARY KEY,
    config_key TEXT UNIQUE NOT NULL,
    config_data JSONB NOT NULL,
    created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    created_by VARCHAR(255) NOT NULL DEFAULT 'system',
    updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
);
-- +goose StatementBegin
INSERT INTO app_config (config_key, config_data) VALUES
('home_page', '{"primary": "Welcome to Our App", "secondary": "Your journey starts here"}'),
('goc_phone_numbers', '[{"name": "Asia", "phone_number": "123-456-7890"}, {"name": "America", "phone_number": "987-654-3210"}, {"name": "Europe", "phone_number": "307-654-3210"}]'),
('app_support', '{"email": "support@app.com", "subject": "Support Request"}'),
('service_desk_phone_numbers', '[{"name": "Service Desk 1", "phone_number": "111-222-3333"}, {"name": "Service Desk 2", "phone_number": "444-555-6666"}]'),
('app_detail', '{"version": "1.0.0", "build_number": "100"}');
-- +goose StatementEnd

-- +goose Down
DROP TABLE app_config;
-- +goose StatementBegin

-- +goose StatementEnd
