package bluebox

type Vessel struct {
	VesselName      *string `json:"vesselName,omitempty"`
	VesselIMONumber *string `json:"vesselIMONumber,omitempty"`
}
