package repository

import (
	"context"
	"robot-system-server/internal/model"
	"time"
)

type FortuneDataRepository interface {
	FindDailyByQQ(ctx context.Context, qq string) (*model.QrFortuneDatum, error)
}

func NewFortuneDataRepository(repository *Repository) FortuneDataRepository {
	return &fortuneDataRepository{
		Repository: repository,
	}
}

type fortuneDataRepository struct {
	*Repository
}

func (r *fortuneDataRepository) FindDailyByQQ(ctx context.Context, qq string) (*model.QrFortuneDatum, error) {
	var fortuneData model.QrFortuneDatum
	err := r.DB(ctx).Where("qq = ? and DATE(FORTUNE_DATE) = ? ", qq, time.Now().Format("2006-01-02")).First(&fortuneData).Error
	if err != nil {
		return nil, err
	}
	return &fortuneData, nil
}
