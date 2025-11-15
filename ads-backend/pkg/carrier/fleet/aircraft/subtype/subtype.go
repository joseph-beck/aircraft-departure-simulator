package subtype

import "github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/maxweight"

type Subtype struct {
	// Unique identifier for the aircraft subtype.
	ID uint `json:"id"`
	// Name of the aircraft subtype.
	Name string `json:"name"`

	// Maximum weight configurations applicable to this aircraft subtype.
	MaxWeights []maxweight.MaxWeight `json:"max_weights"`
}
