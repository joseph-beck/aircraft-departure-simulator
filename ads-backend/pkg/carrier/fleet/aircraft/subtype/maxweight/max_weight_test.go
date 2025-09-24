package maxweight

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxWeightString(t *testing.T) {
	mw := MaxWeight{
		UUID:     "123e4567-e89b-12d3-a456-426614174000",
		Name:     "TestWeight",
		Date:     time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		ZeroFuel: 50000,
		TakeOff:  70000,
		Landing:  60000,
		Taxi:     55000,
	}

	expected := "MaxWeight{UUID: 123e4567-e89b-12d3-a456-426614174000, Name: TestWeight, Date: 2023-10-01 00:00:00 +0000 UTC, ZeroFuel: 50000, TakeOff: 70000, Landing: 60000, Taxi: 55000}"
	assert.Equal(t, expected, mw.String())
}
