package models

type UmsRole struct {
	*Model
	// 名称
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	// 描述
	Description string `gorm:"column:description;type:varchar(500);" json:"description"`
	// 后台用户数量
	AdminCount int `gorm:"column:admin_count;type:int(11);" json:"adminCount"`
	// 0->禁用；1->启用
	Status int `gorm:"column:status;type:tinyint(1);default:1;" json:"status"`
	// 排序
	Sort int `gorm:"column:sort;type:int(11);default:0;" json:"sort"`
}

func (*UmsRole) TableName() string {
	return "ums_role"
}
