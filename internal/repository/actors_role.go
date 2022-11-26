package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type ActorsRoleQuery interface {
	CreateActorsRole(actorsRole datastruct.ActorsRole) (*int64, error)
	ActorsRole(id int64) (*datastruct.ActorsRole, error)
	UpdateActorsRole(actorsRole datastruct.ActorsRole) (*datastruct.ActorsRole, error)
	DeleteActorsRole(id int64) error
	ActorsRolesByPerformance(id uint64) ([]datastruct.ActorsRole, error)
}

type actorsRoleQuery struct{}

func (p *actorsRoleQuery) CreateActorsRole(actorsRole datastruct.ActorsRole) (*int64, error) {
	qb := dbQueryBuilder().
		Insert(datastruct.ActorsRoleTableName).
		Columns("performanceId",
			"actorId",
			"personageId").
		Values(actorsRole.PerformanceId,
			actorsRole.ActorId,
			actorsRole.PersonageId).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create actorsRole error: %w", err)
	}

	return &id, nil
}

func (p *actorsRoleQuery) ActorsRole(id int64) (*datastruct.ActorsRole, error) {
	db := dbQueryBuilder().
		Select("performanceId",
			"actorId",
			"personageId",
			"id").
		From(datastruct.ActorsRoleTableName).
		Where(squirrel.Eq{"id": id})

	actorsRole := datastruct.ActorsRole{}
	err := db.QueryRow().Scan(&actorsRole.PerformanceId,
		&actorsRole.ActorId,
		&actorsRole.PersonageId,
		&actorsRole.Id)
	if err != nil {
		return nil, fmt.Errorf("get actorsRole: %w", err)
	}

	return &actorsRole, nil
}

func (p *actorsRoleQuery) UpdateActorsRole(actorsRole datastruct.ActorsRole) (*datastruct.ActorsRole, error) {
	fromDB, err := p.ActorsRole(actorsRole.Id)
	if err != nil {
		return nil, fmt.Errorf("updated composition error: %w", err)
	}

	updated := updateActorsRole(fromDB, &actorsRole)

	db := dbQueryBuilder().
		Update(datastruct.UserTableName).
		SetMap(map[string]interface{}{
			"performanceId": updated.PerformanceId,
			"actorId":       updated.ActorId,
			"personageId":   updated.PersonageId,
		}).Where(squirrel.Eq{"id": actorsRole.Id}).
		Suffix("RETURNING actorsRoleName, descriptionId, compositionId, id")

	uActorsRole := datastruct.ActorsRole{}
	err = db.QueryRow().Scan(&uActorsRole.PerformanceId,
		&uActorsRole.ActorId,
		&uActorsRole.PersonageId,
		&uActorsRole.Id)
	if err != nil {
		return nil, fmt.Errorf("update actorsRole error: %w", err)
	}

	return &uActorsRole, nil
}

func (p *actorsRoleQuery) DeleteActorsRole(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.ActorsRoleTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (p *actorsRoleQuery) ActorsRolesByPerformance(id uint64) ([]datastruct.ActorsRole, error) {
	db := dbQueryBuilder().
		Select("actorsRoleName", "descriptionId", "compositionId", "id").
		From(datastruct.ActorsRoleTableName).
		Where(squirrel.Eq{"performanceId": id})

	var actorsRoles []datastruct.ActorsRole
	var actorsRole datastruct.ActorsRole
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get actors roles error: %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&actorsRole.PerformanceId,
			&actorsRole.ActorId,
			&actorsRole.PersonageId,
			&actorsRole.Id,
			&actorsRole.Id)
		if err != nil {
			return nil, fmt.Errorf("get actors roles error: %w", err)
		}
		actorsRoles = append(actorsRoles, actorsRole)
	}

	return actorsRoles, nil
}

func updateActorsRole(fromDB, new *datastruct.ActorsRole) (updated datastruct.ActorsRole) {
	updated = *fromDB
	if new.PerformanceId != 0 {
		updated.PerformanceId = new.PerformanceId
	}
	if new.ActorId != 0 {
		updated.ActorId = new.ActorId
	}
	if new.PersonageId != 0 {
		updated.PersonageId = new.PersonageId
	}
	return
}
