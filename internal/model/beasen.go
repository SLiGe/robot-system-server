package model

import "time"

type Beasen struct {
	Id         uint      `gorm:"primarykey" json:"id,omitempty"`
	Sentence   string    `json:"sentence,omitempty"`
	Md5        string    `json:"md5,omitempty"`
	Type       string    `json:"type,omitempty"`
	Source     string    `json:"source,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	DelFlag    string    `json:"delFlag,omitempty"`
}

func (m *Beasen) TableName() string {
	return "qr_beasen"
}
