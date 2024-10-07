package test

import (
	"fmt"
	"github.com/bcjti/bluebox"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var client *bluebox.Client

func init() {
	accountID := os.Getenv("BB_ACCOUNT_ID")
	username := os.Getenv("BB_USERNAME")
	password := os.Getenv("BB_PASSWORD")
	client = bluebox.NewClient(accountID, username, password)
}

func TestGetSubscriptions(t *testing.T) {
	shipmentKeys := []bluebox.Subscription{{
		Hawb: bluebox.Str2Pnt(os.Getenv("BB_HAWB")),
		Mawb: bluebox.Str2Pnt(os.Getenv("BB_MAWB")),
	}}
	response, err := client.GetSubscriptions("", shipmentKeys)
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	fmt.Println(bluebox.Serialize(response))
}

func TestSubscribeShipments(t *testing.T) {
	shipments := []bluebox.Subscription{{
		Mawb:      bluebox.Str2Pnt(os.Getenv("BB_MAWB")),
		PushTypes: []string{"SHIPMENT", "TELEMETRY"},
	}}
	response, err := client.SubscribeShipments(shipments)
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	fmt.Println(bluebox.Serialize(response))
}

func TestGetShipmentStatus(t *testing.T) {
	response, err := client.GetShipmentStatus(os.Getenv("BB_MAWB"))
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	fmt.Println(bluebox.Serialize(response))
}

func TestGetShipmentTelemetries(t *testing.T) {
	response, err := client.GetShipmentTelemetries(os.Getenv("BB_MAWB"))
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	fmt.Println(bluebox.Serialize(response))
}

func TestSubscribeOceanShipments(t *testing.T) {
	//shipments := []bluebox.OceanSubSubscription{{
	//	CommunicationChannelCode: bluebox.Str2Pnt("AO"), //FIX AO (API)
	//	OceanSubCommodities: []bluebox.OceanSubCommodity{{
	//		CommodityRequestedEquipmentLink: bluebox.Str2Pnt("001"),
	//	}},
	//	OceanSubReferences: []bluebox.OceanSubReference{
	//		{
	//			Type:  bluebox.Str2Pnt("FF"),
	//			Value: bluebox.Str2Pnt("WINDT-ES24080014"),
	//		},
	//		{
	//			Type:  bluebox.Str2Pnt("CR"),
	//			Value: bluebox.Str2Pnt("SSZ1520203"),
	//		},
	//	},
	//	OceanSubRequestedEquipments: []bluebox.OceanSubRequestedEquipment{{
	//		EquipmentReferences:             []string{"TEMU3179317"},
	//		CommodityRequestedEquipmentLink: bluebox.Str2Pnt("001"),
	//	}},
	//	OceanSubDocumentParties: []bluebox.OceanSubDocumentParty{{
	//		OceanSubParty: &bluebox.OceanSubParty{
	//			PartyName: bluebox.Str2Pnt("CMA CGM - São Paulo"),
	//			Address: &bluebox.OceanSubAddress{
	//				Name: bluebox.Str2Pnt("CMA CGM - São Paulo"),
	//			},
	//		},
	//		PartyFunction: bluebox.Str2Pnt("OS"),
	//	}},
	//}}

	//timeLoad, _ := bluebox.NewTimestampPointer("2024-09-01 08:06")
	//
	//timeArrival, _ := bluebox.NewTimestampPointer("2024-10-02 08:06")

	shipments := []bluebox.OceanSubSubscription{{
		CommunicationChannelCode:   bluebox.Str2Pnt("AO"), //FIX AO (API)
		TransportDocumentTypeCode:  bluebox.Str2Pnt("BOL"),
		TransportDocumentReference: bluebox.Str2Pnt("SSZ1520203"),
		IncoTerms:                  bluebox.Str2Pnt("CIF"),
		//ExpectedArrivalAtPlaceOfDeliveryStartDate: bluebox.Str2Pnt("2024-09-01"),
		//ExpectedArrivalAtPlaceOfDeliveryEndDate:   bluebox.Str2Pnt("2024-10-02"),
		OceanSubReferences: []bluebox.OceanSubReference{
			{
				Type:  bluebox.Str2Pnt("CR"),
				Value: bluebox.Str2Pnt("WINDT-ES24080014"),
			},
			{
				Type:  bluebox.Str2Pnt("FF"),
				Value: bluebox.Str2Pnt("SSZ1520203"),
			},
			{
				Type:  bluebox.Str2Pnt("SAC"),
				Value: bluebox.Str2Pnt("CMDU"),
			},
		},
		OceanSubRequestedEquipments: []bluebox.OceanSubRequestedEquipment{{
			EquipmentReferences:             []string{"TEMU3179317"},
			CommodityRequestedEquipmentLink: bluebox.Str2Pnt("001"),
		}},
		//OceanSubDocumentParties: []bluebox.OceanSubDocumentParty{{
		//	OceanSubParty: &bluebox.OceanSubParty{
		//		PartyName: bluebox.Str2Pnt("CMA CGM - São Paulo"),
		//		Address: &bluebox.OceanSubAddress{
		//			Name: bluebox.Str2Pnt("CMA CGM - São Paulo"),
		//		},
		//	},
		//	PartyFunction: bluebox.Str2Pnt("OS"),
		//}},
		//ShipmentLocations: []bluebox.OceanSubShipmentLocation{{
		//	Location: &bluebox.OceanSubLocation{
		//		LocationName: bluebox.Str2Pnt("Belize City"),
		//		Address: &bluebox.OceanSubAddress{
		//			City:    bluebox.Str2Pnt("Belize City"),
		//			Country: bluebox.Str2Pnt("BZ"),
		//		},
		//	},
		//	ShipmentLocationTypeCode: bluebox.Str2Pnt("PDE"),
		//	//EventDateTime:            timeArrival,
		//}, {
		//	Location: &bluebox.OceanSubLocation{
		//		LocationName: bluebox.Str2Pnt("Santos"),
		//		Address: &bluebox.OceanSubAddress{
		//			City:    bluebox.Str2Pnt("Belize City"),
		//			Country: bluebox.Str2Pnt("BZ"),
		//		},
		//	},
		//	ShipmentLocationTypeCode: bluebox.Str2Pnt("PRE"),
		//	//EventDateTime:            timeLoad,
		//}},
	}}

	response, err := client.SubscribeOceanShipments(shipments)
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	fmt.Println(bluebox.Serialize(response))
}
