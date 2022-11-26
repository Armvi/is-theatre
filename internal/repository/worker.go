package repository

import (
	"fmt"
	"is-theatre/internal/datastruct"
	"time"
	"github.com/Masterminds/squirrel"
)

type WorkerQuery interface {
	CreateWorker(worker datastruct.Worker) (*int64, error)
	Worker(id int64) (*datastruct.Worker, error)
	Workers(limit, offset uint64) ([]datastruct.Worker, error)
	DeleteWorker(id int64) error
	UpdateWorker(worker *datastruct.Worker) (*datastruct.Worker, error)
	WorkersByBirthDate(beg, end time.Time) ([]datastruct.Worker, error)
	WorkersBySalary(min, max float64) ([]datastruct.Worker, error)
}

type workerQuery struct{}

func (w *workerQuery) CreateWorker(worker datastruct.Worker) (*int64, error) {
	db := dbQueryBuilder().
		Insert(datastruct.WorkerTableName).
		Columns("name",
			"secondName",
			"birthDate",
			"salary").
		Values(worker.Name,
			worker.SecondName,
			worker.BirthDate,
			worker.Salary).
		Suffix("RETURNING id")

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create worker error: %w", err)
	}

	return &id, nil
}

func (w *workerQuery) Worker(id int64) (*datastruct.Worker, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"birthDate",
			"salary",
			"id").
		From(datastruct.WorkerTableName).
		Where(squirrel.Eq{"id": id})

	worker := datastruct.Worker{}
	err := db.QueryRow().
		Scan(&worker.Name,
			&worker.SecondName,
			&worker.BirthDate,
			&worker.Salary,
			&worker.Id)
	if err != nil {
		return nil, fmt.Errorf("get worker: %w", err)
	}

	return &worker, nil
}

func (w *workerQuery) Workers(limit, offset uint64) ([]datastruct.Worker, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"birthDate",
			"salary",
			"id").
		From(datastruct.WorkerTableName).
		Limit(limit).Offset(offset)

	var workers []datastruct.Worker
	var worker datastruct.Worker
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get workers error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&worker.Name,
			&worker.SecondName,
			&worker.BirthDate,
			&worker.Salary,
			&worker.Id)
		if err != nil {
			return nil, fmt.Errorf("get workers error: %w", err)
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func (w *workerQuery) DeleteWorker(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.WorkerTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete worker error: %w", err)
	}
	return nil
}

func (w *workerQuery) UpdateWorker(worker *datastruct.Worker) (*datastruct.Worker, error) {
	fromDB, err := w.Worker(worker.Id)
	if err != nil {
		return nil, fmt.Errorf("update worker error: %w", err)
	}

	updated := updateWorker(fromDB, worker)

	db := dbQueryBuilder().
		Update(datastruct.WorkerTableName).
		SetMap(map[string]interface{}{
			"name":       updated.Name,
			"secondName": updated.SecondName,
			"birthDate":  updated.BirthDate,
			"salary":     updated.Salary,
		}).Where(squirrel.Eq{"id": worker.Id}).
		Suffix("RETURNING name, secondName, brithDate, salary, id")

	uWorker := datastruct.Worker{}
	err = db.QueryRow().Scan(&uWorker.Name,
		&uWorker.SecondName,
		&uWorker.BirthDate,
		&uWorker.Salary,
		&uWorker.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("update worker error: %w", err)
	}

	return &uWorker, nil
}

func (w *workerQuery) WorkersByBirthDate(beg, end time.Time) ([]datastruct.Worker, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"birthDate",
			"salary",
			"id").
		From(datastruct.WorkerTableName).
		Where(squirrel.And{
			squirrel.Lt{"birthDate": end},
			squirrel.GtOrEq{"birthDate": beg},
		})

	var workers []datastruct.Worker
	var worker datastruct.Worker
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get workers by birth date error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&worker.Name,
			&worker.SecondName,
			&worker.BirthDate,
			&worker.Salary,
			&worker.Id)
		if err != nil {
			return nil, fmt.Errorf("get workers by birth date error: %w", err)
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func (w *workerQuery) WorkersBySalary(min, max float64) ([]datastruct.Worker, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"birthDate",
			"salary",
			"id").
		From(datastruct.WorkerTableName).
		Where(squirrel.And{
			squirrel.Lt{"salary": max},
			squirrel.GtOrEq{"salary": min},
		})

	var workers []datastruct.Worker
	var worker datastruct.Worker
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get workers by salary error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&worker.Name,
			&worker.SecondName,
			&worker.BirthDate,
			&worker.Salary,
			&worker.Id)
		if err != nil {
			return nil, fmt.Errorf("get workers by salary error: %w", err)
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func updateWorker(fromDB, new *datastruct.Worker) (updated datastruct.Worker) {
	updated = *fromDB
	if len(new.Name) > 0 {
		updated.Name = new.Name
	}
	if len(new.SecondName) > 0 {
		updated.SecondName = new.SecondName
	}
	t := time.Time{}
	if new.BirthDate.After(t) {
		updated.BirthDate = new.BirthDate
	}
	if new.Salary != 0 {
		updated.Salary = new.Salary
	}
	return
}