package cargohold

import (
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/cargohold/baygroup"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/cargohold/baylocation"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/cargohold/lock"
)

type Compartment struct {
	Locks        []lock.Lock
	BayGroups    []baygroup.BayGroup
	BayLocations []baylocation.BayLocation
}
