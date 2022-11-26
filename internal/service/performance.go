package service

import (
	"is-theatre/internal/repository"
)

type PerformanceService interface {
}

type performanceService struct {
	dao repository.DAO
}

func NewPerformanceService(dao repository.DAO) PerformanceService {
	return &performanceService{
		dao: dao,
	}
}
