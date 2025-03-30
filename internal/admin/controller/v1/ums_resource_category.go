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

type UmsResourceCategoryController struct {
	svc v1.Service
}

func NewUmsResourceCategoryController(store store.Factory) *UmsResourceCategoryController {
	return &UmsResourceCategoryController{
		svc: v1.NewService(store),
	}
}

// @Summary 创建资源分类
// @Description 创建资源分类
// @Tags 资源分类
// @Accept json
// @Produce json
// @Param data body dto.ResourceCategoryCreateDTO true "创建资源分类"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resourceCategory/create [post]
func (c *UmsResourceCategoryController) Create(g *gin.Context) {

	var umsResourceCategory *dto.ResourceCategoryCreateDTO
	if err := g.ShouldBindJSON(umsResourceCategory); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	if err := c.svc.UmsResourceCategorys().Create(g, umsResourceCategory); err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 获取资源分类信息
// @Description 获取资源分类信息
// @Tags 资源分类
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.UmsResourceCategory} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resourceCategory/:id [get]
func (c *UmsResourceCategoryController) GetResourceCategory(g *gin.Context) {

	id, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	resourceCategory, err := c.svc.UmsResourceCategorys().Get(g, id)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, resourceCategory)
	return
}

// @Summary 更新资源分类信息
// @Description 更新资源分类信息
// @Tags 资源分类
// @Accept json
// @Produce json
// @Param id path int true "资源分类ID"
// @Param data body dto.ResourceCategoryCreateDTO true "资源分类信息更新"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resourceCategory/update/:id [post]
func (c *UmsResourceCategoryController) UpdateResourceCategory(g *gin.Context) {
	var umsResourceCategory = &dto.ResourceCategoryCreateDTO{}
	if err := g.ShouldBindJSON(umsResourceCategory); err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	resourceCategoryId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}

	err = c.svc.UmsResourceCategorys().Update(g, resourceCategoryId, umsResourceCategory)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 删除资源分类
// @Description 删除资源分类
// @Tags 资源分类
// @Accept json
// @Produce json
// @Param id path int true "资源分类ID"
// @Success 200 {object} response.Response{} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resourceCategory/delete/{id} [post]
func (c *UmsResourceCategoryController) DeleteResourceCategory(g *gin.Context) {
	resourceCategoryId, err := strconv.ParseInt(g.Param("id"), 10, 64)
	if err != nil {
		response.ValidateError(g, err.Error())
		return
	}
	err = c.svc.UmsResourceCategorys().Delete(g, []int64{resourceCategoryId})
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, nil)
	return
}

// @Summary 查询所有资源分类列表
// @Description 查询所有资源分类列表
// @Tags 资源分类
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.UmsResourceCategory} "成功"
// @Failure 400 {object} response.Response{} "失败"
// @Router /resourceCategory/listAll [get]
func (c *UmsResourceCategoryController) ListAllResourceCategory(g *gin.Context) {

	umsResourceCategorys, err := c.svc.UmsResourceCategorys().ListAll(g)
	if err != nil {
		response.Fail(g, err.Error())
		return
	}
	response.Success(g, umsResourceCategorys)
	return
}
