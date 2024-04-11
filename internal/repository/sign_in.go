package repository

import (
	"context"
	"robot-system-server/internal/model"
)

type SignInRepository interface {
	FirstById(ctx context.Context, id int64) (*model.QrSignInDatum, error)
}

func NewSignInRepository(repository *Repository) SignInRepository {
	return &signInRepository{
		Repository: repository,
	}
}

type signInRepository struct {
	*Repository
}

func (r *signInRepository) FirstById(ctx context.Context, id int64) (*model.QrSignInDatum, error) {
	var signIn model.QrSignInDatum
	// TODO: query db
	return &signIn, nil
}
