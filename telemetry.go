package bluebox

type Telemetry struct {
	Mawb      *string          `json:"mawb,omitempty"`
	Movements *[]Movement      `json:"movements,omitempty"`
	Iatas     *[]TelemetryIata `json:"iatas,omitempty"`
}

type TelemetryInfo struct {
	TelemetryInfo *Telemetry `json:"telemetry_info,omitempty"`
}

type TelemetryIata struct {
	Iata     *string            `json:"iata,omitempty"`
	Position *TelemetryPosition `json:"position,omitempty"`
}

type TelemetryPosition struct {
	Lat *float64 `json:"lat,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
}

type ShipmentTelemetryRequest struct {
	AccountId string `json:"accountId"`
	Mawbs     string `json:"mawbs"`
}

type ShipmentTelemetryResponse struct {
	Header      *Header         `json:"header,omitempty"`
	Telemetries []TelemetryInfo `json:"telemetries,omitempty"`
}

func (c *Client) GetShipmentTelemetries(mawbs string) (*ShipmentTelemetryResponse, error) {

	request := &ShipmentTelemetryRequest{
		AccountId: c.AccountID,
		Mawbs:     mawbs,
	}

	response := new(ShipmentTelemetryResponse)

	err := c.Get("/air/shipment/telemetry", request, nil, response)

	return response, err

}
