package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type PmsBrand struct {
	Id        int64                 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time            `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time            `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time            `gorm:"column:deleted_at;" json:"deletedAt"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`

	// 品牌名称
	Name string `gorm:"column:name" json:"name"`
	// 首字母
	FirstLetter string `gorm:"column:first_letter" json:"firstLetter"`
	// 排序
	Sort int `gorm:"column:sort" json:"sort"`
	// 是否为品牌制造商：0->不是；1->是
	FactoryStatus int `gorm:"column:factory_status" json:"factoryStatus"`
	// 是否显示
	ShowStatus int `gorm:"column:show_status" json:"showStatus"`
	// 产品数量
	ProductCount int `gorm:"column:product_count" json:"productCount"`
	// 产品评论数量
	ProductCommentCount int `gorm:"column:product_comment_count" json:"productCommentCount"`
	// 品牌logo
	Logo string `gorm:"column:logo" json:"logo"`
	// 专区大图
	BigPic string `gorm:"column:big_pic" json:"bigPic"`
	// 品牌故事
	BrandStory string `gorm:"column:brand_story" json:"brandStory"`
}

func (*PmsBrand) TableName() string {
	return "pms_brand"
}
