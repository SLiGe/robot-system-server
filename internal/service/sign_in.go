package service

import (
	"errors"
	"github.com/duke-git/lancet/v2/random"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/internal/repository"
	"robot-system-server/pkg/db"
	"strconv"
	"time"
)

type SignInService interface {
	DoSignIn(req v1.SignInForQqRequest) (v1.SignInDataResponse, error)
	QuerySignInData(req v1.QuerySignInDataRequest) (v1.SignInDetailResponse, error)
}

func NewSignInService(service *Service, signInRepository repository.SignInRepository, userAssetsService UserAssetsService) SignInService {
	return &signInService{
		Service:           service,
		signInRepository:  signInRepository,
		userAssetsService: userAssetsService,
	}
}

type signInService struct {
	*Service
	signInRepository  repository.SignInRepository
	userAssetsService UserAssetsService
}

func (s *signInService) QuerySignInData(req v1.QuerySignInDataRequest) (v1.SignInDetailResponse, error) {
	l := query.QrSignInLevel
	d := query.QrSignInDatum
	y := query.QrSignInDay
	var data model.QrSignInDatum
	_ = d.Select(d.Points.As("POINTS"), d.DayID.As("DAY_ID")).Where(d.Qq.Eq(req.QQ)).Scan(&data)
	var level string
	_ = l.Select(l.Level).Where(l.MaxPoints.Gte(data.Points), l.MinPoints.Lte(data.Points)).Scan(&level)
	var msgOfDay string
	_ = query.QrMsgOfDay.Select(query.QrMsgOfDay.Sentence).Order(field.Func.Rand()).Limit(1).Scan(&msgOfDay)
	var day model.QrSignInDay
	_ = y.Select(y.MonthDays.As("MONTH_DAYS"), y.TotalDays.As("TOTAL_DAYS")).Where(y.ID.Eq(data.DayID)).Scan(&day)
	return v1.SignInDetailResponse{
		CurrentLevel: level,
		TodayMsg:     msgOfDay,
		Qq:           req.QQ,
		Points:       data.Points,
		MonthDay:     day.MonthDays,
		TotalDay:     day.TotalDays,
	}, nil

}

func (s *signInService) DoSignIn(req v1.SignInForQqRequest) (v1.SignInDataResponse, error) {
	if req.Points == 0 || req.Points > 60 {
		req.Points = int64(random.RandInt(20, 60))
	}
	q := query.QrSignInDatum
	d := query.QrSignInDay
	signInDatum, err := q.Select(q.ALL).Where(q.Qq.Eq(req.QQ)).First()
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			data, day := s.SaveSignInData(req.QQ, req.Points)
			return v1.SignInDataResponse{
				Status:  200,
				Message: "签到成功",
				Data:    s.FillDetailResponse(data, day),
			}, nil
		}
		return v1.SignInDataResponse{}, err
	}
	signInDay, err := d.Where(d.ID.Eq(signInDatum.ID)).Attrs(field.Attrs(&model.QrSignInDay{
		Qq:        req.QQ,
		MonthDays: 1,
		TotalDays: 1,
	})).FirstOrCreate()
	if signInDatum.SignInDate.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return v1.SignInDataResponse{
			Status:  203,
			Message: "已经签到",
			Data:    s.FillDetailResponse(*signInDatum, *signInDay),
		}, nil
	} else {
		s.UpdateSignInDataAndDays(req.Points, *signInDatum, *signInDay)
		return v1.SignInDataResponse{
			Status:  200,
			Message: "签到成功",
			Data:    s.FillDetailResponse(*signInDatum, *signInDay),
		}, nil
	}
}

func (s *signInService) SaveSignInData(qq string, points int64) (model.QrSignInDatum, model.QrSignInDay) {
	q := query.QrSignInDatum
	d := query.QrSignInDay
	day := model.QrSignInDay{
		Qq:        qq,
		MonthDays: 1,
		TotalDays: 1,
	}
	_ = d.Create(&day)
	qqNumber, _ := strconv.ParseInt(qq, 10, 64)
	creator := "admin"
	data := model.QrSignInDatum{
		Qq:         qq,
		QqNumber:   &qqNumber,
		Points:     points,
		DayID:      day.ID,
		SignInDate: db.Now(),
		CreateBy:   &creator,
		UpdateBy:   &creator,
	}
	_ = q.Create(&data)
	return data, day
}

func (s *signInService) UpdateSignInDataAndDays(addPoints int64, datum model.QrSignInDatum, day model.QrSignInDay) {
	q := query.QrSignInDatum
	d := query.QrSignInDay
	_, _ = q.Where(q.ID.Eq(datum.ID)).Updates(map[string]interface{}{"Points": datum.Points + addPoints, "Sign_In_Date": time.Now()})
	_, _ = d.Where(d.ID.Eq(day.ID)).Updates(map[string]interface{}{"Month_Days": day.MonthDays + 1, "Total_Days": day.TotalDays + 1})
	s.userAssetsService.AddUserPoints(datum.Qq, addPoints)
}

func (s *signInService) FillDetailResponse(signInDatum model.QrSignInDatum, signInDay model.QrSignInDay) v1.SignInDetailResponse {
	l := query.QrSignInLevel
	var level string
	_ = l.Select(l.Level).Where(l.MaxPoints.Gte(signInDatum.Points), l.MinPoints.Lte(signInDatum.Points)).Scan(&level)
	var msgOfDay string
	_ = query.QrMsgOfDay.Select(query.QrMsgOfDay.Sentence).Order(field.Func.Rand()).Limit(1).Scan(&msgOfDay)
	return v1.SignInDetailResponse{
		CurrentLevel: level,
		TodayMsg:     msgOfDay,
		Qq:           signInDatum.Qq,
		Points:       signInDatum.Points,
		MonthDay:     signInDay.MonthDays,
		TotalDay:     signInDay.TotalDays,
	}
}