package repository

import (
	"robot-system-server/internal/model"
)

type FortuneRepository interface {
	FindByQQ(qq string) (*model.QrFortune, error)
}

func NewFortuneRepository(repository *Repository) FortuneRepository {
	return &fortuneRepository{
		Repository: repository,
	}
}

type fortuneRepository struct {
	*Repository
}

func (r *fortuneRepository) FindByQQ(qq string) (*model.QrFortune, error) {
	var fortune model.QrFortune
	// TODO: query db
	return &fortune, nil
}
