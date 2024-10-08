package bluebox

import "fmt"

type ShipmentSubscription struct {
	UniqueShipmentID   *string  `json:"unique_shipment_id"`
	Mawb               *string  `json:"mawb"`
	Hawb               *string  `json:"hawb"`
	CustomerReference  *string  `json:"customer_reference"`
	Shipper            *string  `json:"shipper"`
	ShipperCtrCd       *string  `json:"shipper_ctr_cd"`
	Consignee          *string  `json:"consignee"`
	ConsigneeCtrCd     *string  `json:"consignee_ctr_cd"`
	Pieces             *int     `json:"pieces"`
	HawbPieces         *int     `json:"hawb_pieces"` // Deprecated
	ChargeableWeight   *float64 `json:"chargeable_weight"`
	HawbWeight         *float64 `json:"hawb_weight"` // Deprecated
	GrossWeight        *float64 `json:"gross_weight"`
	Volume             *float64 `json:"volume"`
	HsGoodsDescription *string  `json:"hs_goods_description"`
	HsCode             *int     `json:"hs_code"`
	Forwarder          *string  `json:"forwarder"`
	IncoTerms          *string  `json:"inco_terms"`
	PickupCity         *string  `json:"pickup_city"`
	PickupZip          *string  `json:"pickup_zip"`
	PickupLocode       *string  `json:"pickup_locode"`
	PickupLocation     *string  `json:"pickup_location"`
	PmsPickupDt        *string  `json:"pms_pickup_dt"`
	EstPickupDt        *string  `json:"est_pickup_dt"` // Deprecated
	PmsPickupTs        *string  `json:"pms_pickup_ts"`
	ActPickupDt        *string  `json:"act_pickup_dt"`
	ActPickupTs        *string  `json:"act_pickup_ts"`
	PmsDepartureDt     *string  `json:"pms_departure_dt"`
	PmsDepartureTs     *string  `json:"pms_departure_ts"`
	PmsArrivalDt       *string  `json:"pms_arrival_dt"`
	PmsArrivalTs       *string  `json:"pms_arrival_ts"`
	DeliveryCity       *string  `json:"delivery_city"`
	DeliveryZip        *string  `json:"delivery_zip"`
	DeliveryLocode     *string  `json:"delivery_locode"`
	DeliveryLocation   *string  `json:"delivery_location"`
	PmsDeliveryDt      *string  `json:"pms_delivery_dt"`
	EstDeliveryDt      *string  `json:"est_delivery_dt"` // Deprecated
	PmsDeliveryTs      *string  `json:"pms_delivery_ts"`
	ActDeliveryDt      *string  `json:"act_delivery_dt"`
	ActDeliveryTs      *string  `json:"act_delivery_ts"`
	Mode               *int     `json:"mode"`
	PushTypes          []string `json:"push_types"`
}

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
	MessageHeader *MessageHeader `json:"message_header"`
	Shipments     []Subscription `json:"shipments"`
}

type ShipmentSubscriptionResponse struct {
	Message string   `json:"message"`
	UIDS    []string `json:"uids"`
}

func (c *Client) SubscribeAirShipments(shipments []Subscription) (*ShipmentSubscriptionResponse, error) {

	request := &ShipmentSubscriptionRequest{
		MessageHeader: &MessageHeader{
			AccountID: c.AccountID,
		},
		Shipments: shipments,
	}

	response := new(ShipmentSubscriptionResponse)

	fmt.Println(Serialize(request))

	err := c.Post("/air/shipment/subscription", request, nil, response)

	return response, err

}
