package datastruct

const CompositionTableName = "Composition"

type Composition struct {
	Id              int64
	CompositionName string
	Description     string
	AuthorId        int64
	GenreId         int64
	AgeRatingId     int64
}
