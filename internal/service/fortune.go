package service

import (
	"errors"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/logic"
	"robot-system-server/internal/model"
	"robot-system-server/internal/repository"
	"strings"
)

type FortuneService interface {
	GetFortuneOfToday(ctx *gin.Context, qq string, isOne int, isGroup int, isIntegral int, groupNum string, oldApi bool) (*v1.GetFortuneResponse, error)
}

func NewFortuneService(service *Service, fortuneRepository repository.FortuneRepository, fortuneDataRepository repository.FortuneDataRepository, fortuneLogic logic.FortuneLogic) FortuneService {
	return &fortuneService{
		Service:               service,
		fortuneRepository:     fortuneRepository,
		fortuneDataRepository: fortuneDataRepository,
		fortuneLogic:          fortuneLogic,
	}
}

type fortuneService struct {
	*Service
	fortuneRepository     repository.FortuneRepository
	fortuneDataRepository repository.FortuneDataRepository
	fortuneLogic          logic.FortuneLogic
}

func (s *fortuneService) GetFortuneOfToday(ctx *gin.Context, qq string, isOne int, isGroup int, isIntegral int, groupNum string, oldApi bool) (*v1.GetFortuneResponse, error) {
	fortuneData, err := s.fortuneDataRepository.FindDailyByQQ(ctx, qq)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//更新群号
		if isOne == 1 && isGroup == 1 && !oldApi {
			if checkGroupAction(groupNum, *fortuneData) {
				return &v1.GetFortuneResponse{
					Status:  201,
					Message: "今日运势已进行",
				}, err
			}
			//更新群号
			s.fortuneLogic.UpdateGroupNum(groupNum, *fortuneData)
		}
		//扣除积分
		if isIntegral == 1 && isGroup == 1 {
			s.fortuneLogic.CostUserPoints(qq)
		}
		var fortune model.QrFortune
		_ = jsoniter.UnmarshalFromString(fortuneData.JSONData, &fortune)
		return &v1.GetFortuneResponse{
			Status:       200,
			Message:      "获取运势成功",
			DataResponse: &fortune,
		}, nil
	}
	qrFortune := s.fortuneLogic.GetFortune()
	s.fortuneLogic.SaveNewFortuneData(qq, groupNum, isOne, isGroup, qrFortune)
	if isIntegral == 1 && isGroup == 1 {
		s.fortuneLogic.CostUserPoints(qq)
	}
	return &v1.GetFortuneResponse{
		Status:       200,
		Message:      "获取运势成功",
		DataResponse: qrFortune,
	}, nil
}

func checkGroupAction(groupNum string, qrFortuneData model.QrFortuneDatum) bool {
	groupNums := qrFortuneData.GroupNum
	if strutil.IsNotBlank(groupNum) && strings.Contains(*groupNums, ",") {
		groupArray := strings.Split(*groupNums, ",")
		groupList := make([]string, len(groupArray))
		for i, v := range groupArray {
			groupList[i] = v
		}
		for _, v := range groupList {
			if v == groupNum {
				return true
			}
		}
	}
	return false
}
