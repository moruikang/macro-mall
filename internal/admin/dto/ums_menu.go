// @Author moruikang
// @Date 2025/3/16 17:46:00
// @Desc

package dto

type MenuCreateDTO struct {
	// 父级ID
	ParentId int64 `json:"parentId"`
	// 菜单名称
	Title string `json:"title"`
	// 菜单级数
	Level int `json:"level"`
	// 菜单排序
	Sort int `json:"sort"`
	// 前端名称
	Name string `json:"name"`
	// 前端图标
	Icon string `json:"icon"`
	// 前端隐藏
	Hidden int `json:"hidden"`
}

type MenuUpdateHiddenDTO struct {
	// 前段隐藏菜单
	Hidden int `json:"hidden" query:"hidden" form:"hidden"`
}
