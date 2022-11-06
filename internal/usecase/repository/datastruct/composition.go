package entity

const CompositionTableName = "Composition"

type Composition struct {
	Id              int64
	CompositionName string
	GenreId         int64
	AgeRatingId     int64
	AuthorId        int64
}
