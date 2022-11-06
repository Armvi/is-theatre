package entity

type Composition struct {
	Id              int64      `json:"id,omitempty"`
	CompositionName string     `json:"composition_name,omitempty"`
	Genre           *Genre     `json:"genre,omitempty"`
	AgeRating       *AgeRating `json:"age_rating,omitempty"`
	Author          *Author    `json:"author,omitempty"`
}
