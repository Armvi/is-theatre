package entity

type Author struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Country    string `json:"country,omitempty"`
	Century    string `json:"century,omitempty"`
}
