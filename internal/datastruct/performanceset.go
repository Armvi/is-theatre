package datastruct

import "time"

const PerformanceSetTableName = "PerformanceSet"

type PerformanceSet struct {
	Id            int64
	RepertoireId  int64
	PerformanceId int64
	Date          time.Time
}
