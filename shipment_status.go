package bluebox

type ShipmentStatus struct {
	HawbInfos []HawbInfo `json:"hawb_infos,omitempty"`
	MawbInfo  *MawbInfo  `json:"mawb_info,omitempty"`
}

type ShipmentStatusRequest struct {
	AccountId string `json:"accountId"`
	Mawbs     string `json:"mawbs"`
}

type ShipmentStatusResponse struct {
	Header    *Header          `json:"header,omitempty"`
	Shipments []ShipmentStatus `json:"shipments,omitempty"`
}

func (c *Client) GetShipmentStatus(mawbs string) (*ShipmentStatusResponse, error) {

	request := &ShipmentStatusRequest{
		AccountId: c.AccountID,
		Mawbs:     mawbs,
	}

	response := new(ShipmentStatusResponse)

	err := c.Get("/air/shipment/status", request, nil, response)

	return response, err

}
