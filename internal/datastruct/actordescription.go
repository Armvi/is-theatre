package entity

const ActorDescriptionTableName = "ActorDescription"

type ActorDescription struct {
	Id      int64
	ActorId int64
	Age     string
	Height  string
	Weight  string
	Voice   string
}
