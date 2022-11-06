package entity

type Genre struct {
	Id        int64  `json:"id,omitempty"`
	GenreName string `json:"genre_name,omitempty"`
}
