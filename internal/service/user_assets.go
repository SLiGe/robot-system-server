package service

import (
	"gorm.io/gen/field"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/internal/repository"
)

type UserAssetsService interface {
	AddUserPoints(userAccount string, points int64) error
}

func NewUserAssetsService(service *Service, userAssetsRepository repository.UserAssetsRepository, userService UserService) UserAssetsService {
	return &userAssetsService{
		Service:              service,
		userAssetsRepository: userAssetsRepository,
		userService:          userService,
	}
}

type userAssetsService struct {
	*Service
	userAssetsRepository repository.UserAssetsRepository
	userService          UserService
}

func (s *userAssetsService) AddUserPoints(userAccount string, points int64) error {
	u := query.QrUserAsset
	qrUser, err := s.userService.QueryOrCreate(userAccount)
	if err != nil {
		return err
	}
	initPoints := float64(0)
	err = query.Q.Transaction(func(tx *query.Query) error {
		qrUserAsset, err := tx.QrUserAsset.Select(u.ALL).Where(u.UserAccount.Eq(userAccount)).Attrs(field.Attrs(&model.QrUserAsset{
			UserID:      &qrUser.UserID,
			UserAccount: &userAccount,
			Points:      &initPoints,
		})).FirstOrCreate()
		if err != nil {
			return err
		}
		if qrUserAsset != nil {
			_, err := tx.QrUserAsset.Where(u.ID.Eq(qrUserAsset.ID)).Update(u.Points, u.Points.Add(float64(points)))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
