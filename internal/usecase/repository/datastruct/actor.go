package entity

const ActorTableName = "Actor"

type Actor struct {
	Id         int64
	WorkerId   int64
	Experience int
}
