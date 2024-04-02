package service

import (
	"context"
	"robot-system-server/internal/model"
	"robot-system-server/internal/repository"
)

type BeasenService interface {
	GetBeasen(ctx context.Context, id int64) (*model.Beasen, error)

	RandResult(ctx context.Context) (*model.Beasen, error)
}

func NewBeasenService(service *Service, beasenRepository repository.BeasenRepository) BeasenService {
	return &beasenService{
		Service:          service,
		beasenRepository: beasenRepository,
	}
}

type beasenService struct {
	*Service
	beasenRepository repository.BeasenRepository
}

func (s *beasenService) GetBeasen(ctx context.Context, id int64) (*model.Beasen, error) {
	return s.beasenRepository.FirstById(ctx, id)
}

func (s *beasenService) RandResult(ctx context.Context) (*model.Beasen, error) {
	return s.beasenRepository.RandResult(ctx)
}
