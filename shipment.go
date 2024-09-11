package bluebox

type Shipment struct {
	Mawb            *string    `json:"mawb,omitempty"`
	OrgIata         *string    `json:"org_iata,omitempty"`
	DstIata         *string    `json:"dst_iata,omitempty"`
	MawbTotalPieces *int       `json:"mawb_total_pieces,omitempty"`
	MawbTotalWeight *float64   `json:"mawb_total_weight,omitempty"`
	RcsTs           *string    `json:"rcs_ts,omitempty"`
	RcsPieces       *int       `json:"rcs_pieces,omitempty"`
	DlvTs           *string    `json:"dlv_ts,omitempty"`
	DlvPieces       *int       `json:"dlv_pieces,omitempty"`
	Status          *string    `json:"status,omitempty"`
	StatusTs        *string    `json:"status_ts,omitempty"`
	Mode            *int       `json:"mode,omitempty"`
	Movements       []Movement `json:"movements,omitempty"`
}

type ShipmentSubscriptionRequest struct {
	MessageHeader *MessageHeader `json:"messageHeader"`
	Shipments     []Shipment     `json:"shipments"`
}

type ShipmentSubscriptionResponse struct {
	Message string   `json:"message"`
	UIDS    []string `json:"uids"`
}

func (c *Client) SubscribeShipments(shipments []Shipment) (*ShipmentSubscriptionResponse, error) {

	request := &ShipmentSubscriptionRequest{
		MessageHeader: &MessageHeader{
			AccountID: c.AccountID,
		},
		Shipments: shipments,
	}

	response := new(ShipmentSubscriptionResponse)

	err := c.Post("/air/shipment/subscription", request, nil, response)

	return response, err

}
