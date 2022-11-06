package entity

const PersonageDescriptionTableName = "PersonageDescription"

type PersonageDescription struct {
	Id             int64  `json:"id,omitempty"`
	Age            string `json:"age,omitempty"`
	Height         string `json:"height,omitempty"`
	Weight         string `json:"weight,omitempty"`
	Voice          string `json:"voice,omitempty"`
	Characteristic string `json:"characteristic"`
}
