CREATE TABLE IF NOT EXISTS booking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- user_id UUID NOT NULL,
    event_id UUID NOT NULL,
    seat_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_booking_status CHECK (status IN ('PENDING', 'CONFIRMED', 'CANCELED')),
    -- CONSTRAINT fk_booking_users FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_booking_events FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT fk_booking_seats FOREIGN KEY (seat_id) REFERENCES seats(id)
);


-- CREATE INDEX idx_booking_user_id ON booking(user_id);
-- CREATE INDEX idx_booking_event_id ON booking(event_id);
CREATE INDEX idx_booking_seat_id ON booking(seat_id);

INSERT INTO booking(seat_id,status,created_at,updated_at)
VALUES (
    '90826a9c-3d90-4fd8-95ff-a38b7a683753','PENDING',now(),now()
);
