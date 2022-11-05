package datastruct

import "time"

const RepertoireTableName = "Repertoire"

type Repertoire struct {
	Id        int64
	BeginDate time.Time
	EndDate   time.Time
}
