package entities

const PerformanceTableName = "Performance"

type Performance struct {
	Id                    int64
	PerformanceName       string
	CompositionId         int64
	PerformanceDirectorId int64
	Description           string
}
