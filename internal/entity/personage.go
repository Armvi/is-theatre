package entity

type Personage struct {
	Id          int64                 `json:"id,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description *PersonageDescription `json:"description,omitempty"`
}
