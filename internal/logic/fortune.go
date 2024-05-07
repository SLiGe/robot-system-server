package logic

import (
	"github.com/duke-git/lancet/v2/strutil"
	jsoniter "github.com/json-iterator/go"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/pkg/db"
)

type FortuneLogic interface {
	GetFortune() *model.QrFortune
	CostUserPoints(qq string)
	UpdateGroupNum(groupNum string, datum model.QrFortuneDatum)
	SaveNewFortuneData(qq string, groupNum string, isOne int, isGroup int, fortune *model.QrFortune)
}

type fortuneLogic struct {
	*Logic
}

func NewFortuneLogic(logic *Logic) FortuneLogic {
	return &fortuneLogic{
		Logic: logic,
	}
}

func (l *fortuneLogic) GetFortune() *model.QrFortune {
	q := query.QrFortune
	var fortuneData model.QrFortune
	if err := q.UnderlyingDB().Order("rand()").First(&fortuneData).Error; err != nil {
		return nil
	}
	return &fortuneData
}

func (l *fortuneLogic) CostUserPoints(qq string) {
	q := query.QrSignInDatum
	signInDatum, err := q.Select().Where(q.Qq.Eq(qq)).First()
	if err != nil {
		return
	}
	if signInDatum.Points > 10 {
		_, err := q.Where(q.ID.Eq(signInDatum.ID)).Update(q.Points, signInDatum.Points-10)
		if err != nil {
			return
		}
	}
}

func (l *fortuneLogic) UpdateGroupNum(groupNum string, datum model.QrFortuneDatum) {
	if strutil.IsNotBlank(groupNum) && "null" != groupNum {
		qrFortuneDataGroupNum := ""
		if datum.GroupNum != nil && strutil.IsNotBlank(*datum.GroupNum) {
			qrFortuneDataGroupNum = *datum.GroupNum
		}
		q := query.QrFortuneDatum
		qrFortuneDataGroupNum = qrFortuneDataGroupNum + groupNum + ","
		_, err := q.Where(q.ID.Eq(datum.ID)).Update(q.GroupNum, qrFortuneDataGroupNum)
		if err != nil {
			return
		}
	}
}

func (l *fortuneLogic) SaveNewFortuneData(qq string, groupNum string, isOne int, isGroup int, fortune *model.QrFortune) {
	q := query.QrFortuneDatum
	fortuneJson, _ := jsoniter.MarshalToString(fortune)
	if fortuneData, err := q.Where(q.Qq.Eq(qq)).First(); err != nil {
		var newFortuneData = model.QrFortuneDatum{
			Qq:          qq,
			FortuneDate: db.Now(),
			JSONData:    fortuneJson,
		}
		if isOne == 1 && isGroup == 1 && (strutil.IsNotBlank(groupNum) && "null" != groupNum) {
			*newFortuneData.GroupNum = groupNum + ","
		}
		_ = q.Create(&newFortuneData)
	} else {
		fortuneData.JSONData = fortuneJson
		fortuneData.FortuneDate = db.Now()
		if isOne == 1 && isGroup == 1 && (strutil.IsNotBlank(groupNum) && "null" != groupNum) {
			qrFortuneDataGroupNum := ""
			if fortuneData.GroupNum != nil && strutil.IsNotBlank(*fortuneData.GroupNum) {
				qrFortuneDataGroupNum = *fortuneData.GroupNum
			}
			*fortuneData.GroupNum = qrFortuneDataGroupNum + groupNum + ","
		}
		_, _ = q.Where(q.ID.Eq(fortuneData.ID)).Updates(fortuneData)
	}
}
