package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type ActorQuery interface {
	CreateActor(actor datastruct.Actor) (*int64, error)
	Actor(id int64) (*datastruct.Actor, error)
	Actors(limit, offset uint64) ([]datastruct.Actor, error)
	DeleteActor(id int64) error
	UpdateActor(actor *datastruct.Actor) (*datastruct.Actor, error)
	ActorsByExperience(min, max float64) ([]datastruct.Actor, error)
}

type actorQuery struct{}

func (w *actorQuery) CreateActor(actor datastruct.Actor) (*int64, error) {
	db := dbQueryBuilder().
		Insert(datastruct.ActorTableName).
		Columns("workerId",
			"descriptionId",
			"experience").
		Values(actor.WorkerId,
			actor.DescriptionId,
			actor.Experience).
		Suffix("RETURNING id")

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create actor error: %w", err)
	}

	return &id, nil
}

func (w *actorQuery) Actor(id int64) (*datastruct.Actor, error) {
	db := dbQueryBuilder().
		Select("workerId",
			"descriptionId",
			"experience",
			"id").
		From(datastruct.ActorTableName).
		Where(squirrel.Eq{"id": id})

	actor := datastruct.Actor{}
	err := db.QueryRow().
		Scan(&actor.WorkerId,
			&actor.DescriptionId,
			&actor.Experience,
			&actor.Id)
	if err != nil {
		return nil, fmt.Errorf("get actor: %w", err)
	}

	return &actor, nil
}

func (w *actorQuery) Actors(limit, offset uint64) ([]datastruct.Actor, error) {
	db := dbQueryBuilder().
		Select("workerId",
			"descriptionId",
			"experience",
			"id").
		From(datastruct.ActorTableName).
		Limit(limit).Offset(offset)

	var actors []datastruct.Actor
	var actor datastruct.Actor
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get actors error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&actor.WorkerId,
			&actor.DescriptionId,
			&actor.Experience,
			&actor.Id)
		if err != nil {
			return nil, fmt.Errorf("get actors error: %w", err)
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

func (w *actorQuery) DeleteActor(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.ActorTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete actor error: %w", err)
	}
	return nil
}

func (w *actorQuery) UpdateActor(actor *datastruct.Actor) (*datastruct.Actor, error) {
	fromDB, err := w.Actor(actor.Id)
	if err != nil {
		return nil, fmt.Errorf("update actor error: %w", err)
	}

	updated := updateActor(fromDB, actor)

	db := dbQueryBuilder().
		Update(datastruct.ActorTableName).
		SetMap(map[string]interface{}{
			"name":       updated.WorkerId,
			"secondName": updated.DescriptionId,
			"birthDate":  updated.Experience,
		}).Where(squirrel.Eq{"id": actor.Id}).
		Suffix("RETURNING name, secondName, brithDate, salary, id")

	uActor := datastruct.Actor{}
	err = db.QueryRow().Scan(&actor.WorkerId,
		&actor.DescriptionId,
		&actor.Experience,
		&actor.Id)
	if err != nil {
		return nil, fmt.Errorf("update actor error: %w", err)
	}

	return &uActor, nil
}

func (w *actorQuery) ActorsByExperience(min, max float64) ([]datastruct.Actor, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"birthDate",
			"salary",
			"id").
		From(datastruct.ActorTableName).
		Where(squirrel.And{
			squirrel.Lt{"experience": max},
			squirrel.GtOrEq{"experience": min},
		})

	var actors []datastruct.Actor
	var actor datastruct.Actor
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get actors by salary error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&actor.WorkerId,
			&actor.DescriptionId,
			&actor.Experience,
			&actor.Id)
		if err != nil {
			return nil, fmt.Errorf("get actors by salary error: %w", err)
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

func updateActor(fromDB, new *datastruct.Actor) (updated datastruct.Actor) {
	updated = *fromDB
	if new.WorkerId != 0 {
		updated.WorkerId = new.WorkerId
	}
	if new.DescriptionId != 0 {
		updated.DescriptionId = new.DescriptionId
	}
	if new.Experience != 0 {
		updated.Experience = new.Experience
	}
	return
}
