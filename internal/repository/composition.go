package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type CompositionQuery interface {
	CreateComposition(composition datastruct.Composition) (*int64, error)
	GetComposition(id int64) (*datastruct.Composition, error)
	DeleteComposition(id int64) error
	UpdatedComposition(composition datastruct.Composition) (*datastruct.Composition, error)
	CompositionsByAuthor(authorId int64) ([]datastruct.Composition, error)
	CompositionsByRating(ratingId int64) ([]datastruct.Composition, error)
}

type compositionQuery struct{}

func (c *compositionQuery) CreateComposition(composition datastruct.Composition) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.CompositionTableName).
		Columns("compositionName",
			"description",
			"authorId",
			"genreId",
			"ageRatingId").
		Values(composition.CompositionName,
			composition.Description,
			composition.AuthorId,
			composition.GenreId,
			composition.AgeRatingId).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create composition error: %w", err)
	}

	return &id, nil
}

func (c *compositionQuery) GetComposition(id int64) (*datastruct.Composition, error) {
	db := dbQueryBuilder().
		Select("compositionName",
			"description",
			"authorId",
			"genreId",
			"ageRatingId",
			"id").
		From(datastruct.CompositionTableName).
		Where(squirrel.Eq{"id": id})

	composition := datastruct.Composition{}
	err := db.QueryRow().Scan(&composition.CompositionName,
		&composition.Description,
		&composition.AuthorId,
		&composition.GenreId,
		&composition.AgeRatingId,
		&composition.Id)
	if err != nil {
		return nil, fmt.Errorf("get compostion error: %w", err)
	}

	return &composition, nil
}

func (c *compositionQuery) DeleteComposition(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.CompositionTableName).
		From(datastruct.CompositionTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete composition error: %w", err)
	}
	return nil
}

func (c *compositionQuery) UpdatedComposition(composition datastruct.Composition) (*datastruct.Composition, error) {
	fromDB, err := c.GetComposition(composition.Id)
	if err != nil {
		return nil, fmt.Errorf("updated composition error: %w", err)
	}

	updated := updateComposition(fromDB, &composition)

	qb := dbQueryBuilder().
		Update(datastruct.CompositionTableName).
		SetMap(map[string]interface{}{
			"compositionName": updated.CompositionName,
			"description":     updated.Description,
			"authorId":        updated.AuthorId,
			"genreId":         updated.GenreId,
			"ageRatingId":     updated.AgeRatingId,
		}).Where(squirrel.Eq{"id": composition.Id}).
		Suffix("RETURNING compositionName, description, authorId, genreId, ageRatingId, id")

	updatedComposition := datastruct.Composition{}
	err = qb.QueryRow().Scan(&updatedComposition.CompositionName,
		&updatedComposition.Description,
		&updatedComposition.AuthorId,
		&updatedComposition.GenreId,
		&updatedComposition.AgeRatingId,
		&updatedComposition.Id)
	if err != nil {
		return nil, fmt.Errorf("update compostion error: %w", err)
	}

	return &composition, nil
}

func (c *compositionQuery) CompositionsByAuthor(authorId int64) ([]datastruct.Composition, error) {
	db := dbQueryBuilder().
		Select("compositionName",
			"description",
			"authorId",
			"genreId",
			"ageRatingId",
			"id").
		From(datastruct.CompositionTableName).
		Where(squirrel.Eq{"authorId": authorId})

	var compositions []datastruct.Composition
	var comp datastruct.Composition
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by country error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&comp.CompositionName,
			&comp.Description,
			&comp.AuthorId,
			&comp.GenreId,
			&comp.AgeRatingId,
			&comp.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by country error: %w", err)
		}
		compositions = append(compositions, comp)
	}

	return compositions, nil
}

func (c *compositionQuery) CompositionsByRating(ratingId int64) ([]datastruct.Composition, error) {
	db := dbQueryBuilder().
		Select("compositionName",
			"description",
			"authorId",
			"genreId",
			"ageRatingId",
			"id").
		From(datastruct.CompositionTableName).
		Where(squirrel.Eq{"ageRatingId": ratingId})

	var compositions []datastruct.Composition
	var comp datastruct.Composition
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by country error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&comp.CompositionName,
			&comp.Description,
			&comp.AuthorId,
			&comp.GenreId,
			&comp.AgeRatingId,
			&comp.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by country error: %w", err)
		}
		compositions = append(compositions, comp)
	}

	return compositions, nil
}

func updateComposition(fromDB, new *datastruct.Composition) (updated datastruct.Composition) {
	updated = *fromDB
	if len(new.CompositionName) > 0 {
		updated.CompositionName = new.CompositionName
	}
	if len(new.Description) > 0 {
		updated.Description = new.Description
	}
	if new.AuthorId > 0 {
		updated.AuthorId = new.AuthorId
	}
	if new.GenreId > 0 {
		updated.GenreId = new.GenreId
	}
	if new.AgeRatingId > 0 {
		updated.AgeRatingId = new.AgeRatingId
	}
	return
}
