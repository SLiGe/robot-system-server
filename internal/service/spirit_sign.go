package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/pkg/db"
	"strconv"
	"time"
)

type SpiritSignService interface {
	OneSignPerDay(req v1.SignPerDayReq) (v1.SignPerDayRes, error)
	ViewGyLq(ctx *gin.Context)
	ViewCsLq(ctx *gin.Context)
	ViewYlLq(ctx *gin.Context)
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
	var localSign model.QrSpiritSignUDatum
	err := u.UnderlyingDB().Where("user_qq =? and sign_type =? and sign_date =?", req.Qq, req.Type, time.Now().Format("2006-01-02")).Limit(1).First(&localSign).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return v1.SignPerDayRes{}, err
	}
	var signId = localSign.SignID
	ss := query.QrSpiritSign
	if errors.Is(err, gorm.ErrRecordNotFound) {
		spiritSign, err := ss.Select().Where(ss.DataType.Eq(req.Type)).Order(field.Func.Rand()).Limit(1).First()
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
		signId = &spiritSign.ID
		err = u.Create(&spiritSignUDatum)
		if err != nil {
			return v1.SignPerDayRes{}, err
		}
	}
	spiritSign, err := ss.Select(ss.ALL).Where(ss.ID.Eq(*signId)).First()
	if err != nil {
		return v1.SignPerDayRes{}, err
	}

	path, data := chooseSignType(req.Type, *spiritSign.DataJSON)
	return v1.SignPerDayRes{
		SignData: data,
		ViewUrl:  fmt.Sprintf("/lq/v/%s/%s", path, req.Qq),
	}, nil
}

func (s *spiritSignService) ViewGyLq(ctx *gin.Context) {
	qq := ctx.Param("qq")
	signPerDayRes, err := s.OneSignPerDay(v1.SignPerDayReq{
		Type: signTypeList[0],
		Qq:   qq,
	})
	if err != nil {
		return
	}
	ctx.HTML(http.StatusOK, "lq/gylq.html", gin.H{
		"data":        signPerDayRes.SignData,
		"newLineChar": "\\n",
	})
}

func (s *spiritSignService) ViewCsLq(ctx *gin.Context) {
	qq := ctx.Param("qq")
	signPerDayRes, err := s.OneSignPerDay(v1.SignPerDayReq{
		Type: signTypeList[2],
		Qq:   qq,
	})
	if err != nil {
		return
	}
	var cslq = signPerDayRes.SignData.(Cslq)
	var title string
	if cslq.Id == 61 {
		title = "财神灵签 签王解签"
	} else {
		title = "财神灵签第" + strconv.Itoa(cslq.Id) + "签解签 财神灵签" + message.NewPrinter(language.Chinese).Sprintf("第%d签", cslq.Id)
	}
	ctx.HTML(http.StatusOK, "lq/cslq.html", gin.H{
		"data":        signPerDayRes.SignData,
		"newLineChar": "\\n",
		"title":       title,
	})
}

func (s *spiritSignService) ViewYlLq(ctx *gin.Context) {
	qq := ctx.Param("qq")
	signPerDayRes, err := s.OneSignPerDay(v1.SignPerDayReq{
		Type: signTypeList[1],
		Qq:   qq,
	})
	if err != nil {
		return
	}
	ctx.HTML(http.StatusOK, "lq/yllq.html", gin.H{
		"data":        signPerDayRes.SignData,
		"newLineChar": "\\n",
	})
}

func chooseSignType(typeStr string, dataJson string) (path string, data any) {
	if typeStr == "gylq" {
		var gylqRes Gylq
		err := jsoniter.UnmarshalFromString(dataJson, &gylqRes)
		if err != nil {
			return "gy", nil
		}
		return "gy", gylqRes
	} else if typeStr == "yllq" {
		var yllqRes Yllq
		err := jsoniter.UnmarshalFromString(dataJson, &yllqRes)
		if err != nil {
			return "yl", nil
		}
		return "yl", yllqRes
	} else if typeStr == "cslq" {
		var cslqRes Cslq
		err := jsoniter.UnmarshalFromString(dataJson, &cslqRes)
		if err != nil {
			return "cs", nil
		}
		if cslqRes.Id == 61 {
			cslqRes.Title = "财神灵签 签王解签"
		} else {
			cslqRes.Title = "财神灵签第" + strconv.Itoa(cslqRes.Id) + "签解签 财神灵签第" +
				message.NewPrinter(language.Chinese).Sprintf("%d", cslqRes.Id) + "签"

		}

		return "cs", cslqRes
	}
	return "", nil
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
