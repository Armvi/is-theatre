package entity

type Ticket struct {
	PerformanceId int64   `json:"performance,omitempty"`
	PlaceNumber   int64   `json:"place_number,omitempty"`
	Cost          float64 `json:"cost,omitempty"`
	UserId        int64   `json:"user_id"`
}
