package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type AgeRatingQuery interface {
	AddRating(rating datastruct.AgeRating) (*int64, error)
	GetRating(id int64) (*datastruct.AgeRating, error)
	DeleteRating(id int64) error
}

type ageRatingQuery struct{}

func (a *ageRatingQuery) AddRating(rating datastruct.AgeRating) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.AgeRatingTableName).
		Columns("rating").
		Values(rating.Rating).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("add rating error: %w", err)
	}
	return &id, nil
}

func (a *ageRatingQuery) GetRating(id int64) (*datastruct.AgeRating, error) {
	qb := dbQueryBuilder().
		Select("rating", "id").
		From(datastruct.AgeRatingTableName).
		Where(squirrel.Eq{"id": id})

	r := datastruct.AgeRating{}
	err := qb.QueryRow().Scan(&r.Rating, &r.Id)
	if err != nil {
		return nil, fmt.Errorf("get rating error: %w", err)
	}

	return &r, nil
}

func (a *ageRatingQuery) DeleteRating(id int64) error {
	qb := dbQueryBuilder().
		Delete(datastruct.AgeRatingTableName).
		From(datastruct.AgeRatingTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return fmt.Errorf("delete rating error: %w", err)
	}

	return nil
}
