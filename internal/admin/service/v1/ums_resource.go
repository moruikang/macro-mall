// @Author moruikang
// @Date 2025/3/16 17:35:00
// @Desc

package v1

import (
	"github.com/gin-gonic/gin"
	"macro-mall/internal/admin/dto"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/admin/store/models"
	"math"
)

type UmsResourceService interface {
	Create(g *gin.Context, dto *dto.ResourceCreateDTO) error
	Get(g *gin.Context, id int64) (*models.UmsResource, error)
	Update(g *gin.Context, id int64, dto *dto.ResourceCreateDTO) error
	Delete(g *gin.Context, ids []int64) error
	List(g *gin.Context, parentId int64, dto *dto.ResourcePageQueryDTO) (*dto.PageResult, error)
	ListAll(g *gin.Context) ([]*models.UmsResource, error)
}

type umsResourceService struct {
	store store.Factory
}

var _ UmsResourceService = (*umsResourceService)(nil)

func NewUmsResourceService(store store.Factory) UmsResourceService {
	return &umsResourceService{
		store: store,
	}
}

func (u umsResourceService) Create(g *gin.Context, dto *dto.ResourceCreateDTO) error {
	resource := &models.UmsResource{
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Url:         dto.Url,
		Description: dto.Description,
	}
	return u.store.UmsResources().Create(g, resource)
}

func (u umsResourceService) Get(g *gin.Context, id int64) (*models.UmsResource, error) {
	return u.store.UmsResources().GetById(g, id)
}

func (u umsResourceService) Update(g *gin.Context, id int64, dto *dto.ResourceCreateDTO) error {
	resource, err := u.store.UmsResources().GetById(g, id)
	if err != nil {
		return err
	}
	resource.CategoryId = dto.CategoryId
	resource.Name = dto.Name
	resource.Url = dto.Url
	resource.Description = dto.Description
	return u.store.UmsResources().Update(g, resource)
}

func (u umsResourceService) Delete(g *gin.Context, ids []int64) error {
	return u.store.UmsResources().DeleteCollection(g, ids)
}

func (u umsResourceService) List(g *gin.Context, parentId int64, queryDTO *dto.ResourcePageQueryDTO) (*dto.PageResult, error) {
	total, resources, err := u.store.UmsResources().List(g,
		parentId,
		queryDTO.NameKeyword,
		queryDTO.UrlKeyword,
		queryDTO.PageSize,
		queryDTO.PageNum,
	)
	if err != nil {
		return nil, err
	}
	totalPage := int(math.Ceil(float64(total) / float64(queryDTO.PageSize)))
	return &dto.PageResult{
		List: resources,
		Pagination: dto.Pagination{
			Total:     total,
			PageNum:   queryDTO.PageNum,
			PageSize:  queryDTO.PageSize,
			TotalPage: totalPage,
		},
	}, nil
}

func (u umsResourceService) ListAll(g *gin.Context) ([]*models.UmsResource, error) {
	return u.store.UmsResources().ListAll(g)
}
