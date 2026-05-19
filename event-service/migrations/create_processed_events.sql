CREATE TABLE processed_events (
    id BIGSERIAL PRIMARY KEY,
    event_id VARCHAR(100) NOT NULL UNIQUE,
    event_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);