package bluebox

import (
	"errors"
	"fmt"
)

// enum DateType ACT (Actual) PLN (Planned) EST (Estimated)

type DateTypeEnum string

var (
	DateTypeActual    DateTypeEnum = "ACT"
	DateTypePlanned   DateTypeEnum = "PLN"
	DateTypeEstimated DateTypeEnum = "EST"
)

func (e DateTypeEnum) String() string {
	return string(e)
}

type EventTypeEnum string

var (
	EventTypeArrived  EventTypeEnum = "ARRI"
	EventTypeDeparted EventTypeEnum = "DEPA"
)

func (e EventTypeEnum) String() string {
	return string(e)
}

type OceanSubscription struct {
	MessageHeaderOceanSub *MessageHeader         `json:"message_header,omitempty"`
	OceanSubscriptions    []OceanSubSubscription `json:"ocean_subscriptions,omitempty"`
}

type OceanSubSubscription struct {
	CommunicationChannelCode                  *string                      `json:"communicationChannelCode,omitempty"`                  // for Subscribe single shipment
	ExpectedArrivalAtPlaceOfDeliveryStartDate *string                      `json:"expectedArrivalAtPlaceOfDeliveryStartDate,omitempty"` // for Subscribe with master bill of lading
	ExpectedArrivalAtPlaceOfDeliveryEndDate   *string                      `json:"expectedArrivalAtPlaceOfDeliveryEndDate,omitempty"`   // for Subscribe with master bill of lading
	TransportDocumentTypeCode                 *string                      `json:"transportDocumentTypeCode,omitempty"`                 // for Subscribe with master bill of lading
	TransportDocumentReference                *string                      `json:"transportDocumentReference,omitempty"`                // for Subscribe with master bill of lading
	IncoTerms                                 *string                      `json:"incoTerms,omitempty"`                                 // for Subscribe with master bill of lading
	OceanSubCommodities                       []OceanSubCommodity          `json:"commodities,omitempty"`                               // for Both Subscribes
	OceanSubReferences                        []OceanSubReference          `json:"references,omitempty"`                                // for Both Subscribes
	OceanSubRequestedEquipments               []OceanSubRequestedEquipment `json:"requestedEquipments,omitempty"`                       // for Subscribe single shipment
	OceanSubDocumentParties                   []OceanSubDocumentParty      `json:"documentParties,omitempty"`                           // for Both Subscribes
	ShipmentLocations                         []OceanSubShipmentLocation   `json:"shipmentLocations,omitempty"`                         // for Subscribe with master bill of lading
}

type OceanSubCommodity struct {
	CommodityRequestedEquipmentLink *string `json:"commodityRequestedEquipmentLink,omitempty"` // for Subscribe single shipment
	CommodityType                   *string `json:"commodityType,omitempty"`                   // for Subscribe with master bill of lading
	CargoGrossWeight                *int    `json:"cargoGrossWeight,omitempty"`                // for Subscribe with master bill of lading
	NumberOfPackages                *int    `json:"numberOfPackages,omitempty"`                // for Subscribe with master bill of lading
}

type OceanSubReference struct {
	Type  *string `json:"type,omitempty"`  // for Both Subscribes
	Value *string `json:"value,omitempty"` // for Both Subscribes
}

type OceanSubRequestedEquipment struct {
	EquipmentReferences             []string `json:"equipmentReferences,omitempty"`             // for Both Subscribes
	CommodityRequestedEquipmentLink *string  `json:"commodityRequestedEquipmentLink,omitempty"` // for Both Subscribes
}

type OceanSubDocumentParty struct {
	OceanSubParty *OceanSubParty `json:"party,omitempty"`         // for Both Subscribes
	PartyFunction *string        `json:"partyFunction,omitempty"` // for Both Subscribes
}

type OceanSubParty struct {
	PartyName *string          `json:"partyName,omitempty"` // for Both Subscribes
	Address   *OceanSubAddress `json:"address,omitempty"`   // for Both Subscribes
}

type OceanSubAddress struct {
	Name    *string `json:"name,omitempty"`    // for Subscribe single shipment
	City    *string `json:"city,omitempty"`    // for Subscribe with master bill of lading
	Country *string `json:"country,omitempty"` // for Subscribe with master bill of lading
}

type OceanSubShipmentLocation struct {
	Location                 *OceanSubLocation `json:"location,omitempty"`                 // for Subscribe with master bill of lading
	ShipmentLocationTypeCode *string           `json:"shipmentLocationTypeCode,omitempty"` // for Subscribe with master bill of lading
	EventDateTime            *Timestamp        `json:"eventDateTime,omitempty"`            // for Subscribe with master bill of lading
}

// Estrutura que representa uma localização.
type OceanSubLocation struct {
	LocationName *string          `json:"locationName"` // for Subscribe with master bill of lading
	Address      *OceanSubAddress `json:"address"`      // for Subscribe with master bill of lading
}

func GetCarrierNaftaCode(carrier string) (*string, error) {
	if code, ok := Carriers[carrier]; ok {
		return &code.NMFTACode, nil
	}
	return nil, errors.New("Carrier not found " + carrier)
}

func (c *Client) SubscribeOceanShipments(shipments []OceanSubSubscription) (*ShipmentSubscriptionResponse, error) {

	request := &OceanSubscription{
		MessageHeaderOceanSub: &MessageHeader{
			AccountID: c.AccountID,
		},
		OceanSubscriptions: shipments,
	}

	fmt.Println(Serialize(request))

	response := new(ShipmentSubscriptionResponse)

	err := c.Post("/ocean/shipment/subscription", request, nil, response)

	return response, err

}

type SeaHeader struct {
	MessageTs   *Timestamp   `json:"messageTs,omitempty"`
	Co2Emission *Co2Emission `json:"co2Emission,omitempty"`
	ApiVersion  *string      `json:"apiVersion,omitempty"`
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
