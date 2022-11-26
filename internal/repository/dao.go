package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type DAO interface {
	NewUserQuery() UserQuery
	NewAuthorQuery() AuthorQuery
	NewAgeRatingQuery() AgeRatingQuery
	NewGenreQuery() GenreQuery
	NewCompositionQuery() CompositionQuery
	NewPersonageDescriptionQuery() PersonageDescriptionQuery
	NewPersonageQuery() PersonageQuery
	NewWorkerQuery() WorkerQuery
	NewDirectorQuery() DirectorQuery
	NewActorDescriptionQuery() ActorDescriptionQuery
	NewActorQuery() ActorQuery
	NewActorsRoleQuery() ActorsRoleQuery
	NewPerformanceQuery() PerformanceQuery
	NewRepertoireQuery() RepertoireQuery
	NewTicketQuery() TicketQuery
}

type dao struct{}

var DB *sql.DB

func dbQueryBuilder() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}

// QueryBuilder only for test
func QueryBuilder() squirrel.StatementBuilderType {
	return dbQueryBuilder()
}

func NewDAO(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (d *dao) NewAuthorQuery() AuthorQuery {
	return &authorQuery{}
}

func (d *dao) NewAgeRatingQuery() AgeRatingQuery {
	return &ageRatingQuery{}
}

func (d *dao) NewGenreQuery() GenreQuery {
	return &genreQuery{}
}

func (d *dao) NewCompositionQuery() CompositionQuery {
	return &compositionQuery{}
}

func (d *dao) NewPersonageDescriptionQuery() PersonageDescriptionQuery {
	return &personageDescriptionQuery{}
}

func (d *dao) NewPersonageQuery() PersonageQuery {
	return &personageQuery{}
}

func (d *dao) NewWorkerQuery() WorkerQuery {
	return &workerQuery{}
}

func (d *dao) NewDirectorQuery() DirectorQuery {
	return &directorQuery{}
}

func (d *dao) NewActorDescriptionQuery() ActorDescriptionQuery {
	return &actorDescriptionQuery{}
}

func (d *dao) NewActorQuery() ActorQuery {
	return &actorQuery{}
}

func (d *dao) NewActorsRoleQuery() ActorsRoleQuery {
	return &actorsRoleQuery{}
}

func (d *dao) NewPerformanceQuery() PerformanceQuery {
	return &performanceQuery{}
}

func (d *dao) NewRepertoireQuery() RepertoireQuery {
	return &repertoireQuery{}
}

func (d *dao) NewTicketQuery() TicketQuery {
	return &ticketQuery{}
}
