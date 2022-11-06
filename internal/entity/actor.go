package entity

type Actor struct {
	Id          int64             `json:"id,omitempty"`
	Worker      *Worker           `json:"worker"`
	Experience  int               `json:"experience,omitempty"`
	Description *ActorDescription `json:"description"`
}
