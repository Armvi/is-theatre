package datastruct

const ActorDescriptionTableName = "ActorDescription"

type ActorDescription struct {
	Id     int64
	Age    string
	Voice  string
	Height string
	Weight string
	Gender bool
}
