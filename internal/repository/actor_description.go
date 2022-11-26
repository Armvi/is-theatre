package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type ActorDescriptionQuery interface {
	CreateActorDescription(description datastruct.ActorDescription) (*int64, error)
	ActorDescription(id int64) (*datastruct.ActorDescription, error)
	DeleteActorDescription(id int64) error
	UpdateActorDescription(description datastruct.ActorDescription) (*datastruct.ActorDescription, error)
}

type actorDescriptionQuery struct{}

func (p *actorDescriptionQuery) CreateActorDescription(description datastruct.ActorDescription) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.ActorDescriptionTableName).
		Columns("age",
			"voice",
			"height",
			"weight",
			"gender").
		Values(description.Age,
			description.Voice,
			description.Height,
			description.Weight,
			description.Gender).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create actor description error: %w", err)
	}

	return &id, nil
}

func (p *actorDescriptionQuery) ActorDescription(id int64) (*datastruct.ActorDescription, error) {
	db := dbQueryBuilder().
		Select("age",
			"voice",
			"height",
			"weight",
			"gender",
			"id").
		From(datastruct.ActorDescriptionTableName).
		Where(squirrel.Eq{"id": id})

	description := datastruct.ActorDescription{}
	err := db.QueryRow().Scan(&description.Age,
		&description.Voice,
		&description.Height,
		&description.Weight,
		&description.Gender,
		&description.Id)
	if err != nil {
		return nil, fmt.Errorf("get actor description error: %w", err)
	}

	return &description, nil
}

func (p *actorDescriptionQuery) DeleteActorDescription(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.ActorDescriptionTableName).
		From(datastruct.ActorDescriptionTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete actor description error: %w", err)
	}
	return nil
}

func (p *actorDescriptionQuery) UpdateActorDescription(description datastruct.ActorDescription) (*datastruct.ActorDescription, error) {
	fromDB, err := p.ActorDescription(description.Id)
	if err != nil {
		return nil, fmt.Errorf("updated actor description error: %w", err)
	}

	updated := updateActorDescription(fromDB, &description)

	qb := dbQueryBuilder().
		Update(datastruct.ActorDescriptionTableName).
		SetMap(map[string]interface{}{
			"age":    updated.Age,
			"voice":  updated.Voice,
			"height": updated.Height,
			"weight": updated.Weight,
			"gender": updated.Gender,
		}).Where(squirrel.Eq{"id": description.Id}).
		Suffix("RETURNING age, voice, height, weight, gender, id")

	uDescription := datastruct.ActorDescription{}
	err = qb.QueryRow().Scan(&description.Age,
		&uDescription.Voice,
		&uDescription.Height,
		&uDescription.Weight,
		&uDescription.Gender,
		&uDescription.Id)
	if err != nil {
		return nil, fmt.Errorf("update actor description error: %w", err)
	}

	return &uDescription, nil
}

func updateActorDescription(fromDB, new *datastruct.ActorDescription) (updated datastruct.ActorDescription) {
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
	return
}
