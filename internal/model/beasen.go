package model

import (
	"gorm.io/plugin/soft_delete"
	"robot-system-server/pkg/db"
)

type Beasen struct {
	Id         uint                  `gorm:"primarykey" json:"id,omitempty"`
	Sentence   string                `json:"sentence,omitempty"`
	Md5        string                `json:"md5,omitempty"`
	Type       string                `json:"type,omitempty"`
	Source     string                `json:"source,omitempty"`
	CreateBy   string                `json:"createBy,omitempty"`
	UpdateBy   string                `json:"updateBy,omitempty"`
	CreateTime *db.LocalDateTime     `json:"createTime"`
	UpdateTime *db.LocalDateTime     `json:"updateTime"`
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;comment:删除标识;0否1是;softDelete:flag" json:"delFlag"`
}

func (m *Beasen) TableName() string {
	return "qr_beasen"
}
