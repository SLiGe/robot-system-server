// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameQrFortune = "qr_fortune"

// QrFortune 运势表
type QrFortune struct {
	ID             int64  `gorm:"column:ID;primaryKey;autoIncrement:true" json:"id"`
	FortuneSummary string `gorm:"column:FORTUNE_SUMMARY;not null;comment:运情总结" json:"fortuneSummary"` // 运情总结
	LuckyStar      string `gorm:"column:LUCKY_STAR;not null;comment:幸运星" json:"luckyStar"`            // 幸运星
	SignText       string `gorm:"column:SIGN_TEXT;not null;comment:签文" json:"signText"`               // 签文
	UnSignText     string `gorm:"column:UN_SIGN_TEXT;not null;comment:解签" json:"unSignText"`          // 解签
}

// TableName QrFortune's table name
func (*QrFortune) TableName() string {
	return TableNameQrFortune
}
