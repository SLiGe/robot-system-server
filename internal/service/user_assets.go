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
	qrUser := s.userService.QueryOrCreate(userAccount)
	initPoints := float64(0)
	qrUserAsset, _ := u.Select(u.ALL).Where(u.UserAccount.Eq(userAccount)).Attrs(field.Attrs(&model.QrUserAsset{
		UserID:      &qrUser.UserID,
		UserAccount: &userAccount,
		Points:      &initPoints,
	})).FirstOrCreate()
	if qrUserAsset != nil {
		_, _ = u.Where(u.ID.Eq(qrUserAsset.ID)).Update(u.Points, u.Points.Add(float64(points)))
	}
}
