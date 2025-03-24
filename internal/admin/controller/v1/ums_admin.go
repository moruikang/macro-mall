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
	"strconv"
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

// @Summary 获取管理员信息
// @Description 获取管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=map[string]interface{}} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/info [get]
func (c *UmsAdminController) GetAdminInfo(g *gin.Context) {
	username := g.GetHeader("username")
	adminInfo, err := c.svc.UmsAdmins().GetAdminInfo(g, username)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, adminInfo)
	return
}

// @Summary 获取管理员信息
// @Description 获取管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Success 200 {object} response.Response{data=models.UmsAdmin} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/:id [get]
func (c *UmsAdminController) GetAdmin(g *gin.Context) {
	userId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	umsAdmin, err := c.svc.UmsAdmins().Get(g, userId)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsAdmin)
	return
}

// @Summary 获取管理员角色
// @Description 获取管理员角色
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Success 200 {object} response.Response{data=[]models.UmsRole} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/role/:id [get]
func (c *UmsAdminController) GetAdminRoles(g *gin.Context) {
	userId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	userRoleList, err := c.svc.UmsAdmins().GetUserRoleList(g, userId)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, userRoleList)
	return
}

// @Summary 分页查询管理员列表
// @Description 分页查询管理员列表
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data param dto.PublicPageQuery true "分页查询参数"
// @Success 200 {object} response.Response{data=dto.PageResult} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/list [get]
func (c *UmsAdminController) ListAdmin(g *gin.Context) {

	var query *dto.PublicPageQuery
	if err := g.ShouldBindQuery(query); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	umsAdmins, err := c.svc.UmsAdmins().List(g, query)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsAdmins)
	return
}

// @Summary 更新管理员信息
// @Description 更新管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Param data body dto.AdminUpdateDTO true "管理员信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/update/:id [post]
func (c *UmsAdminController) UpdateAdmin(g *gin.Context) {
	var umsAdmin *dto.AdminUpdateDTO
	if err := g.ShouldBindJSON(umsAdmin); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	userId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	umsAdmin.Id = userId
	err = c.svc.UmsAdmins().Update(g, umsAdmin)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 更新管理员信息
// @Description 更新管理员信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Param data body dto.AdminRoleRelationDTO true "管理员信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/role/update [post]
func (c *UmsAdminController) UpdateAdminRole(g *gin.Context) {
	var req *dto.AdminRoleRelationDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err := c.svc.UmsAdmins().UpdateRoles(g, req)
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
// @Router /admin/delete/{id} [post]
func (c *UmsAdminController) DeleteAdmin(g *gin.Context) {
	userId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsAdmins().Delete(g, []int64{userId})
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 更新管理员密码
// @Description 更新管理员密码
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body dto.AdminUpdatePasswordDTO true "管理员信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/updatePassword [post]
func (c *UmsAdminController) UpdatePassword(g *gin.Context) {
	var req *dto.AdminUpdatePasswordDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err := c.svc.UmsAdmins().UpdatePassword(g, req)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 更新管理员状态
// @Description 更新管理员状态
// @Tags 管理员
// @Accept json
// @Produce json
// @Param id path int true "管理员ID"
// @Param data body dto.AdminUpdateStatusDTO true "管理员信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /admin/updateStatus/:id [put]
func (c *UmsAdminController) UpdateStatus(g *gin.Context) {
	var req *dto.AdminUpdateStatusDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	userId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsAdmins().UpdateStatus(g, userId, req)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}
