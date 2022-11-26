package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type PersonageDescriptionQuery interface {
	CreatePersonageDescription(description datastruct.PersonageDescription) (*int64, error)
	PersonageDescription(id int64) (*datastruct.PersonageDescription, error)
	DeletePersonageDescription(id int64) error
	UpdatePersonageDescription(description datastruct.PersonageDescription) (*datastruct.PersonageDescription, error)
}

type personageDescriptionQuery struct{}

func (p *personageDescriptionQuery) CreatePersonageDescription(description datastruct.PersonageDescription) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.PersonageDescriptionTableName).
		Columns("age",
			"voice",
			"height",
			"weight",
			"gender",
			"description").
		Values(description.Age,
			description.Voice,
			description.Height,
			description.Weight,
			description.Gender,
			description.Description).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create personage description error: %w", err)
	}

	return &id, nil
}

func (p *personageDescriptionQuery) PersonageDescription(id int64) (*datastruct.PersonageDescription, error) {
	db := dbQueryBuilder().
		Select("age",
			"voice",
			"height",
			"weight",
			"gender",
			"description",
			"id").
		From(datastruct.PersonageDescriptionTableName).
		Where(squirrel.Eq{"id": id})

	description := datastruct.PersonageDescription{}
	err := db.QueryRow().Scan(&description.Age,
		&description.Voice,
		&description.Height,
		&description.Weight,
		&description.Gender,
		&description.Description,
		&description.Id)
	if err != nil {
		return nil, fmt.Errorf("get description error: %w", err)
	}

	return &description, nil
}

func (p *personageDescriptionQuery) DeletePersonageDescription(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.PersonageDescriptionTableName).
		From(datastruct.PersonageDescriptionTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete description error: %w", err)
	}
	return nil
}

func (p *personageDescriptionQuery) UpdatePersonageDescription(description datastruct.PersonageDescription) (*datastruct.PersonageDescription, error) {
	fromDB, err := p.PersonageDescription(description.Id)
	if err != nil {
		return nil, fmt.Errorf("updated description error: %w", err)
	}

	updated := updatePersonageDescription(fromDB, &description)

	qb := dbQueryBuilder().
		Update(datastruct.PersonageDescriptionTableName).
		SetMap(map[string]interface{}{
			"age":         updated.Age,
			"voice":       updated.Voice,
			"height":      updated.Height,
			"weight":      updated.Weight,
			"gender":      updated.Gender,
			"description": updated.Description,
		}).Where(squirrel.Eq{"id": description.Id}).
		Suffix("RETURNING age, voice, height, weight, gender, description, id")

	uDescription := datastruct.PersonageDescription{}
	err = qb.QueryRow().Scan(&description.Age,
		&uDescription.Voice,
		&uDescription.Height,
		&uDescription.Weight,
		&uDescription.Gender,
		&uDescription.Description,
		&uDescription.Id)
	if err != nil {
		return nil, fmt.Errorf("update description error: %w", err)
	}

	return &uDescription, nil
}

func updatePersonageDescription(fromDB, new *datastruct.PersonageDescription) (updated datastruct.PersonageDescription) {
	updated = *fromDB
	if len(new.Age) > 0 {
		updated.Age = new.Age
	}
	if len(new.Voice) > 0 {
		updated.Voice = new.Voice
	}
	if len(new.Height) > 0 {
		updated.Height = new.Height
	}
	if len(new.Weight) > 0 {
		updated.Weight = new.Weight
	}
	if !new.Gender {
		updated.Gender = new.Gender
	}
	if len(new.Description) > 0 {
		updated.Description = new.Description
	}
	return
}
