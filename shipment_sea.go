package bluebox

type OceanSubscription struct {
	MessageHeaderOceanSub *MessageHeader         `json:"message_header,omitempty"`
	OceanSubscriptions    []OceanSubSubscription `json:"ocean_subscriptions,omitempty"`
}

type OceanSubSubscription struct {
	CommunicationChannelCode    *string                      `json:"communicationChannelCode,omitempty"`
	OceanSubCommodities         []OceanSubCommodity          `json:"commodities,omitempty"`
	OceanSubReferences          []OceanSubReference          `json:"references,omitempty"`
	OceanSubRequestedEquipments []OceanSubRequestedEquipment `json:"requestedEquipments,omitempty"`
	OceanSubDocumentParties     []OceanSubDocumentParty      `json:"documentParties,omitempty"`
}

type OceanSubCommodity struct {
	CommodityRequestedEquipmentLink *string `json:"commodityRequestedEquipmentLink,omitempty"`
}

type OceanSubReference struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type OceanSubRequestedEquipment struct {
	EquipmentReferences             []string `json:"equipmentReferences,omitempty"`
	CommodityRequestedEquipmentLink *string  `json:"commodityRequestedEquipmentLink,omitempty"`
}

type OceanSubDocumentParty struct {
	OceanSubParty *OceanSubParty `json:"party,omitempty"`
	PartyFunction *string        `json:"partyFunction,omitempty"`
}

type OceanSubParty struct {
	PartyName *string          `json:"partyName,omitempty"`
	Address   *OceanSubAddress `json:"address,omitempty"`
}

type OceanSubAddress struct {
	Name *string `json:"name,omitempty"`
}

func (c *Client) SubscribeOceanShipments(shipments []OceanSubSubscription) (*ShipmentSubscriptionResponse, error) {

	request := &OceanSubscription{
		MessageHeaderOceanSub: &MessageHeader{
			AccountID: c.AccountID,
		},
		OceanSubscriptions: shipments,
	}

	response := new(ShipmentSubscriptionResponse)

	err := c.Post("/ocean/shipment/subscription", request, nil, response)

	return response, err

}

type SeaHeader struct {
	MessageTs  *Timestamp `json:"messageTs,omitempty"`
	ApiVersion *string    `json:"apiVersion,omitempty"`
}

type SeaShipmentWebHook struct {
	Header    *SeaHeader    `json:"header,omitempty"`
	Shipments []SeaShipment `json:"shipments,omitempty"`
}

type SeaShipment struct {
	Header *SeaHeader  `json:"header,omitempty"`
	Events []SeaEvents `json:"events,omitempty"`
}

type SeaEvents struct {
	Metadata *SeaMetadata `json:"metadata,omitempty"`
	Payload  *SeaPayload  `json:"payload,omitempty"`
}

type SeaMetadata struct {
	EventType *string       `json:"eventType,omitempty"`
	Publisher *SeaPublisher `json:"publisher,omitempty"`
}

type SeaPublisher struct {
	PartyName               *string `json:"partyName,omitempty"`
	CarrierCode             *string `json:"carrierCode,omitempty"`
	CarrierCodeListProvider *string `json:"carrierCodeListProvider,omitempty"`
}

type SeaPayload struct {
	EventClassifierCode       *string                `json:"eventClassifierCode,omitempty"`
	EventDateTime             *Timestamp             `json:"eventDateTime,omitempty"`
	RelatedDocumentReferences []SeaDocumentReference `json:"relatedDocumentReferences,omitempty"`
	References                []SeaReference         `json:"references,omitempty"`
	EquipmentEventTypeCode    *string                `json:"equipmentEventTypeCode,omitempty"`
	EquipmentReference        *string                `json:"equipmentReference,omitempty"`
	EmptyIndicatorCode        *string                `json:"emptyIndicatorCode,omitempty"`
	IsTransshipmentMove       *bool                  `json:"isTransshipmentMove,omitempty"`
	Location                  *SeaLocation           `json:"location,omitempty"`
	FacilityTypeCode          *string                `json:"facilityTypeCode,omitempty"`
	TransportCall             *SeaTransportCall      `json:"transportCall,omitempty"`
	TransportEventTypeCode    *string                `json:"transportEventTypeCode,omitempty"`
	ShipmentEventTypeCode     *string                `json:"shipmentEventTypeCode,omitempty"`
}

type SeaDocumentReference struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type SeaReference struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type SeaLocation struct {
	LocationName   *string `json:"locationName"`
	LocationType   *string `json:"locationType"`
	UNLocationCode *string `json:"UNLocationCode"`
}

type SeaTransportCall struct {
	ModeOfTransport           *string      `json:"modeOfTransport"`
	Location                  *SeaLocation `json:"location"`
	CarrierExportVoyageNumber *string      `json:"carrierExportVoyageNumber,omitempty"`
	FacilityTypeCode          *string      `json:"facilityTypeCode"`
	Vessel                    *SeaVessel   `json:"vessel"`
	CarrierImportVoyageNumber *string      `json:"carrierImportVoyageNumber,omitempty"`
}

type SeaVessel struct {
	VesselIMONumber                 *string `json:"vesselIMONumber"`
	Name                            *string `json:"name"`
	OperatorCarrierCode             *string `json:"operatorCarrierCode"`
	OperatorCarrierCodeListProvider *string `json:"operatorCarrierCodeListProvider"`
}
