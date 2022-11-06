package entity

type PerformanceDirector struct {
	Id     int64   `json:"id,omitempty"`
	Worker *Worker `json:"worker,omitempty"`
}
