package repository

import (
	"context"
	"robot-system-server/internal/model"
)

type UserAssetsRepository interface {
	FirstById(ctx context.Context, id int64) (*model.QrUserAsset, error)
}

func NewUserAssetsRepository(repository *Repository) UserAssetsRepository {
	return &userAssetsRepository{
		Repository: repository,
	}
}

type userAssetsRepository struct {
	*Repository
}

func (r *userAssetsRepository) FirstById(ctx context.Context, id int64) (*model.QrUserAsset, error) {
	var userAssets model.QrUserAsset
	// TODO: query db
	return &userAssets, nil
}
