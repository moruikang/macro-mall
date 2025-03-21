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

type UmsRoleService interface {
	Create(g *gin.Context, dto *dto.RoleCreateDTO) error
	Get(g *gin.Context, id int64) (*models.UmsRole, error)
	Update(g *gin.Context, id int64, dto *dto.RoleCreateDTO) error
	Delete(g *gin.Context, ids []int64) error
	List(g *gin.Context, dto *dto.PublicPageQuery) (*dto.PageResult, error)
	ListAll(g *gin.Context) ([]*models.UmsRole, error)
	UpdateStatus(g *gin.Context, id int64, status int) error
	// 根据管理员Id获取对应菜单
	GetMenuList(g *gin.Context, adminId int64) ([]*models.UmsMenu, error)
	// 根据角色Id获取对应菜单
	ListMenus(g *gin.Context, roleId int64) ([]*models.UmsMenu, error)
	// 根据角色Id获取对应资源
	ListResources(g *gin.Context, roleId int64) ([]*models.UmsResource, error)
	// 给角色分配菜单
	AllocMenu(g *gin.Context, roleId int64, menuIds []int64) error
	// 给角色分配资源
	AllocResource(g *gin.Context, roleId int64, resourceIds []int64) error
}

type umsRoleService struct {
	store store.Factory
}

var _ UmsRoleService = (*umsRoleService)(nil)

func NewUmsRoleService(store store.Factory) UmsRoleService {
	return &umsRoleService{
		store: store,
	}
}

func (svc *umsRoleService) Create(g *gin.Context, dto *dto.RoleCreateDTO) error {
	role := &models.UmsRole{
		Name:        dto.Name,
		Description: dto.Description,
		AdminCount:  0,
		Sort:        0,
		Status:      dto.Status,
	}
	return svc.store.UmsRoles().Create(g, role)
}

func (svc *umsRoleService) Get(g *gin.Context, id int64) (*models.UmsRole, error) {
	return svc.store.UmsRoles().GetById(g, id)
}

func (svc *umsRoleService) Update(g *gin.Context, id int64, dto *dto.RoleCreateDTO) error {
	role, err := svc.store.UmsRoles().GetById(g, id)
	if err != nil {
		return err
	}
	role.Name = dto.Name
	role.Description = dto.Description
	role.AdminCount = 0
	role.Status = dto.Status
	role.Sort = dto.Sort
	return svc.store.UmsRoles().Update(g, role)
}

func (svc *umsRoleService) Delete(g *gin.Context, ids []int64) error {
	return svc.store.UmsRoles().DeleteCollection(g, ids)
}

func (svc *umsRoleService) List(g *gin.Context, req *dto.PublicPageQuery) (*dto.PageResult, error) {
	total, roles, err := svc.store.UmsRoles().List(g, req.Keyword, req.PageSize, req.PageNum)
	if err != nil {
		return nil, err
	}
	totalPage := int(math.Ceil(float64(total) / float64(req.PageSize)))
	return &dto.PageResult{
		List: roles,
		Pagination: dto.Pagination{
			Total:     total,
			PageNum:   req.PageNum,
			PageSize:  req.PageSize,
			TotalPage: totalPage,
		},
	}, nil
}

func (svc *umsRoleService) ListAll(g *gin.Context) ([]*models.UmsRole, error) {
	return svc.store.UmsRoles().ListAll(g)
}

func (svc *umsRoleService) UpdateStatus(g *gin.Context, id int64, status int) error {

	return svc.store.UmsRoles().UpdateStatus(g, id, status)
}

func (svc *umsRoleService) GetMenuList(g *gin.Context, adminId int64) ([]*models.UmsMenu, error) {
	return svc.store.UmsRoles().GetMenuList(g, adminId)
}

func (svc *umsRoleService) ListMenus(g *gin.Context, roleId int64) ([]*models.UmsMenu, error) {
	return svc.store.UmsRoles().ListMenus(g, roleId)
}

func (svc *umsRoleService) ListResources(g *gin.Context, roleId int64) ([]*models.UmsResource, error) {
	return svc.store.UmsRoles().ListResources(g, roleId)
}

func (svc *umsRoleService) AllocMenu(g *gin.Context, roleId int64, menuIds []int64) error {
	return svc.store.UmsRoles().AllocMenus(g, roleId, menuIds)
}

func (svc *umsRoleService) AllocResource(g *gin.Context, roleId int64, resourceIds []int64) error {
	return svc.store.UmsRoles().AllocResources(g, roleId, resourceIds)
}
