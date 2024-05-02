package bluebox

type IdentifyingCode struct {
	CodeListProvider *string `json:"codeListProvider,omitempty"`
	PartyCode        *string `json:"partyCode,omitempty"`
	CodeListName     *string `json:"codeListName,omitempty"`
}
