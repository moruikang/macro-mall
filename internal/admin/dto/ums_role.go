// @Author moruikang
// @Date 2025/3/16 17:46:00
// @Desc

package dto

type RoleCreateDTO struct {
	// 角色名称
	Name string `json:"name" binding:"required"`
	// 角色描述
	Description string `json:"description"`
	// 后台用户数量
	AdminCount int `json:"adminCount"`
	// 启用状态：0->禁用；1->启用
	Status int `json:"status"`
	// 排序
	Sort int `json:"sort"`
}

type RoleDeleteDTO struct {
	// 角色id列表
	Ids []int64 `json:"ids" binding:"required"`
}

type AdminRoleStatusDTO struct {
	// 角色启用状态 0->禁用，1->启用
	Status int `json:"status" form:"status"`
}

type RoleAllocMenuDTO struct {
	// 角色id
	RoleId int64 `json:"roleId" form:"roleId"`
	// 菜单id列表
	MenuIds []int64 `json:"menuIds" form:"menuIds"`
}

type RoleAllocResourceDTO struct {
	// 角色id列表
	RoleId int64 `json:"roleId" form:"roleId"`
	// 资源id列表
	ResourceIds []int64 `json:"resourceIds" form:"resourceIds"`
}
