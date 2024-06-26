// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"robot-system-server/pkg/db"

	"gorm.io/plugin/soft_delete"
)

const TableNameQrSignInDay = "qr_sign_in_days"

// QrSignInDay 签到天数关联表
type QrSignInDay struct {
	ID         int32                 `gorm:"column:ID;primaryKey;autoIncrement:true;comment:主键" json:"id"`                        // 主键
	Qq         string                `gorm:"column:QQ;not null;comment:用户QQ" json:"qq"`                                           // 用户QQ
	MonthDays  int32                 `gorm:"column:MONTH_DAYS;not null;comment:月连续签到天数" json:"monthDays"`                         // 月连续签到天数
	TotalDays  int32                 `gorm:"column:TOTAL_DAYS;not null;comment:总天数" json:"totalDays"`                             // 总天数
	UpdateDate *db.LocalDateTime     `gorm:"column:UPDATE_DATE;autoUpdateTime;comment:更新日期" json:"updateDate"`                    // 更新日期
	CreateDate *db.LocalDateTime     `gorm:"column:CREATE_DATE;autoCreateTime;comment:创建日期" json:"createDate"`                    // 创建日期
	CreateBy   *string               `gorm:"column:create_by;comment:创建人" json:"createBy"`                                        // 创建人
	CreateTime *db.LocalDateTime     `gorm:"column:create_time;autoCreateTime;comment:创建时间" json:"createTime"`                    // 创建时间
	UpdateBy   *string               `gorm:"column:update_by;comment:更新人" json:"updateBy"`                                        // 更新人
	UpdateTime *db.LocalDateTime     `gorm:"column:update_time;autoUpdateTime;comment:更新时间" json:"updateTime"`                    // 更新时间
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;not null;default:0;comment:删除标识;0否1是;softDelete:flag" json:"delFlag"` // 删除标识;0否1是
}

// TableName QrSignInDay's table name
func (*QrSignInDay) TableName() string {
	return TableNameQrSignInDay
}
