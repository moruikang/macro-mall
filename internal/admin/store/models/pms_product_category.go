package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type PmsProductCategory struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deletedAt"`
	// 原有表结构字段 用于兼容
	CreateTime *time.Time            `gorm:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`
	// 父分类的编号
	ParentId int64 `json:"parentId" gorm:"column:parent_id"`
	// 分类名称
	Name string `json:"name" gorm:"column:name"`
	// 分类级别
	Level int `json:"level" gor:"column:level"`
	// 分类单位
	ProductUnit string `json:"productUnit" gorm:"column:product_unit"`
	// 分类数量
	ProductCount int `json:"productCount" gorm:"column:product_count"`
	// 是否显示在导航栏
	NavStatus int `json:"navStatus" gorm:"column:nav_status"`
	// 显示状态
	ShowStatus int `json:"showStatus" gorm:"column:show_status"`
	// 排序
	Sort int `json:"sort" gorm:"column:sort"`
	// 图标
	Icon string `json:"icon" gorm:"column:icon"`
	// 关键字
	Keywords string `json:"keywords" gorm:"column:keywords"`
	// 描述
	Description string `json:"description" gorm:"column:description"`
}

type PmsProductCategoryWithChildrenItem struct {
	PmsProductCategory
	Children []*PmsProductCategory `json:"children" gorm:"foreignKey:ParentId"`
}

func (*PmsProductCategory) TableName() string {
	return "pms_product_category"
}
