package maxweight

import "time"

var MockMaxWeight = MaxWeight{
	Name:     "Test Max Weight",
	Date:     time.Now(),
	ZeroFuel: 40000,
	TakeOff:  70000,
	Landing:  60000,
	Taxi:     72000,
}
