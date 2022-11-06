package entities

const ActorsRoleTableName = "ActorsRole"

type ActorsRole struct {
	Id            int64
	PerformanceId int64
	ActorId       int64
	PersonageId   int64
}
