package bluebox

type HawbInfo struct {
	Hawb               *string             `json:"hawb,omitempty"`
	Shipper            *string             `json:"shipper,omitempty"`
	AccountID          *string             `json:"account_id,omitempty"`
	CustomerReference  *string             `json:"customer_reference,omitempty"`
	Co2Emission        *Co2Emission        `json:"co2_emission,omitempty"`
	SubscriptionStatus *SubscriptionStatus `json:"subscription_status,omitempty"`
	UniqueShipmentId   *string             `json:"unique_shipment_id,omitempty"`
}
