package datastruct

import "time"

const PerformanceTableName = "Performance"

type Performance struct {
	Id              int64
	CompositionId   int64
	PerformanceName string
	Description     string
	DirectorId      int64
	Date            time.Time
	Time            time.Time
}
