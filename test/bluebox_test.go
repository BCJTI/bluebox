package test

import (
	"os"
	"testing"

	"github.com/bcjti/bluebox"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
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
	spew.Dump(response)
}

func TestSubscribeShipments(t *testing.T) {
	shipments := []bluebox.Shipment{{
		Movements: nil,
	}}
	response, err := client.SubscribeShipments(shipments)
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
}

func TestGetShipmentStatus(t *testing.T) {
	response, err := client.GetShipmentStatus(os.Getenv("BB_MAWB"))
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	spew.Dump(response)
}

func TestGetShipmentTelemetries(t *testing.T) {
	response, err := client.GetShipmentTelemetries(os.Getenv("BB_MAWB"))
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
	spew.Dump(response)
}

func TestSubscribeOceanShipments(t *testing.T) {
	shipments := []bluebox.OceanSubSubscription{{
		OceanSubRequestedEquipments: []bluebox.OceanSubRequestedEquipment{{EquipmentReferences: []string{"TEMU3179317"}}},
	}}

	response, err := client.SubscribeOceanShipments(shipments)
	assert.NoError(t, err)
	assert.NotNil(t, response, "response cannot be nil")
}
