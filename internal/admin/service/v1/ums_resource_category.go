// @Author moruikang
// @Date 2025/3/16 17:35:00
// @Desc

package v1

import (
	"github.com/gin-gonic/gin"
	"macro-mall/internal/admin/dto"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/admin/store/models"
)

type UmsResourceCategoryService interface {
	Create(g *gin.Context, dto *dto.ResourceCategoryCreateDTO) error
	Get(g *gin.Context, id int64) (*models.UmsResourceCategory, error)
	Update(g *gin.Context, id int64, dto *dto.ResourceCategoryCreateDTO) error
	Delete(g *gin.Context, ids []int64) error
	ListAll(g *gin.Context) ([]*models.UmsResourceCategory, error)
}

type umsResourceCategoryService struct {
	store store.Factory
}

var _ UmsResourceCategoryService = (*umsResourceCategoryService)(nil)

func NewUmsResourceCategoryService(store store.Factory) UmsResourceCategoryService {
	return &umsResourceCategoryService{
		store: store,
	}
}

func (svc umsResourceCategoryService) Create(g *gin.Context, dto *dto.ResourceCategoryCreateDTO) error {
	resourceCategory := &models.UmsResourceCategory{
		Sort: dto.Sort,
		Name: dto.Name,
	}
	return svc.store.UmsResourceCategorys().Create(g, resourceCategory)
}

func (svc umsResourceCategoryService) Get(g *gin.Context, id int64) (*models.UmsResourceCategory, error) {
	return svc.store.UmsResourceCategorys().GetById(g, id)
}

func (svc umsResourceCategoryService) Update(g *gin.Context, id int64, dto *dto.ResourceCategoryCreateDTO) error {
	resourceCategory, err := svc.store.UmsResourceCategorys().GetById(g, id)
	if err != nil {
		return err
	}
	resourceCategory.Sort = dto.Sort
	resourceCategory.Name = dto.Name
	return svc.store.UmsResourceCategorys().Update(g, resourceCategory)
}

func (svc umsResourceCategoryService) Delete(g *gin.Context, ids []int64) error {
	return svc.store.UmsResourceCategorys().DeleteCollection(g, ids)
}

func (svc umsResourceCategoryService) ListAll(g *gin.Context) ([]*models.UmsResourceCategory, error) {
	return svc.store.UmsResourceCategorys().ListAll(g)
}
