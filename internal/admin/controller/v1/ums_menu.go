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

type UmsMenuController struct {
	svc v1.Service
}

func NewUmsMenuController(store store.Factory) *UmsMenuController {
	return &UmsMenuController{
		svc: v1.NewService(store),
	}
}

// @Summary 菜单创建
// @Description 菜单创建
// @Tags 菜单
// @Accept json
// @Produce json
// @Param data body dto.MenuCreateDTO true "菜单创建"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/create [post]
func (c *UmsMenuController) Create(g *gin.Context) {

	var umsMenu *dto.MenuCreateDTO
	if err := g.ShouldBindJSON(umsMenu); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	if err := c.svc.UmsMenus().Create(g, umsMenu); err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 获取菜单信息
// @Description 获取菜单信息
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.UmsMenu} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/:id [get]
func (c *UmsMenuController) GetMenu(g *gin.Context) {

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	menu, err := c.svc.UmsMenus().Get(g, id)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, menu)
	return
}

// @Summary 分页查询菜单列表
// @Description 分页查询菜单列表
// @Tags 菜单
// @Accept json
// @Produce json
// @Param data query dto.PublicPageQuery true "分页查询参数"
// @Success 200 {object} response.Response{data=dto.PageResult} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/list/:parentId [get]
func (c *UmsMenuController) ListMenu(g *gin.Context) {

	var query *dto.PageQuery
	if err := g.ShouldBindQuery(query); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	parentId, err := strconv.ParseInt(g.Param("parentId"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	umsMenus, err := c.svc.UmsMenus().List(g, parentId, query)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsMenus)
	return
}

// @Summary 更新菜单信息
// @Description 更新菜单信息
// @Tags 菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param data body dto.MenuCreateDTO true "菜单信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/update/:id [post]
func (c *UmsMenuController) UpdateMenu(g *gin.Context) {
	var umsMenu *dto.MenuCreateDTO
	if err := g.ShouldBindJSON(umsMenu); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	menuId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err = c.svc.UmsMenus().Update(g, menuId, umsMenu)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 删除菜单
// @Description 删除菜单
// @Tags 菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/delete/{id} [post]
func (c *UmsMenuController) DeleteMenu(g *gin.Context) {
	menuId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsMenus().Delete(g, []int64{menuId})
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 更新菜单显示状态
// @Description 更新菜单显示状态
// @Tags 菜单
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param data body dto.MenuUpdateHiddenDTO true "请求体"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/updateHidden/{id} [post]
func (c *UmsMenuController) UpdateMenuHidden(g *gin.Context) {

	var req *dto.MenuUpdateHiddenDTO
	if err := g.ShouldBindJSON(req); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	menuId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsMenus().UpdateHidden(g, menuId, req.Hidden)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 获取菜单树
// @Description 获取菜单树
// @Tags 菜单
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.UmsMenuNode} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /menu/treeList [get]
func (c *UmsMenuController) MenuTreeList(g *gin.Context) {

	menuTree, err := c.svc.UmsMenus().Tree(g)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, menuTree)
	return
}
