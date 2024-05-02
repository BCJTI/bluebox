package bluebox

type Subscription struct {
	Description           *string    `json:"description,omitempty"`
	ShipmentID            *string    `json:"shipment_id,omitempty"`
	Mawb                  *string    `json:"mawb,omitempty"`
	Hawb                  *string    `json:"hawb,omitempty"`
	CustomerReference     *string    `json:"customer_reference,omitempty"`
	SubscriptionReference *string    `json:"subscription_reference,omitempty"`
	Shipper               *string    `json:"shipper,omitempty"`
	ShipperCtrCd          *string    `json:"shipper_ctr_cd,omitempty"`
	Consignee             *string    `json:"consignee,omitempty"`
	ConsigneeCtrCd        *string    `json:"consignee_ctr_cd,omitempty"`
	Pieces                *int       `json:"pieces,omitempty"`
	HawbPieces            *int       `json:"hawb_pieces,omitempty"`
	ChargeableWeight      *float64   `json:"chargeable_weight,omitempty"`
	HawbWeight            *float64   `json:"hawb_weight,omitempty"`
	GrossWeight           *float64   `json:"gross_weight,omitempty"`
	Volume                *float64   `json:"volume,omitempty"`
	HsGoodsDescription    *string    `json:"hs_goods_description,omitempty"`
	HsCode                *int       `json:"hs_code,omitempty"`
	Forwarder             *string    `json:"forwarder,omitempty"`
	IncoTerms             *string    `json:"inco_terms,omitempty"`
	PickupCity            *string    `json:"pickup_city,omitempty"`
	PickupZip             *string    `json:"pickup_zip,omitempty"`
	PickupLocode          *string    `json:"pickup_locode,omitempty"`
	PickupLocation        *string    `json:"pickup_location,omitempty"`
	PmsPickupDt           *DateOnly  `json:"pms_pickup_dt,omitempty"`
	EstPickupDt           *DateOnly  `json:"est_pickup_dt,omitempty"`
	PmsPickupTs           *Timestamp `json:"pms_pickup_ts,omitempty"`
	ActPickupDt           *DateOnly  `json:"act_pickup_dt,omitempty"`
	ActPickupTs           *Timestamp `json:"act_pickup_ts,omitempty"`
	PmsDepartureDt        *DateOnly  `json:"pms_departure_dt,omitempty"`
	PmsDepartureTs        *Timestamp `json:"pms_departure_ts,omitempty"`
	PmsArrivalDt          *DateOnly  `json:"pms_arrival_dt,omitempty"`
	PmsArrivalTs          *Timestamp `json:"pms_arrival_ts,omitempty"`
	DeliveryCity          *string    `json:"delivery_city,omitempty"`
	DeliveryZip           *string    `json:"delivery_zip,omitempty"`
	DeliveryLocode        *string    `json:"delivery_locode,omitempty"`
	DeliveryLocation      *string    `json:"delivery_location,omitempty"`
	PmsDeliveryDt         *DateOnly  `json:"pms_delivery_dt,omitempty"`
	EstDeliveryDt         *DateOnly  `json:"est_delivery_dt,omitempty"`
	PmsDeliveryTs         *Timestamp `json:"pms_delivery_ts,omitempty"`
	ActDeliveryDt         *DateOnly  `json:"act_delivery_dt,omitempty"`
	ActDeliveryTs         *Timestamp `json:"act_delivery_ts,omitempty"`
	Mode                  *int       `json:"mode,omitempty"`
	Status                *string    `json:"status,omitempty"`
	PushTypes             []string   `json:"push_types,omitempty"`
}

type SubscriptionStatus struct {
	Status *string `json:"status,omitempty"`
}

type SubscriptionRequest struct {
	AccountID    string         `json:"accountId"`
	RequestRefs  string         `json:"requestRefs,omitempty"`
	ShipmentKeys []Subscription `json:"shipmentKeys,omitempty"`
}

type SubscriptionResponse struct {
	Header        *Header        `json:"header"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (c *Client) GetSubscriptions(requestRefs string, shipmentKeys []Subscription) (*SubscriptionResponse, error) {

	request := &SubscriptionRequest{
		AccountID:    c.AccountID,
		RequestRefs:  requestRefs,
		ShipmentKeys: shipmentKeys,
	}

	response := new(SubscriptionResponse)

	err := c.Get("shipment/subscription", request, nil, response)

	return response, err

}
