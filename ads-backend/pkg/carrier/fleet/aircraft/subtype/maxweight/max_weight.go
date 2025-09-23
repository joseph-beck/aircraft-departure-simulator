package maxweight

import (
	"time"
)

// Maximum weight configurations for an aircraft.
type MaxWeight struct {
	// Unique identifier for the maximum weight entry.
	ID uint `json:"id"`
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
