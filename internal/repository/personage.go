package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type PersonageQuery interface {
	CreatePersonage(personage datastruct.Personage) (*int64, error)
	Personage(id int64) (*datastruct.Personage, error)
	Personages(limit, offset uint64) ([]datastruct.Personage, error)
	UpdatePersonage(personage datastruct.Personage) (*datastruct.Personage, error)
	DeletePersonage(id int64) error
	PersonagesByComposition(id int64) ([]datastruct.Personage, error)
}

type personageQuery struct{}

func (p *personageQuery) CreatePersonage(personage datastruct.Personage) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.PersonageTableName).
		Columns("personageName", "descriptionId", "compositionId").
		Values(personage.Name, personage.DescriptionId, personage.CompositionId).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create personage error: %w", err)
	}

	return &id, nil
}

func (p *personageQuery) Personage(id int64) (*datastruct.Personage, error) {
	db := dbQueryBuilder().
		Select("personageName", "descriptionId", "compositionId", "id").
		From(datastruct.PersonageTableName).
		Where(squirrel.Eq{"id": id})

	personage := datastruct.Personage{}
	err := db.QueryRow().Scan(&personage.Name, &personage.DescriptionId, &personage.CompositionId, &personage.Id)
	if err != nil {
		return nil, fmt.Errorf("get personage: %w", err)
	}

	return &personage, nil
}

func (p *personageQuery) Personages(limit, offset uint64) ([]datastruct.Personage, error) {
	db := dbQueryBuilder().
		Select("personageName", "descriptionId", "compositionId", "id").
		From(datastruct.PersonageTableName).
		Limit(limit).Offset(offset)

	var personages []datastruct.Personage
	var personage datastruct.Personage
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get personages error: %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&personage.Name, &personage.DescriptionId, &personage.CompositionId, &personage.Id)
		if err != nil {
			return nil, fmt.Errorf("get persoanges error: %w", err)
		}
		personages = append(personages, personage)
	}

	return personages, nil
}

func (p *personageQuery) UpdatePersonage(personage datastruct.Personage) (*datastruct.Personage, error) {
	fromDB, err := p.Personage(personage.Id)
	if err != nil {
		return nil, fmt.Errorf("updated composition error: %w", err)
	}

	updated := updatePersonage(fromDB, &personage)

	db := dbQueryBuilder().
		Update(datastruct.UserTableName).
		SetMap(map[string]interface{}{
			"personageName": updated.Name,
			"descriptionId": updated.DescriptionId,
			"compositionId": updated.CompositionId,
		}).Where(squirrel.Eq{"id": personage.Id}).
		Suffix("RETURNING personageName, descriptionId, compositionId, id")

	uPersonage := datastruct.Personage{}
	err = db.QueryRow().Scan(
		&uPersonage.Name,
		&uPersonage.DescriptionId,
		&uPersonage.CompositionId,
		&uPersonage.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("update personage error: %w", err)
	}

	return &uPersonage, nil
}

func (p *personageQuery) DeletePersonage(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.PersonageTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (p *personageQuery) PersonagesByComposition(compositionId int64) ([]datastruct.Personage, error) {
	db := dbQueryBuilder().
		Select("personageName", "descriptionId", "compositionId", "id").
		From(datastruct.PersonageTableName).
		Where(squirrel.Eq{"compositionId": compositionId})

	var personages []datastruct.Personage
	var personage datastruct.Personage
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get personages by composition error: %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&personage.Name,
			&personage.DescriptionId,
			&personage.CompositionId,
			&personage.Id)
		if err != nil {
			return nil, fmt.Errorf("get persoanges by composition error: %w", err)
		}
		personages = append(personages, personage)
	}

	return personages, nil
}

func updatePersonage(fromDB, new *datastruct.Personage) (updated datastruct.Personage) {
	updated = *fromDB
	if len(new.Name) > 0 {
		updated.Name = new.Name
	}
	if new.CompositionId != 0 {
		updated.CompositionId = new.CompositionId
	}
	if new.DescriptionId != 0 {
		updated.DescriptionId = new.DescriptionId
	}
	return
}
