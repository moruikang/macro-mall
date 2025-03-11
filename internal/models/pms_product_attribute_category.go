package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type PmsProductAttributeCategory struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deletedAt"`
	// 原有表结构字段 用于兼容
	CreateTime *time.Time            `gorm:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`
	Name       string                `json:"value" gorm:"value"`
	// 属性数量
	AttributeCount int `json:"attributeCount" gorm:"attribute_count"`
	// 参数数量
	ParamCount int `json:"paramCount" gorm:"param_count"`
}
