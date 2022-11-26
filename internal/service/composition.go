package service

import (
	"is-theatre/internal/entity"
	"is-theatre/internal/repository"
)

type CompositionService interface {
	CreateComposition(composition entity.Composition) (*int64, error)
	Composition(id int64) (*entity.Composition, error)
	DeleteComposition(id int64) error
	UpdateComposition(composition entity.Composition) (*entity.Composition, error)
	Compositions(limit, offset uint64) ([]entity.Composition, error)
	CompositionsByConditions(genre entity.Genre, rating entity.AgeRating, author entity.Author)
}

type compositionService struct {
	dao repository.DAO
}

func NewCompositionService(dao repository.DAO) AuthorService {
	return &authorService{dao: dao}
}
