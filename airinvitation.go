package bluebox

import (
	"fmt"
)

type InvitationKey struct {
	AccountId        string `json:"account_id"`
	UniqueShipmentId string `json:"unique_shipment_id"`
}

type InvitationInput struct {
	Expiry *DateOnly         `json:"expiry"`
	Keys   [][]InvitationKey `json:"keys"`
}

type InvitationResponse struct {
	Header      Header       `json:"header"`
	Invitations []Invitation `json:"invitations"`
}

type Invitation struct {
	Url string `json:"url"`
}

func (c *Client) AirInvitation(shipments InvitationInput) (*InvitationResponse, error) {

	response := new(InvitationResponse)

	fmt.Println(Serialize(shipments))

	err := c.Post("/cargo/shipment/invitation", shipments, nil, response)

	return response, err

}
