package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type DirectorQuery interface {
	CreateDirector(director datastruct.Director) (*int64, error)
	Director(id int64) (*datastruct.Director, error)
	Directors(limit, offset uint64) ([]datastruct.Director, error)
	UpdateDirector(director *datastruct.Director) (*datastruct.Director, error)
	DeleteDirector(id int64) error
}

type directorQuery struct{}

func (d directorQuery) CreateDirector(director datastruct.Director) (*int64, error) {
	db := dbQueryBuilder().
		Insert(datastruct.DirectorTableName).
		Columns("workerId").
		Values(director.WorkerId).
		Suffix("RETURNING id")

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create director error: %w", err)
	}

	return &id, nil
}

func (d directorQuery) Director(id int64) (*datastruct.Director, error) {
	db := dbQueryBuilder().
		Select("workerId",
			"id").
		From(datastruct.DirectorTableName).
		Where(squirrel.Eq{"id": id})

	dir := datastruct.Director{}
	err := db.QueryRow().
		Scan(&dir.WorkerId,
			&dir.Id)
	if err != nil {
		return nil, fmt.Errorf("get director error: %w", err)
	}

	return &dir, nil
}

func (d directorQuery) Directors(limit, offset uint64) ([]datastruct.Director, error) {
	db := dbQueryBuilder().
		Select("workerId",
			"id").
		From(datastruct.DirectorTableName).
		Limit(limit).Offset(offset)

	var directors []datastruct.Director
	var director datastruct.Director
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get directors error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&director.WorkerId,
			&director.Id)
		if err != nil {
			return nil, fmt.Errorf("get directors error: %w", err)
		}
		directors = append(directors, director)
	}

	return directors, nil
}

func (d directorQuery) UpdateDirector(director *datastruct.Director) (*datastruct.Director, error) {
	fromDB, err := d.Director(director.Id)
	if err != nil {
		return nil, fmt.Errorf("update director error: %w", err)
	}

	updated := updateDirector(fromDB, director)

	db := dbQueryBuilder().
		Update(datastruct.WorkerTableName).
		SetMap(map[string]interface{}{
			"workerId": updated.WorkerId,
		}).Where(squirrel.Eq{"id": director.Id}).
		Suffix("RETURNING name, secondName, brithDate, salary, id")

	uDirector := datastruct.Director{}
	err = db.QueryRow().Scan(&uDirector.WorkerId,
		&uDirector.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("update worker error: %w", err)
	}

	return &uDirector, nil
}

func (d directorQuery) DeleteDirector(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.DirectorTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete director error: %w", err)
	}
	return nil
}

func updateDirector(fromDB, new *datastruct.Director) (updated datastruct.Director) {
	updated = *fromDB
	if new.WorkerId != 0 {
		updated.WorkerId = new.WorkerId
	}
	return
}
