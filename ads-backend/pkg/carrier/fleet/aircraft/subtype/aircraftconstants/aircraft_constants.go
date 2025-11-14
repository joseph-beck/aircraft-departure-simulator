package aircraftconstants

import (
	"fmt"

	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/generated/sqlc"
)

// AircraftConstants for an aircraft subtype.
type AircraftConstant struct {
	// Unique identifier for the aircraft constant entry.
	UUID string `json:"uuid"`
	// Name of the aircraft constant.
	Name string `json:"name"`
	// C constant, or multiplier.
	C uint `json:"c"`
	// K constant, or offset value.
	K uint `json:"k"`
	// Reference station for weight and balance calculations.
	ReferenceStation uint `json:"reference_station"`
	// Length of Mean Aerodynamic Chord.
	LengthOfMAC uint `json:"length_of_mac"`
	// Leading edge of Mean Aerodynamic Chord.
	LeadingEdgeMAC uint `json:"leading_edge_mac"`
}

// More readable string representation of AircraftConstant.
func (ac AircraftConstant) String() string {
	return fmt.Sprintf(
		"AircraftConstant{UUID: %s, Name: %s, C: %d, K: %d, ReferenceStation: %d, LengthOfMAC: %d, LeadingEdgeMAC: %d}",
		ac.UUID,
		ac.Name,
		ac.C,
		ac.K,
		ac.ReferenceStation,
		ac.LengthOfMAC,
		ac.LeadingEdgeMAC,
	)
}

// Preparer for converting sqlc generated AircraftConstant to AircraftConstant.
// Implements the Preparer interface.
type AircraftConstantPreparer struct{}

// Prepare converts a sqlc.AircraftConstant to an AircraftConstant.
func (AircraftConstantPreparer) Prepare(ac sqlc.AircraftConstant) (AircraftConstant, error) {
	return AircraftConstant{
		UUID:             ac.Uuid.String(),
		Name:             ac.Name,
		C:                uint(ac.C),
		K:                uint(ac.K),
		ReferenceStation: uint(ac.ReferenceStation),
		LengthOfMAC:      uint(ac.LengthOfMac),
		LeadingEdgeMAC:   uint(ac.LeadingEdgeMac),
	}, nil
}
