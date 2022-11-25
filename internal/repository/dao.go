package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type DAO interface {
	NewUserQuery() UserQuery
	NewAuthorQuery() AuthorQuery
	NewAgeRatingQuery() AgeRatingQuery
	NewGenreQuery() GenreQuery
	NewCompositionQuery() CompositionQuery
	NewPersonageQuery() PersonageQuery
}

type dao struct{}

var DB *sql.DB

func dbQueryBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}

func NewDAO(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewAuthorQuery() AuthorQuery {
	return &authorQuery{}
}

func (d *dao) NewAgeRatingQuery() AgeRatingQuery {
	return &ageRatingQuery{}
}

func (d *dao) NewGenreQuery() GenreQuery {
	return &genreQuery{}
}

func (d *dao) NewCompositionQuery() CompositionQuery {
	return &compositionQuery{}
}

func (d *dao) NewPersonageQuery() PersonageQuery {
	return &personageQuery{}
}
