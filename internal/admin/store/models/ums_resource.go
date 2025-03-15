package models

type UmsResource struct {
	Model
	// 资源名称
	Name string `gorm:"column:name;not null" json:"name"`
	// 资源URL
	Url string `gorm:"column:url;not null" json:"url"`
	// 描述
	Description string `gorm:"column:description;" json:"description"`
	// 资源分类ID
	CategoryId int64 `gorm:"column:category_id;not null" json:"categoryId"`
	// 对应的拥有该资源的角色，在数据库忽略该字段
	RoleId int64 `gorm:"-" json:"roleId"`
}

func (*UmsResource) TableName() string {
	return "ums_resource"
}
