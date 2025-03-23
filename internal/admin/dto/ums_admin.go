// @Author moruikang
// @Date 2025/3/16 17:37:00
// @Desc

package dto

type AdminBaseDTO struct {
	// 账号
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}

type AdminCreateDTO struct {
	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 密码
	Password string `json:"password" form:"password" binding:"required"`
	// 图标
	Icon string `json:"icon,omitempty" form:"icon"`
	// 昵称
	Nickname string `json:"nickname,omitempty" form:"nickname"`
	// 邮箱
	Email string `json:"email,omitempty" form:"email"`
	// 备注
	Note string `json:"note,omitempty" form:"note"`
}

type AdminLoginDTO struct {
	AdminBaseDTO
}

type AdminRegisterDTO struct {
	AdminCreateDTO
}

type AdminUpdateDTO struct {
	Base
	AdminCreateDTO
}

type AdminUpdatePasswordDTO struct {
	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 旧密码
	OldPassword string `json:"oldPassword" form:"oldPassword" binding:"required"`
	// 新密码
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}

type AdminUpdateStatusDTO struct {
	// 账号启用状态 0->禁用，1->启用
	Status int `json:"status" form:"status"`
}

type AdminRoleRelationDTO struct {
	//
	AdminId int64 `json:"adminId" form:"adminId"`
	//
	RoleIds []int64 `json:"roleIds" form:"roleIds"`
}
