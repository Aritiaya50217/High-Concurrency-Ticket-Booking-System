package aggregate_test

import (
	"testing"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	bookingAggregate "github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/aggregate"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/event"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/booking-service/internal/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewBooking(t *testing.T) {

	booking := bookingAggregate.NewBooking(1, 1, 10)

	assert.NotNil(t, booking)

	assert.Equal(t, uint(1), booking.UserID)

	assert.Equal(t, uint(1), booking.EventID)

	assert.Equal(t, uint(10), booking.SeatID)

	assert.Equal(t, valueobject.BookingPending, booking.Status)

	events := booking.Events()

	assert.Len(t, events, 1)

	_, ok := events[0].(event.BookingCreated)

	assert.True(t, ok)

}

func TestBooking_Confirm_Success(t *testing.T) {
	booking := aggregate.NewBooking(1, 100, 50)

	err := booking.Confirm()

	assert.NoError(t, err)

	assert.Equal(t, valueobject.BookingConfirmed, booking.Status)

	events := booking.Events()

	assert.Len(t, events, 2)

	_, ok := events[1].(event.BookingConfirmed)

	assert.True(t, ok)

}

func TestBooking_Confirm_AlreadyConfirmd(t *testing.T) {
	booking := aggregate.NewBooking(1, 100, 50)

	err := booking.Confirm()

	assert.NoError(t, err)

	before := len(booking.Events())

	err = booking.Confirm()

	assert.NoError(t, err)

	after := len(booking.Events())

	assert.Equal(t, before, after)
}

func TestBooking_Confirm_InvalidState(t *testing.T) {
	booking := aggregate.NewBooking(1, 100, 50)

	booking.Status = valueobject.BookingStatus("INVALID")

	err := booking.Confirm()

	assert.Error(t, err)

	if err != nil {
		assert.Equal(t, "booking must be PENDING to confirm", err.Error())
	}

	assert.Equal(t, valueobject.BookingStatus("INVALID"), booking.Status)
}

func TestBooking_ClearEvents(t *testing.T) {
	booking := aggregate.NewBooking(1, 100, 50)

	assert.NotEmpty(t, booking.Events())

	booking.ClearEvents()

	assert.Empty(t, booking.Events())
}

func TestBooking_AddEvent(t *testing.T) {
	booking := aggregate.NewBooking(1, 100, 50)

	before := len(booking.Events())

	e := event.NewBookingConfirmed(1, 1, 100, 50)

	booking.AddEvent(e)

	after := len(booking.Events())

	assert.Equal(t, before+1, after)
}
