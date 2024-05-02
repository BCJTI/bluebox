package bluebox

type TaxLegalReference struct {
	TaxLegalReferenceType  *string `json:"taxLegalReferenceType,omitempty"`
	CountryCode            *string `json:"countryCode,omitempty"`
	TaxLegalReferenceValue *string `json:"taxLegalReferenceValue,omitempty"`
}
