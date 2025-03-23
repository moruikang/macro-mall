// @Author moruikang
// @Date 2025/3/23 09:16:00
// @Desc

package v1

import (
	"github.com/gin-gonic/gin"
	"macro-mall/internal/admin/dto"
	v1 "macro-mall/internal/admin/service/v1"
	"macro-mall/internal/admin/store"
	"macro-mall/pkg/response"
)

type UmsAdminController struct {
	svc v1.Service
}

func NewUmsAdminController(store store.Factory) *UmsAdminController {
	return &UmsAdminController{
		svc: v1.NewService(store),
	}
}

// @Summary 管理员注册
// @Description 管理员注册
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body dto.AdminRegisterDTO true "管理员注册"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/register [post]
func (c *UmsAdminController) Register(g *gin.Context) {

	var umsAdmin *dto.AdminRegisterDTO
	if err := g.ShouldBindJSON(umsAdmin); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	if err := c.svc.UmsAdmins().Register(g, umsAdmin); err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body dto.AdminLoginDTO true "管理员登录"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/login [post]
func (c *UmsAdminController) Login(g *gin.Context) {
	var umsAdmin *dto.AdminLoginDTO
	if err := g.ShouldBindJSON(umsAdmin); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	token, err := c.svc.UmsAdmins().Login(g, umsAdmin)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, gin.H{"token": token})
	return
}

// @Summary 获取管理员信息
// @Description 获取管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/info [get]
func (c *UmsAdminController) GetAdminInfo(g *gin.Context) {
	adminInfo, err := c.svc.UmsAdmins().GetAdminInfo(g)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, adminInfo)
	return
}

// @Summary 更新管理员信息
// @Description 更新管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body dto.AdminUpdateDTO true "管理员信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/update [post]
func (c *UmsAdminController) UpdateAdminInfo(g *gin.Context) {
	var umsAdmin *dto.AdminUpdateDTO
	if err := g.ShouldBindJSON(umsAdmin); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err := c.svc.UmsAdmins().UpdateAdminInfo(g, umsAdmin)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 删除管理员
// @Description 删除管理员
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/delete/{id} [delete]
func (c *UmsAdminController) DeleteAdmin(g *gin.Context) {
	id := g.Param("id")
	err := c.svc.UmsAdmins().DeleteAdmin(g, id)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}
