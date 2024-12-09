CREATE TABLE audit_logs (
    id BIGSERIAL PRIMARY KEY,
    table_name VARCHAR(255) NOT NULL,
    actor VARCHAR(255) NOT NULL DEFAULT 'anonymous',
    action VARCHAR(255) NOT NULL,
    ip_address VARCHAR(15) NOT NULL,
    record_key VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    timestamp BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())
);