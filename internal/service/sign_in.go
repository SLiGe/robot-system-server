package service

import (
	"errors"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
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
	DoSignIn(ctx *gin.Context, req v1.SignInForQqRequest) (v1.SignInDataResponse, error)
	QuerySignInData(req v1.QuerySignInDataRequest) (v1.SignInDetailResponse, error)
	AddSignInPoints(req v1.AddSignPointsRequest) (v1.SignInDataResponse, error)
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

func (s *signInService) DoSignIn(ctx *gin.Context, req v1.SignInForQqRequest) (v1.SignInDataResponse, error) {
	if req.Points == 0 || req.Points > 60 {
		req.Points = int32(random.RandInt(20, 60))
	}
	q := query.QrSignInDatum
	d := query.QrSignInDay
	signInDatum, err := q.Select(q.ALL).Where(q.Qq.Eq(req.QQ)).First()
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			data, day, saveErr := s.SaveSignInData(req.QQ, req.Points)
			if saveErr != nil {
				return v1.SignInDataResponse{}, saveErr
			}
			return v1.SignInDataResponse{
				Data: s.FillDetailResponse(data, day),
			}, nil
		}
		return v1.SignInDataResponse{}, err
	}
	var signInDay *model.QrSignInDay
	err = query.Q.Transaction(func(tx *query.Query) error {
		signInDay, err = tx.QrSignInDay.Where(d.ID.Eq(signInDatum.DayID)).Attrs(field.Attrs(&model.QrSignInDay{
			Qq:        req.QQ,
			MonthDays: 1,
			TotalDays: 1,
		})).FirstOrCreate()
		return err
	})
	if err != nil {
		return v1.SignInDataResponse{}, err
	}

	if signInDatum.SignInDate.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return v1.SignInDataResponse{
			Status:  203,
			Message: "已经签到",
			Data:    s.FillDetailResponse(*signInDatum, *signInDay),
		}, nil
	} else {
		err := s.UpdateSignInDataAndDays(ctx, req.Points, *signInDatum, *signInDay)
		if err != nil {
			return v1.SignInDataResponse{}, err
		}
		return v1.SignInDataResponse{
			Status:  200,
			Message: "签到成功",
			Data:    s.FillDetailResponse(*signInDatum, *signInDay),
		}, nil
	}
}

func (s *signInService) AddSignInPoints(req v1.AddSignPointsRequest) (v1.SignInDataResponse, error) {
	update, err := query.QrSignInDatum.Where(query.QrSignInDatum.Qq.Eq(req.QQ)).Update(query.QrSignInDatum.Points, query.QrSignInDatum.Points.Add(req.Points))
	if err != nil {
		return v1.SignInDataResponse{}, err
	}
	if update.RowsAffected == 0 {
		return v1.SignInDataResponse{}, v1.ErrNotFound
	}
	signInDetailResponse, err := s.QuerySignInData(v1.QuerySignInDataRequest{QQ: req.QQ})
	if err != nil {
		return v1.SignInDataResponse{}, err
	}
	return v1.SignInDataResponse{
		Status:  200,
		Message: "操作成功",
		Data:    signInDetailResponse,
	}, nil
}

func (s *signInService) SaveSignInData(qq string, points int32) (model.QrSignInDatum, model.QrSignInDay, error) {
	day := model.QrSignInDay{
		Qq:        qq,
		MonthDays: 1,
		TotalDays: 1,
	}
	qqNumber, _ := strconv.ParseInt(qq, 10, 64)
	creator := "admin"
	data := model.QrSignInDatum{
		Qq:         qq,
		QqNumber:   &qqNumber,
		Points:     points,
		SignInDate: db.Now(),
		CreateBy:   &creator,
		UpdateBy:   &creator,
	}
	err := query.Q.Transaction(func(tx *query.Query) error {
		err := tx.QrSignInDay.Create(&day)
		if err != nil {
			return err
		}
		data.DayID = day.ID
		err = tx.QrSignInDatum.Create(&data)
		return err
	})

	if err != nil {
		return model.QrSignInDatum{}, model.QrSignInDay{}, err
	}
	return data, day, nil
}

func (s *signInService) UpdateSignInDataAndDays(ctx *gin.Context, addPoints int32, datum model.QrSignInDatum, day model.QrSignInDay) error {
	q := query.QrSignInDatum
	d := query.QrSignInDay

	err := query.Q.Transaction(func(tx *query.Query) error {
		_, err := tx.QrSignInDatum.WithContext(ctx).Where(q.ID.Eq(datum.ID)).Updates(map[string]interface{}{"Points": datum.Points + addPoints, "Sign_In_Date": time.Now()})
		if err != nil {
			return err
		}
		_, err = tx.QrSignInDay.WithContext(ctx).Where(d.ID.Eq(day.ID)).Updates(map[string]interface{}{"Month_Days": day.MonthDays + 1, "Total_Days": day.TotalDays + 1})
		if err != nil {
			return err
		}
		err = s.userAssetsService.AddUserPoints(datum.Qq, addPoints)
		return err
	})
	return err
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
