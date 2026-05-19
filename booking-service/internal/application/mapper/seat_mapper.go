package mapper

import (
	"encoding/json"

	domainEvent "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/interface/dto"
)

func ToSeatReserved(data json.RawMessage) (domainEvent.SeatReserved, error) {
	var d dto.SeatReservedEvent

	if err := json.Unmarshal(data, &d); err != nil {
		return domainEvent.SeatReserved{}, err
	}

	return domainEvent.SeatReserved{
		BookingID: d.BookingID,
		SeatID:    d.SeatID,
		EventID:   d.EventID,
	}, nil
}
