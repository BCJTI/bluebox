package bluebox

type MessageHeader struct {
	AccountID string `json:"account_id,omitempty"`
}

type Header struct {
	MessageTs  *Timestamp `json:"message_ts,omitempty"`
	ApiVersion string     `json:"api_version,omitempty"`
}
