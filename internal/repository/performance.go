package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
	"time"
)

type PerformanceQuery interface {
	CreatePerformance(performance datastruct.Performance) (*int64, error)
	Performance(id int64) (*datastruct.Performance, error)
	Performances(limit, offset uint64) ([]datastruct.Performance, error)
	DeletePerformance(id int64) error
	UpdatePerformance(performance *datastruct.Performance) (*datastruct.Performance, error)
	PerformancesByDate(beg, end time.Time) ([]datastruct.Performance, error)
	PerformancesByDirector(id int64) ([]datastruct.Performance, error)
	PerformancesByComposition(id int64) ([]datastruct.Performance, error)
}

type performanceQuery struct{}

func (w *performanceQuery) CreatePerformance(performance datastruct.Performance) (*int64, error) {
	db := dbQueryBuilder().
		Insert(datastruct.PerformanceTableName).
		Columns("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime").
		Values(performance.CompositionId,
			performance.PerformanceName,
			performance.Description,
			performance.DirectorId,
			performance.Date,
			performance.Time).
		Suffix("RETURNING id")

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create performance error: %w", err)
	}

	return &id, nil
}

func (w *performanceQuery) Performance(id int64) (*datastruct.Performance, error) {
	db := dbQueryBuilder().
		Select("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime",
			"id").
		From(datastruct.PerformanceTableName).
		Where(squirrel.Eq{"id": id})

	performance := datastruct.Performance{}
	err := db.QueryRow().
		Scan(&performance.CompositionId,
			&performance.PerformanceName,
			&performance.Description,
			&performance.DirectorId,
			&performance.Date,
			&performance.Time,
			&performance.Id)
	if err != nil {
		return nil, fmt.Errorf("get performance: %w", err)
	}

	return &performance, nil
}

func (w *performanceQuery) Performances(limit, offset uint64) ([]datastruct.Performance, error) {
	db := dbQueryBuilder().
		Select("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime",
			"id").
		From(datastruct.PerformanceTableName).
		Limit(limit).Offset(offset)

	var performances []datastruct.Performance
	var performance datastruct.Performance
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get performances error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&performance.CompositionId,
			&performance.PerformanceName,
			&performance.Description,
			&performance.DirectorId,
			&performance.Date,
			&performance.Time,
			&performance.Id)
		if err != nil {
			return nil, fmt.Errorf("get performances error: %w", err)
		}
		performances = append(performances, performance)
	}

	return performances, nil
}

func (w *performanceQuery) DeletePerformance(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.PerformanceTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete performance error: %w", err)
	}
	return nil
}

func (w *performanceQuery) UpdatePerformance(performance *datastruct.Performance) (*datastruct.Performance, error) {
	fromDB, err := w.Performance(performance.Id)
	if err != nil {
		return nil, fmt.Errorf("update performance error: %w", err)
	}

	updated := updatePerformance(fromDB, performance)

	db := dbQueryBuilder().
		Update(datastruct.PerformanceTableName).
		SetMap(map[string]interface{}{
			"compositionId":   updated.CompositionId,
			"performanceName": updated.PerformanceName,
			"description":     updated.Description,
			"directorId":      updated.DirectorId,
			"performanceDate": updated.Date,
			"performanceTime": updated.Time,
		}).Where(squirrel.Eq{"id": performance.Id}).
		Suffix("RETURNING name, secondName, brithDate, salary, id")

	uPerformance := datastruct.Performance{}
	err = db.QueryRow().Scan(&uPerformance.CompositionId,
		&uPerformance.PerformanceName,
		&uPerformance.Description,
		&uPerformance.DirectorId,
		&uPerformance.Date,
		&uPerformance.Time,
		&uPerformance.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("update performance error: %w", err)
	}

	return &uPerformance, nil
}

func (w *performanceQuery) PerformancesByDate(beg, end time.Time) ([]datastruct.Performance, error) {
	db := dbQueryBuilder().
		Select("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime",
			"id").
		From(datastruct.PerformanceTableName).
		Where(squirrel.And{
			squirrel.Lt{"birthDate": end},
			squirrel.GtOrEq{"birthDate": beg},
		})

	var performances []datastruct.Performance
	var performance datastruct.Performance
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get performances by birth date error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&performance.CompositionId,
			&performance.PerformanceName,
			&performance.Description,
			&performance.DirectorId,
			&performance.Date,
			&performance.Time,
			&performance.Id)
		if err != nil {
			return nil, fmt.Errorf("get performances by birth date error: %w", err)
		}
		performances = append(performances, performance)
	}

	return performances, nil
}

func (w *performanceQuery) PerformancesByDirector(id int64) ([]datastruct.Performance, error) {
	db := dbQueryBuilder().
		Select("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime",
			"id").
		From(datastruct.PerformanceTableName).
		Where(squirrel.Eq{"directorId": id})

	var performances []datastruct.Performance
	var performance datastruct.Performance
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get performances by salary error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&performance.CompositionId,
			&performance.PerformanceName,
			&performance.Description,
			&performance.DirectorId,
			&performance.Date,
			&performance.Time,
			&performance.Id)
		if err != nil {
			return nil, fmt.Errorf("get performances by salary error: %w", err)
		}
		performances = append(performances, performance)
	}

	return performances, nil
}

func (w *performanceQuery) PerformancesByComposition(id int64) ([]datastruct.Performance, error) {
	db := dbQueryBuilder().
		Select("compositionId",
			"performanceName",
			"description",
			"directorId",
			"performanceDate",
			"performanceTime",
			"id").
		From(datastruct.PerformanceTableName).
		Where(squirrel.Eq{"compositionId": id})

	var performances []datastruct.Performance
	var performance datastruct.Performance
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get performances by salary error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&performance.CompositionId,
			&performance.PerformanceName,
			&performance.Description,
			&performance.DirectorId,
			&performance.Date,
			&performance.Time,
			&performance.Id)
		if err != nil {
			return nil, fmt.Errorf("get performances by salary error: %w", err)
		}
		performances = append(performances, performance)
	}

	return performances, nil
}

func updatePerformance(fromDB, new *datastruct.Performance) (updated datastruct.Performance) {
	updated = *fromDB
	if new.CompositionId != 0 {
		updated.CompositionId = new.CompositionId
	}
	if len(new.PerformanceName) > 0 {
		updated.PerformanceName = new.PerformanceName
	}
	if len(new.Description) > 0 {
		updated.Description = new.Description
	}
	if new.DirectorId != 0 {
		updated.DirectorId = new.DirectorId
	}
	t := time.Time{}
	if new.Date.After(t) {
		updated.Date = new.Date
	}
	if new.Time.After(t) {
		updated.Time = new.Time
	}
	return
}
