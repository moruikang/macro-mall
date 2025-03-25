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

type UmsResourceController struct {
	svc v1.Service
}

func NewUmsResourceController(store store.Factory) *UmsResourceController {
	return &UmsResourceController{
		svc: v1.NewService(store),
	}
}

// @Summary 创建资源
// @Description 创建资源
// @Tags 资源
// @Accept json
// @Produce json
// @Param data body dto.ResourceCreateDTO true "创建资源"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/create [post]
func (c *UmsResourceController) Create(g *gin.Context) {

	var umsResource *dto.ResourceCreateDTO
	if err := g.ShouldBindJSON(umsResource); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	if err := c.svc.UmsResources().Create(g, umsResource); err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 获取资源信息
// @Description 获取资源信息
// @Tags 资源
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.UmsResource} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/:id [get]
func (c *UmsResourceController) GetResource(g *gin.Context) {

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	resource, err := c.svc.UmsResources().Get(g, id)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, resource)
	return
}

// @Summary 分页查询资源列表
// @Description 分页查询资源列表
// @Tags 资源
// @Accept json
// @Produce json
// @Param data param dto.ResourcePageQueryDTO true "分页查询参数"
// @Success 200 {object} response.Response{data=dto.PageResult} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/list [get]
func (c *UmsResourceController) ListResource(g *gin.Context) {

	var query *dto.ResourcePageQueryDTO
	if err := g.ShouldBindQuery(query); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	umsResources, err := c.svc.UmsResources().List(g, query)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsResources)
	return
}

// @Summary 更新资源信息
// @Description 更新资源信息
// @Tags 资源
// @Accept json
// @Produce json
// @Param id path int true "资源ID"
// @Param data body dto.ResourceCreateDTO true "资源信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/update/:id [post]
func (c *UmsResourceController) UpdateResource(g *gin.Context) {
	var umsResource *dto.ResourceCreateDTO
	if err := g.ShouldBindJSON(umsResource); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	resourceId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err = c.svc.UmsResources().Update(g, resourceId, umsResource)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 删除资源
// @Description 删除资源
// @Tags 资源
// @Accept json
// @Produce json
// @Param id path int true "资源ID"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/delete/{id} [post]
func (c *UmsResourceController) DeleteResource(g *gin.Context) {
	resourceId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsResources().Delete(g, []int64{resourceId})
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 查询所有资源列表
// @Description 查询所有资源列表
// @Tags 资源
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.UmsResource} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resource/listAll [get]
func (c *UmsResourceController) ListAllResource(g *gin.Context) {

	umsResources, err := c.svc.UmsResources().ListAll(g)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsResources)
	return
}
