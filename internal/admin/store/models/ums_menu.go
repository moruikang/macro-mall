package models

type UmsMenu struct {
	Model
	// 父级ID
	ParentId int64 `gorm:"column:parent_id;" json:"parentId"`
	// 菜单名称
	Title string `gorm:"column:title;not null" json:"title"`
	// 菜单级数
	Level int `gorm:"column:level;not null" json:"level"`
	// 菜单排序
	Sort int `gorm:"column:sort;default:0" json:"sort"`
	// 前端名称
	Name string `gorm:"column:name;" json:"name"`
	// 前端图标
	Icon string `gorm:"column:icon;" json:"icon"`
	// 前端隐藏
	Hidden int `gorm:"column:hidden;default:0" json:"hidden"`
}

type UmsMenuNode struct {
	UmsMenu
	Children []*UmsMenuNode `json:"children"`
}

func (*UmsMenu) TableName() string {
	return "ums_menu"
}
