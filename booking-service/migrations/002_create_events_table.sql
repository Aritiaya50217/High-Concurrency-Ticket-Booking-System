CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    location VARCHAR(200),
    event_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_events_event_date ON events(event_date);

INSERT INTO events (name, location, event_date)
VALUES 
('Rock Concert', 'Bangkok Arena', '2026-06-01 19:00:00'),
('Tech Conference', 'BITEC Bangna', '2026-07-10 09:00:00'),
('Startup Meetup', 'True Digital Park', '2026-05-15 18:30:00');