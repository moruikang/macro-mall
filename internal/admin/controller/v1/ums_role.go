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

type UmsRoleController struct {
	svc v1.Service
}

func NewUmsRoleController(store store.Factory) *UmsRoleController {
	return &UmsRoleController{
		svc: v1.NewService(store),
	}
}

// @Summary 创建角色
// @Description 创建角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param data body dto.RoleCreateDTO true "创建角色"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/create [post]
func (c *UmsRoleController) Create(g *gin.Context) {

	var umsRole *dto.RoleCreateDTO
	if err := g.ShouldBindJSON(umsRole); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	if err := c.svc.UmsRoles().Create(g, umsRole); err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 分页查询角色列表
// @Description 分页查询角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param data param dto.PublicPageQuery true "分页查询参数"
// @Success 200 {object} response.Response{data=dto.PageResult} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/list [get]
func (c *UmsRoleController) ListRole(g *gin.Context) {

	var query *dto.PublicPageQuery
	if err := g.ShouldBindQuery(query); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	umsRoles, err := c.svc.UmsRoles().List(g, query)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsRoles)
	return
}

// @Summary 更新角色信息
// @Description 更新角色信息
// @Tags 角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleCreateDTO true "角色信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/update/:id [post]
func (c *UmsRoleController) UpdateRole(g *gin.Context) {
	var umsRole *dto.RoleCreateDTO
	if err := g.ShouldBindJSON(umsRole); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	roleId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err = c.svc.UmsRoles().Update(g, roleId, umsRole)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 删除角色
// @Description 删除角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param ids body dto.DeleteDTO true "角色ID"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/delete [post]
func (c *UmsRoleController) DeleteRole(g *gin.Context) {

	var req *dto.RoleDeleteDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err := c.svc.UmsRoles().Delete(g, req.Ids)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 查询所有角色列表
// @Description 查询所有角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.UmsRole} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/listAll [get]
func (c *UmsRoleController) ListAllRole(g *gin.Context) {

	umsRoles, err := c.svc.UmsRoles().ListAll(g)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsRoles)
	return
}

// @Summary 更新角色状态
// @Description 更新角色状态
// @Tags 角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.AdminRoleStatusDTO true "角色撞塌更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/updateStatus/:id [post]
func (c *UmsRoleController) UpdateRoleStatus(g *gin.Context) {
	var req *dto.AdminRoleStatusDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	roleId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err = c.svc.UmsRoles().UpdateStatus(g, roleId, req.Status)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 获取角色菜单
// @Description 获取角色菜单
// @Tags 角色
// @Accept json
// @Produce json
// Param id path int true "角色ID"
// @Success 200 {object} response.Response{data=[]models.UmsMenu} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/listMenu/:roleId [get]
func (c *UmsRoleController) ListRoleMenu(g *gin.Context) {

	roleId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	roleMenu, err := c.svc.UmsRoles().ListMenus(g, roleId)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, roleMenu)
	return
}

// @Summary 获取角色资源
// @Description 获取角色资源
// @Tags 角色
// @Accept json
// @Produce json
// Param id path int true "角色ID"
// @Success 200 {object} response.Response{data=[]models.UmsResource} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/listResource/:roleId [get]
func (c *UmsRoleController) ListRoleResource(g *gin.Context) {

	roleId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	roleMenu, err := c.svc.UmsRoles().ListResources(g, roleId)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, roleMenu)
	return
}

// @Summary 给角色分配菜单
// @Description 给角色分配菜单
// @Tags 角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleAllocMenuDTO true "给角色分配菜单"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/allocMenu [post]
func (c *UmsRoleController) AllocRoleMenu(g *gin.Context) {
	var req *dto.RoleAllocMenuDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err := c.svc.UmsRoles().AllocMenu(g, req.RoleId, req.MenuIds)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 给角色分配资源
// @Description 给角色分配资源
// @Tags 角色
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param data body dto.RoleAllocMenuDTO true "给角色分配资源"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /role/allocResource [post]
func (c *UmsRoleController) AllocRoleResource(g *gin.Context) {
	var req *dto.RoleAllocResourceDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err := c.svc.UmsRoles().AllocResource(g, req.RoleId, req.ResourceIds)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}
