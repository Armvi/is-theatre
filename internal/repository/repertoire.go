package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
	"time"
)

type RepertoireQuery interface {
	CreateRepertoire(repertoire datastruct.Repertoire) (*int64, error)
	Repertoire(id int64) (*datastruct.Repertoire, error)
	UpdateRepertoire(repertoire datastruct.Repertoire) (*datastruct.Repertoire, error)
	DeleteRepertoire(id int64) error
	RepertoireByDate(date time.Time) (*datastruct.Repertoire, error)
}

type repertoireQuery struct{}

func (p *repertoireQuery) CreateRepertoire(repertoire datastruct.Repertoire) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.RepertoireTableName).
		Columns("periodBegin",
			"periodEnd",
			"description").
		Values(repertoire.BeginDate,
			repertoire.EndDate,
			repertoire.Description).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create repertoire error: %w", err)
	}

	return &id, nil
}

func (p *repertoireQuery) Repertoire(id int64) (*datastruct.Repertoire, error) {
	db := dbQueryBuilder().
		Select("periodBegin",
			"periodEnd",
			"description",
			"id").
		From(datastruct.RepertoireTableName).
		Where(squirrel.Eq{"id": id})

	repertoire := datastruct.Repertoire{}
	err := db.QueryRow().Scan(&repertoire.BeginDate,
		&repertoire.EndDate,
		&repertoire.Description,
		&repertoire.Id)
	if err != nil {
		return nil, fmt.Errorf("get repertoire: %w", err)
	}

	return &repertoire, nil
}

func (p *repertoireQuery) UpdateRepertoire(repertoire datastruct.Repertoire) (*datastruct.Repertoire, error) {
	err := p.DeleteRepertoire(repertoire.Id)
	if err != nil {
		return nil, fmt.Errorf("update repertoire err: %w", err)
	}

	id, err := p.CreateRepertoire(repertoire)

	if err != nil {
		return nil, fmt.Errorf("update repertoire err: %w", err)
	}

	return p.Repertoire(*id)
}

func (p *repertoireQuery) DeleteRepertoire(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.RepertoireTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (p *repertoireQuery) RepertoireByDate(date time.Time) (*datastruct.Repertoire, error) {
	db := dbQueryBuilder().
		Select("periodBegin",
			"periodEnd",
			"description",
			"id").
		From(datastruct.RepertoireTableName).
		Where(squirrel.And{
			squirrel.Lt{"periodEnd": date},
			squirrel.GtOrEq{"periodBegin": date},
		})

	repertoire := datastruct.Repertoire{}
	err := db.QueryRow().Scan(&repertoire.BeginDate,
		&repertoire.EndDate,
		&repertoire.Description,
		&repertoire.Id)
	if err != nil {
		return nil, fmt.Errorf("get repertoire: %w", err)
	}

	return &repertoire, nil
}
