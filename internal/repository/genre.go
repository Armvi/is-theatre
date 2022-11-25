package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type GenreQuery interface {
	AddGenre(genre datastruct.Genre) (*int64, error)
	GetGenre(id int64) (*datastruct.Genre, error)
	DeleteGenre(id int64) error
}

type genreQuery struct{}

func (a *genreQuery) AddGenre(genre datastruct.Genre) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.GenreTableName).
		Columns("genre").
		Values(genre.GenreName).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("add genre error: %w", err)
	}
	return &id, nil
}

func (a *genreQuery) GetGenre(id int64) (*datastruct.Genre, error) {
	qb := dbQueryBuilder().
		Select("genre", "id").
		From(datastruct.GenreTableName).
		Where(squirrel.Eq{"id": id})

	g := datastruct.Genre{}
	err := qb.QueryRow().Scan(&g.GenreName, &g.Id)
	if err != nil {
		return nil, fmt.Errorf("get rating error: %w", err)
	}

	return &g, nil
}

func (a *genreQuery) DeleteGenre(id int64) error {
	qb := dbQueryBuilder().
		Delete(datastruct.GenreTableName).
		From(datastruct.GenreTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return fmt.Errorf("delete genre error: %w", err)
	}

	return nil
}
