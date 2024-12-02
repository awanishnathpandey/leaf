-- +goose Up
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        first_name VARCHAR(255) NOT NULL,
        last_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        email_verified_at BIGINT,
        last_seen_at BIGINT NOT NULL DEFAULT 0,
        created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
        deleted_at BIGINT,
        created_by VARCHAR(255) NOT NULL DEFAULT 'system',
        updated_by VARCHAR(255) NOT NULL DEFAULT 'system'
    );

-- +goose StatementBegin
INSERT INTO users (first_name, last_name, email, password, email_verified_at) VALUES ('admin', '-', 'admin@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
INSERT INTO users (first_name, last_name, email, password, email_verified_at) VALUES ('admin2', '-', 'admin2@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
INSERT INTO users (first_name, last_name, email, password, email_verified_at) VALUES ('user', '-', 'user@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user1', 'user1@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user2', 'user2@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user3', 'user3@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user4', 'user4@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user5', 'user5@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user6', 'user6@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user7', 'user7@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user8', 'user8@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user9', 'user9@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user10', 'user10@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user11', 'user11@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user12', 'user12@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user13', 'user13@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user14', 'user14@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user15', 'user15@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user16', 'user16@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user17', 'user17@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user18', 'user18@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user19', 'user19@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user20', 'user20@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user21', 'user21@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user22', 'user22@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user23', 'user23@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user24', 'user24@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user25', 'user25@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user26', 'user26@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user27', 'user27@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user28', 'user28@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user29', 'user29@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user30', 'user30@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user31', 'user31@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user32', 'user32@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user33', 'user33@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user34', 'user34@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user35', 'user35@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user36', 'user36@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user37', 'user37@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user38', 'user38@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user39', 'user39@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user40', 'user40@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user41', 'user41@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user42', 'user42@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user43', 'user43@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user44', 'user44@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user45', 'user45@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user46', 'user46@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user47', 'user47@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user48', 'user48@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user49', 'user49@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user50', 'user50@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user51', 'user51@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user52', 'user52@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user53', 'user53@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user54', 'user54@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user55', 'user55@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user56', 'user56@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user57', 'user57@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user58', 'user58@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user59', 'user59@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user60', 'user60@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user61', 'user61@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user62', 'user62@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user63', 'user63@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user64', 'user64@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user65', 'user65@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user66', 'user66@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user67', 'user67@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user68', 'user68@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user69', 'user69@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user70', 'user70@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user71', 'user71@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user72', 'user72@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user73', 'user73@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user74', 'user74@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user75', 'user75@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user76', 'user76@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user77', 'user77@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user78', 'user78@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user79', 'user79@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user80', 'user80@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user81', 'user81@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user82', 'user82@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user83', 'user83@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user84', 'user84@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));            
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user85', 'user85@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user86', 'user86@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user87', 'user87@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user88', 'user88@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user89', 'user89@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user90', 'user90@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user91', 'user91@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user92', 'user92@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user93', 'user93@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user94', 'user94@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user95', 'user95@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user96', 'user96@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user97', 'user97@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user98', 'user98@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user99', 'user99@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));
-- INSERT INTO users (name, email, password, email_verified_at) VALUES ('user100', 'user100@example.com', '$2a$14$v.59rKsYjjC5K.LacDrHCO/hCoXr/IZiR3HFUEG7IenOU.nV.MXfK', EXTRACT(EPOCH FROM NOW()));

-- +goose StatementEnd

-- +goose Down
DROP TABLE users;

-- +goose StatementBegin
-- DELETE FROM users WHERE email = 'admin@example.com';
-- +goose StatementEnd