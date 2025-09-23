package aircraft

import (
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/cabin"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/registration"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype"
)

type Aircraft struct {
	Subtype       subtype.Subtype
	Registrations []registration.Registration
	Cabins        []cabin.Cabin
}
