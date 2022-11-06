package entity

type Repertoire struct {
	Id           int64         `json:"id,omitempty"`
	BeginDate    Date          `json:"begin_date"`
	EndDate      Date          `json:"end_date"`
	Performances []Performance `json:"performances"`
}
