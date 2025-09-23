package subtype

import "github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/maxweight"

var MockSubtype = Subtype{
	MaxWeights: []maxweight.MaxWeight{maxweight.MockMaxWeight},
}
