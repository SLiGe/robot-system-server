package service

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/pkg/db"
	"strconv"
	"time"
)

type SpiritSignService interface {
	OneSignPerDay(req v1.SignPerDayReq) (v1.SignPerDayRes, error)
}

func NewSpiritSignService(service *Service, userService UserService) SpiritSignService {
	return &spiritSignService{
		Service:     service,
		userService: userService,
	}
}

type spiritSignService struct {
	*Service
	userService UserService
}

var signTypeList = []string{"gylq", "yllq", "cslq"}

func (s *spiritSignService) OneSignPerDay(req v1.SignPerDayReq) (v1.SignPerDayRes, error) {
	if !utils.Contains(signTypeList, req.Type) || req.Qq == "" {
		return v1.SignPerDayRes{}, v1.ErrBadRequest
	}
	u := query.QrSpiritSignUDatum
	var signId int64
	err := u.UnderlyingDB().Select("id").Where("user_qq =? and sign_type =? and sign_date =?", req.Qq, req.Type, time.Now().Format("2006-01-02")).Limit(1).First(&signId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return v1.SignPerDayRes{}, err
	}
	ss := query.QrSpiritSign
	if errors.Is(err, gorm.ErrRecordNotFound) {
		spiritSign, err := ss.Select().Order(field.Func.Rand()).Limit(1).First()
		if err != nil {
			return v1.SignPerDayRes{}, err
		}
		user, err := s.userService.QueryOrCreate(req.Qq)
		if err != nil {
			return v1.SignPerDayRes{}, err
		}
		spiritSignUDatum := model.QrSpiritSignUDatum{
			UserID:   &user.UserID,
			UserQq:   &req.Qq,
			SignType: &req.Type,
			SignID:   &spiritSign.ID,
			SignDate: db.Now(),
		}
		err = u.Create(&spiritSignUDatum)
		if err != nil {
			return v1.SignPerDayRes{}, err
		}
	}
	spiritSign, err := ss.Select(ss.ALL).Where(ss.ID.Eq(signId)).First()
	if err != nil {
		return v1.SignPerDayRes{}, err
	}

	return v1.SignPerDayRes{
		SignData: chooseSignType(req.Type, *spiritSign.DataJSON),
		ViewUrl:  "",
	}, nil
}

func chooseSignType(typeStr string, dataJson string) any {
	if typeStr == "gylq" {
		var gylqRes Gylq
		err := jsoniter.UnmarshalFromString(dataJson, &gylqRes)
		if err != nil {
			return nil
		}
		return gylqRes
	} else if typeStr == "yllq" {
		var yllqRes Yllq
		err := jsoniter.UnmarshalFromString(dataJson, &yllqRes)
		if err != nil {
			return nil
		}
		return yllqRes
	} else if typeStr == "cslq" {
		var cslqRes Cslq
		err := jsoniter.UnmarshalFromString(dataJson, &cslqRes)
		if err != nil {
			return nil
		}
		if cslqRes.Id == 61 {
			cslqRes.Title = "财神灵签 签王解签"
		} else {
			cslqRes.Title = "财神灵签第" + strconv.Itoa(cslqRes.Id) + "签解签 财神灵签第" +
				message.NewPrinter(language.Chinese).Sprintf("%d", cslqRes.Id) + "签"

		}

		return cslqRes
	}
	return nil
}

type Gylq struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ShiYue   string `json:"shi_yue"`
	ShiYi    string `json:"shi_yi"`
	JieYue   string `json:"jie_yue"`
	XiangJie string `json:"xiang_jie"`
	XianJi   string `json:"xian_ji"`
	LqImg    string `json:"lq_img"`
	Ztjq     string `json:"ztjq"`
	Bqjs     string `json:"bqjs"`
	Fszs     string `json:"fszs"`
	Aqhy     string `json:"aqhy"`
	Gzqz     string `json:"gzqz"`
	Ksjs     string `json:"ksjs"`
	Tzlc     string `json:"tzlc"`
	Jssy     string `json:"jssy"`
	Fdjy     string `json:"fdjy"`
	Zbjk     string `json:"zbjk"`
	Zhbg     string `json:"zhbg"`
	Qyqz     string `json:"qyqz"`
	Gsss     string `json:"gsss"`
	Xrxw     string `json:"xrxw"`
	Yxcg     string `json:"yxcg"`
	QsgsList []struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"qsgsList"`
}

type Yllq struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	QianShi string `json:"qian_shi"`
	LqImg   string `json:"lq_img"`
	JieQian string `json:"jie_qian"`
	Zhu     string `json:"zhu"`
	Bhqs    string `json:"bhqs"`
}

type Cslq struct {
	Title   string `json:"title"`
	Id      int    `json:"id"`
	Desc    string `json:"desc"`
	ShiYue  string `json:"shi_yue"`
	LqImg   string `json:"lq_img"`
	Mlxz    string `json:"mlxz"`
	JiXiong string `json:"ji_xiong"`
}
