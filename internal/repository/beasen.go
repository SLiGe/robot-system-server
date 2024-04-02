package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"robot-system-server/internal/model"
)

type BeasenRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Beasen, error)

	RandResult(ctx context.Context) (*model.Beasen, error)
}

func NewBeasenRepository(repository *Repository) BeasenRepository {
	return &beasenRepository{
		Repository: repository,
	}
}

type beasenRepository struct {
	*Repository
}

func (r *beasenRepository) RandResult(ctx context.Context) (*model.Beasen, error) {
	var bean model.Beasen
	if err := r.DB(ctx).Where("del_flag = ?", "0").Order("rand()").First(&bean).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &bean, nil
}

func (r *beasenRepository) FirstById(ctx context.Context, id int64) (*model.Beasen, error) {
	var beasen model.Beasen
	if err := r.DB(ctx).Where("id = ?", id).First(&beasen).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &beasen, nil
}
