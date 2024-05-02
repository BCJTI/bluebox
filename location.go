package bluebox

type Location struct {
	AddressLocation    *AddressLocation    `json:"addressLocation,omitempty"`
	UNLocationLocation *UnLocationLocation `json:"unLocationLocation,omitempty"`
}

type Address struct {
	Name         *string `json:"name,omitempty"`
	Street       *string `json:"street,omitempty"`
	StreetNumber *string `json:"streetNumber,omitempty"`
	Floor        *string `json:"floor,omitempty"`
	PostCode     *string `json:"postCode,omitempty"`
	City         *string `json:"city,omitempty"`
	StateRegion  *string `json:"stateRegion,omitempty"`
	Country      *string `json:"country,omitempty"`
}

type AddressLocation struct {
	LocationName string   `json:"locationName,omitempty"`
	LocationType string   `json:"locationType,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

type UnLocationLocation struct {
	LocationName   *string `json:"locationName,omitempty"`
	LocationType   *string `json:"locationType,omitempty"`
	UNLocationCode *string `json:"UNLocationCode,omitempty"`
}
