package datastruct

const TicketTableName = "Ticket"

type Ticket struct {
	PerformanceSetId int64
	PlaceNumber      int64
	Cost             float64
	UserId           int64
	Id               int64
}
