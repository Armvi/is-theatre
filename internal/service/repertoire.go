package service

import (
	"is-theatre/internal/repository"
)

type RepertoireService interface {
}

type repertoireService struct {
	dao repository.DAO
}

func NewRepertoireService(dao repository.DAO) RepertoireService {
	return &repertoireService{
		dao: dao,
	}
}
