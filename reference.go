package bluebox

type Reference struct {
	ReferenceType  *string `json:"referenceType,omitempty"`
	ReferenceValue *string `json:"referenceValue,omitempty"`
}
