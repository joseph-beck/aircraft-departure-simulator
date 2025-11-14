package aircraftconstants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAircraftConstantString(t *testing.T) {
	ac := AircraftConstant{
		UUID:             "123e4567-e89b-12d3-a456-426614174001",
		Name:             "TestAircraft",
		C:                45,
		K:                90,
		ReferenceStation: 30,
		LengthOfMAC:      60,
		LeadingEdgeMAC:   70,
	}

	expected := "AircraftConstant{UUID: 123e4567-e89b-12d3-a456-426614174001, Name: TestAircraft, C: 45, K: 90, ReferenceStation: 30, LengthOfMAC: 60, LeadingEdgeMAC: 70}"
	assert.Equal(t, expected, ac.String())
}
