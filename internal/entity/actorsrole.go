package entity

type ActorsRole struct {
	Id        int64      `json:"id,omitempty"`
	Actor     *Actor     `json:"actor,omitempty"`
	Personage *Personage `json:"personage,omitempty"`
}
