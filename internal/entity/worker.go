package entity

type Worker struct {
	Id         int64   `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	SecondName string  `json:"second_name,omitempty"`
	BirthDate  Date    `json:"birth_date"`
	Salary     float64 `json:"salary,omitempty"`
}
