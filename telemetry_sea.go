package bluebox

import "time"

type SeaTelemetryWebHook struct {
	Header      SeaTelemetryHeader `json:"header"`
	Telemetries []Telemetry        `json:"telemetries"`
}

type SeaTelemetryHeader struct {
	MessageTs  time.Time `json:"messageTs"`
	ApiVersion string    `json:"apiVersion"`
}

type SeaTelemetry struct {
	Header SeaTelemetryHeader `json:"header"`
	Events []Event            `json:"events"`
}

type Event struct {
	Metadata EventMetadata `json:"metadata"`
	Payload  EventPayload  `json:"payload"`
}

type EventMetadata struct {
	EventType string         `json:"eventType"`
	Publisher EventPublisher `json:"publisher"`
}

type EventPublisher struct {
	PartyName               string `json:"partyName"`
	CarrierCode             string `json:"carrierCode"`
	CarrierCodeListProvider string `json:"carrierCodeListProvider"`
}

type EventPayload struct {
	RelatedDocumentReferences []DocumentReference `json:"relatedDocumentReferences"`
	References                []DocumentReference `json:"references"`
	TransportCall             TransportCall       `json:"transportCall"`
}

type DocumentReference struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TransportCall struct {
	ModeOfTransport           string               `json:"modeOfTransport"`
	Location                  SeaLocationTelemetry `json:"location"`
	CarrierExportVoyageNumber string               `json:"carrierExportVoyageNumber"`
	Vessel                    SeaVesselTelemetry   `json:"vessel"`
}

type SeaLocationTelemetry struct {
	LocationType string  `json:"locationType"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type SeaVesselTelemetry struct {
	VesselIMONumber                 string `json:"vesselIMONumber"`
	Name                            string `json:"name"`
	OperatorCarrierCode             string `json:"operatorCarrierCode"`
	OperatorCarrierCodeListProvider string `json:"operatorCarrierCodeListProvider"`
}
