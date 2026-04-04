CREATE TABLE IF NOT EXISTS booking_models (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    event_id UUID NOT NULL,
    seat_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_booking_models_status CHECK (status IN ('PENDING', 'CONFIRMED', 'CANCELED')),
    CONSTRAINT fk_booking_models_users FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_booking_models_events FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT fk_booking_models_seats FOREIGN KEY (seat_id) REFERENCES seats(id)
);


CREATE INDEX idx_booking_models_user_id ON booking_models(user_id);
CREATE INDEX idx_booking_models_event_id ON booking_models(event_id);
CREATE INDEX idx_booking_models_seat_id ON booking_models(seat_id);

INSERT INTO booking_models(user_id,event_id,seat_id,status,created_at,updated_at)
VALUES (
    '2059fd41-0855-4822-9065-43fd834f16f8','16deccf9-ac6a-4638-89c2-6dcfb18e4a9e','90826a9c-3d90-4fd8-95ff-a38b7a683753','PENDING',now(),now()
);
