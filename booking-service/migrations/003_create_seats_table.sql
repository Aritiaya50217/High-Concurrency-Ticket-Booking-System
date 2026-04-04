CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL,
    seat_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_seats_event FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT unique_event_seat UNIQUE(event_id, seat_number)
);

CREATE INDEX idx_seats_event_id ON seats(event_id);

INSERT INTO seats (event_id, seat_number)
VALUES
('16deccf9-ac6a-4638-89c2-6dcfb18e4a9e', 'A1'),
('721d46fc-7c8d-4fed-aac1-bf52d03318d9', 'A2'),
('8115f221-1ca7-4564-aa02-7ae637bffa58', 'A3');