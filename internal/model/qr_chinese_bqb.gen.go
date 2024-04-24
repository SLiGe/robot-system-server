// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"robot-system-server/pkg/db"

	"gorm.io/plugin/soft_delete"
)

const TableNameQrChineseBqb = "qr_chinese_bqb"

// QrChineseBqb 表情包
type QrChineseBqb struct {
	ID         int64                 `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`     // 主键
	ImgName    *string               `gorm:"column:img_name;comment:图片名称" json:"imgName"`                      // 图片名称
	ImgSize    *int64                `gorm:"column:img_size;comment:图片大小" json:"imgSize"`                      // 图片大小
	ImgKeyword *string               `gorm:"column:img_keyword;comment:关键词" json:"imgKeyword"`                 // 关键词
	ImgClass   *string               `gorm:"column:img_class;comment:图片分类" json:"imgClass"`                    // 图片分类
	ImgURL     *string               `gorm:"column:img_url;comment:图片地址" json:"imgUrl"`                        // 图片地址
	Md5        *string               `gorm:"column:md5;comment:MD5" json:"md5"`                                // MD5
	CreateBy   *string               `gorm:"column:create_by;comment:创建人" json:"createBy"`                     // 创建人
	CreateTime *db.LocalDateTime     `gorm:"column:create_time;autoCreateTime;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy   *string               `gorm:"column:update_by;comment:更新人" json:"updateBy"`                     // 更新人
	UpdateTime *db.LocalDateTime     `gorm:"column:update_time;autoUpdateTime;comment:更新时间" json:"updateTime"` // 更新时间
	DelFlag    soft_delete.DeletedAt `gorm:"column:del_flag;comment:删除标识;0否1是;softDelete:flag" json:"delFlag"` // 删除标识;0否1是
}

// TableName QrChineseBqb's table name
func (*QrChineseBqb) TableName() string {
	return TableNameQrChineseBqb
}
