package bluebox

type MawbInfo struct {
	Mawb            *string    `json:"mawb,omitempty"`
	OrgIata         *string    `json:"org_iata,omitempty"`
	DstIata         *string    `json:"dst_iata,omitempty"`
	MawbTotalPieces *int       `json:"mawb_total_pieces,omitempty"`
	MawbTotalWeight *float64   `json:"mawb_total_weight,omitempty"`
	RcsTs           *Timestamp `json:"rcs_ts,omitempty"`
	RcsPieces       *int       `json:"rcs_pieces,omitempty"`
	DlvTs           *Timestamp `json:"dlv_ts,omitempty"`
	DlvPieces       *int       `json:"dlv_pieces,omitempty"`
	Status          *string    `json:"status,omitempty"`
	StatusTs        *Timestamp `json:"status_ts,omitempty"`
	PdcArrivalTs    *Timestamp `json:"pdc_arrival_ts,omitempty"`
	Movements       []Movement `json:"movements,omitempty"`
}
