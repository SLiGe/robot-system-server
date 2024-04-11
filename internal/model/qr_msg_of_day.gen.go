// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"robot-system-server/pkg/db"
	"time"
)

const TableNameQrMsgOfDay = "qr_msg_of_day"

// QrMsgOfDay 每日一句
type QrMsgOfDay struct {
	ID         int64             `gorm:"column:ID;primaryKey;autoIncrement:true" json:"id"`
	Sentence   *string           `gorm:"column:SENTENCE" json:"sentence"`
	CreateDate *db.LocalDateTime `gorm:"column:CREATE_DATE;autoCreateTime" json:"createDate"`
	SendDate   *time.Time        `gorm:"column:SEND_DATE;comment:发送日期" json:"sendDate"` // 发送日期
}

// TableName QrMsgOfDay's table name
func (*QrMsgOfDay) TableName() string {
	return TableNameQrMsgOfDay
}