package bluebox

import "time"

type SeaHeader struct {
	MessageTs  time.Time `json:"messageTs"`
	ApiVersion string    `json:"apiVersion"`
}

type SeaShipmentWebHook struct {
	Header    SeaHeader     `json:"header"`
	Shipments []SeaShipment `json:"shipments"`
}

type SeaShipment struct {
	Header SeaHeader   `json:"header"`
	Events []SeaEvents `json:"events"`
}

type SeaEvents struct {
	Metadata SeaMetadata `json:"metadata"`
	Payload  SeaPayload  `json:"payload"`
}

type SeaMetadata struct {
	EventType string       `json:"eventType"`
	Publisher SeaPublisher `json:"publisher"`
}

type SeaPublisher struct {
	PartyName               string `json:"partyName"`
	CarrierCode             string `json:"carrierCode"`
	CarrierCodeListProvider string `json:"carrierCodeListProvider"`
}

type SeaPayload struct {
	EventClassifierCode       string                 `json:"eventClassifierCode"`
	EventDateTime             time.Time              `json:"eventDateTime"`
	RelatedDocumentReferences []SeaDocumentReference `json:"relatedDocumentReferences"`
	References                []SeaReference         `json:"references"`
	EquipmentEventTypeCode    string                 `json:"equipmentEventTypeCode,omitempty"`
	EquipmentReference        string                 `json:"equipmentReference,omitempty"`
	EmptyIndicatorCode        string                 `json:"emptyIndicatorCode,omitempty"`
	IsTransshipmentMove       bool                   `json:"isTransshipmentMove,omitempty"`
	Location                  SeaLocation            `json:"location,omitempty"`
	FacilityTypeCode          string                 `json:"facilityTypeCode,omitempty"`
	TransportCall             SeaTransportCall       `json:"transportCall,omitempty"`
	TransportEventTypeCode    string                 `json:"transportEventTypeCode,omitempty"`
	ShipmentEventTypeCode     string                 `json:"shipmentEventTypeCode,omitempty"`
}

type SeaDocumentReference struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SeaReference struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SeaLocation struct {
	LocationName   string `json:"locationName"`
	LocationType   string `json:"locationType"`
	UNLocationCode string `json:"UNLocationCode"`
}

type SeaTransportCall struct {
	ModeOfTransport           string      `json:"modeOfTransport"`
	Location                  SeaLocation `json:"location"`
	CarrierExportVoyageNumber string      `json:"carrierExportVoyageNumber,omitempty"`
	FacilityTypeCode          string      `json:"facilityTypeCode"`
	Vessel                    SeaVessel   `json:"vessel"`
	CarrierImportVoyageNumber string      `json:"carrierImportVoyageNumber,omitempty"`
}

type SeaVessel struct {
	VesselIMONumber                 string `json:"vesselIMONumber"`
	Name                            string `json:"name"`
	OperatorCarrierCode             string `json:"operatorCarrierCode"`
	OperatorCarrierCodeListProvider string `json:"operatorCarrierCodeListProvider"`
}
