// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"robot-system-server/pkg/db"
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameQrSignInLevel = "qr_sign_in_level"

// QrSignInLevel 等级表
type QrSignInLevel struct {
	ID         int64                 `gorm:"column:ID;primaryKey;autoIncrement:true;comment:主键" json:"id"`                        // 主键
	Level      *string               `gorm:"column:LEVEL;comment:等级" json:"level"`                                                // 等级
	MaxPoints  *int64                `gorm:"column:MAX_POINTS;comment:小于积分" json:"maxPoints"`                                     // 小于积分
	MinPoints  *int64                `gorm:"column:MIN_POINTS;comment:大于积分" json:"minPoints"`                                     // 大于积分
	GroupID    *int64                `gorm:"column:group_id;comment:群组ID" json:"groupId"`                                         // 群组ID
	UpdateDate *db.LocalDateTime     `gorm:"column:UPDATE_DATE;autoUpdateTime;comment:更新日期" json:"updateDate"`                    // 更新日期
	CreateDate *db.LocalDateTime     `gorm:"column:CREATE_DATE;autoCreateTime;comment:创建日期" json:"createDate"`                    // 创建日期
	CreateBy   *string               `gorm:"column:create_by;comment:创建人" json:"createBy"`                                        // 创建人
	CreateTime *db.LocalDateTime     `gorm:"column:create_time;autoCreateTime;autoUpdateTime;comment:创建时间" json:"createTime"`     // 创建时间
	UpdateBy   *string               `gorm:"column:update_by;comment:更新人" json:"updateBy"`                                        // 更新人
	UpdateTime *time.Time            `gorm:"column:update_time;comment:更新时间" json:"updateTime"`                                   // 更新时间
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;not null;default:0;comment:删除标识;0否1是;softDelete:flag" json:"delFlag"` // 删除标识;0否1是
}

// TableName QrSignInLevel's table name
func (*QrSignInLevel) TableName() string {
	return TableNameQrSignInLevel
}
