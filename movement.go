package bluebox

type Movement struct {
	Description      *string    `json:"description,omitempty"`
	DepIata          *string    `json:"dep_iata,omitempty"`
	ArrIata          *string    `json:"arr_iata,omitempty"`
	FlightNo         *string    `json:"flight_no,omitempty"`
	VehicleType      *string    `json:"vehicle_type,omitempty"`
	MawbFlightPieces *int       `json:"mawb_flight_pieces,omitempty"`
	MawbFlightWeight *float64   `json:"mawb_flight_weight,omitempty"`
	BkdTs            *Timestamp `json:"bkd_ts,omitempty"`
	RcfTs            *Timestamp `json:"rcf_ts,omitempty"`
	ManTs            *Timestamp `json:"man_ts,omitempty"`
	ActDepartureTs   *Timestamp `json:"act_departure_ts,omitempty"`
	ActArrivalTs     *Timestamp `json:"act_arrival_ts,omitempty"`
	SkdDepartureTs   *Timestamp `json:"skd_departure_ts,omitempty"`
	SkdDepartureDt   *DateOnly  `json:"skd_departure_dt,omitempty"`
	SkdArrivalTs     *Timestamp `json:"skd_arrival_ts,omitempty"`
	EstDepartureTs   *Timestamp `json:"est_departure_ts,omitempty"`
	EstArrivalTs     *Timestamp `json:"est_arrival_ts,omitempty"`
	Status           *string    `json:"status,omitempty"`
}
