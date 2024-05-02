package bluebox

type DocumentParty struct {
	Party            *Party   `json:"party,omitempty"`
	PartyFunction    *string  `json:"partyFunction,omitempty"`
	DisplayedAddress []string `json:"displayedAddress,omitempty"`
	IsToBeNotified   *bool    `json:"isToBeNotified,omitempty"`
}

type Party struct {
	PartyName           *string             `json:"partyName,omitempty"`
	Address             *Address            `json:"address,omitempty"`
	PartyContactDetails *ContactDetail      `json:"partyContactDetails,omitempty"`
	IdentifyingCodes    []IdentifyingCode   `json:"identifyingCodes,omitempty"`
	TaxLegalReferences  []TaxLegalReference `json:"taxLegalReferences,omitempty"`
}
