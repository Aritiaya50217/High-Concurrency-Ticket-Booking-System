CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- event_id UUID NOT NULL,
    seat_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_seats_event FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT unique_event_seat UNIQUE(event_id, seat_number)
);

-- CREATE INDEX idx_seats_event_id ON seats(event_id);

INSERT INTO seats (seat_number)
VALUES
('A1'),
('A2'),
('A3');