CREATE TABLE IF NOT EXISTS users_models (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name  VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(150) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users_models(email);

INSERT INTO users_models (name,email, password,created_at, updated_at)
VALUES 
('Alice Johnson','alice@example.com', '$2a$10$VQG3Sk0DsVW/qn7HXMF6nevPl0baoxx1XnoIvlh2qQSBNXlIUudeu',now(),now()), -- password : admin1234 
('Bob Smith','bob@example.com', '$2a$10$Kd.f60dC6nJUnGebiYMDW.g8b/8JTMavYN9S6Ni7Kpjfxz70aD3le',now(),now())  -- password : staff1234
ON CONFLICT DO NOTHING;