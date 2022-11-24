package entity

const PersonageTableName = "Personage"

type Personage struct {
	Id            int64
	Name          string
	CompositionId int64
}
