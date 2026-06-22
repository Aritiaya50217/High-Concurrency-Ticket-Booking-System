CREATE TABLE IF NOT EXISTS booking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seat_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expire_at TIMESTAMP NULL

    CONSTRAINT chk_booking_status CHECK (status IN ('PENDING', 'CONFIRMED', 'CANCELED')),
);

CREATE INDEX idx_booking_seat_id ON booking(seat_id);

INSERT INTO booking(seat_id,status,created_at,updated_at)
VALUES (
    '90826a9c-3d90-4fd8-95ff-a38b7a683753','PENDING',now(),now()
);
