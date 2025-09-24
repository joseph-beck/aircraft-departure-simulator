package maxweight

import (
	"fmt"
	"time"

	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/generated/sqlc"
)

// Maximum weight configurations for an aircraft.
type MaxWeight struct {
	// Unique identifier for the maximum weight entry.
	UUID string `json:"uuid"`
	// Name of the maximum weight.
	Name string `json:"name"`
	// Date for which the data is effective.
	Date time.Time `json:"date"`
	// Maximum weight, when the aircraft has no fuel.
	ZeroFuel uint `json:"zero_fuel"`
	// Maximum weight for aircraft take off.
	TakeOff uint `json:"take_off"`
	// Maximum weight the aircraft can land at.
	Landing uint `json:"landing"`
	// Maximum weight for taxiing the aircraft.
	Taxi uint `json:"taxi"`
}

// More readable string representation of MaxWeight.
func (m MaxWeight) String() string {
	return fmt.Sprintf(
		"MaxWeight{UUID: %s, Name: %s, Date: %s, ZeroFuel: %d, TakeOff: %d, Landing: %d, Taxi: %d}",
		m.UUID,
		m.Name,
		m.Date,
		m.ZeroFuel,
		m.TakeOff,
		m.Landing,
		m.Taxi,
	)
}

// Preparer for converting sqlc generated MaximumWeight to MaxWeight.
// Implements the Preparer interface.
type MaxWeightPreparer struct{}

// Prepare converts a sqlc.MaximumWeight to a MaxWeight.
func (MaxWeightPreparer) Prepare(m sqlc.MaximumWeight) (MaxWeight, error) {
	return MaxWeight{
		UUID:     m.Uuid.String(),
		Name:     m.Name,
		Date:     time.Now(),
		ZeroFuel: uint(m.ZeroFuel),
		TakeOff:  uint(m.TakeOff),
		Landing:  uint(m.Landing),
		Taxi:     uint(m.Taxi),
	}, nil
}
