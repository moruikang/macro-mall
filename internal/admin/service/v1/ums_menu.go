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

type UmsMenuService interface {
	Create(g *gin.Context, dto *dto.MenuCreateDTO) error
	Get(g *gin.Context, id int64) (*models.UmsMenu, error)
	Update(g *gin.Context, id int64, dto *dto.MenuCreateDTO) error
	Delete(g *gin.Context, ids []int64) error
	List(g *gin.Context, parentId int64, dto *dto.PageQuery) (*dto.PageResult, error)
	// 树形结构返回所有菜单列表
	Tree(ctx *gin.Context) ([]*models.UmsMenuNode, error)
	// 修改菜单显示状态
	UpdateHidden(ctx *gin.Context, menuId, hidden int64) error
}

type umsMenuService struct {
	store store.Factory
}

var _ UmsMenuService = (*umsMenuService)(nil)

func NewUmsMenuService(store store.Factory) UmsMenuService {
	return &umsMenuService{
		store: store,
	}
}

func (svc *umsMenuService) Create(g *gin.Context, dto *dto.MenuCreateDTO) error {

	menu := &models.UmsMenu{
		ParentId: dto.ParentId,
		Title:    dto.Title,
		Level:    dto.Level,
		Sort:     dto.Sort,
		Hidden:   dto.Hidden,
	}
	return svc.store.UmsMenus().Create(g, menu)
}

func (svc *umsMenuService) Get(g *gin.Context, id int64) (*models.UmsMenu, error) {

	return svc.store.UmsMenus().GetById(g, id)
}

func (svc *umsMenuService) Update(g *gin.Context, id int64, dto *dto.MenuCreateDTO) error {

	menu, err := svc.store.UmsMenus().GetById(g, id)
	if err != nil {
		return err
	}
	menu.ParentId = dto.ParentId
	menu.Title = dto.Title
	menu.Level = dto.Level
	menu.Sort = dto.Sort
	menu.Name = dto.Name
	menu.Icon = dto.Icon
	menu.Hidden = dto.Hidden
	return svc.store.UmsMenus().Update(g, menu)
}

func (svc *umsMenuService) Delete(g *gin.Context, ids []int64) error {

	return svc.store.UmsMenus().DeleteCollection(g, ids)
}

func (svc *umsMenuService) List(g *gin.Context, parentId int64, req *dto.PageQuery) (*dto.PageResult, error) {

	total, menus, err := svc.store.UmsMenus().List(g, parentId, req.PageSize, req.PageNum)
	if err != nil {
		return nil, err
	}
	totalPage := int(math.Ceil(float64(total) / float64(req.PageSize)))
	return &dto.PageResult{
		List: menus,
		Pagination: dto.Pagination{
			Total:     total,
			PageNum:   req.PageNum,
			PageSize:  req.PageSize,
			TotalPage: totalPage,
		},
	}, nil
}

func (svc *umsMenuService) Tree(ctx *gin.Context) ([]*models.UmsMenuNode, error) {

	allMenus, err := svc.store.UmsMenus().ListAll(ctx)
	if err != nil {
		return nil, err
	}
	menuNodes := make([]*models.UmsMenuNode, 0)
	for _, menu := range allMenus {
		if menu.ParentId == 0 {
			menuNodes = append(menuNodes, converMenuNode(menu, allMenus))
		}
	}
	return menuNodes, nil

}

func (svc *umsMenuService) UpdateHidden(ctx *gin.Context, menuId, hidden int64) error {

	return svc.store.UmsMenus().UpdateHidden(ctx, menuId, hidden)
}

func converMenuNode(menu *models.UmsMenu, menus []*models.UmsMenu) *models.UmsMenuNode {

	var umsMenuNode = &models.UmsMenuNode{}
	var umsMenuNodeChildren []*models.UmsMenuNode
	umsMenuNode.UmsMenu = *menu
	for _, umsMenu := range menus {
		if umsMenu.ParentId == menu.Id {
			umsMenuNodeChildren = append(umsMenuNodeChildren, converMenuNode(umsMenu, menus))
		}
	}
	umsMenuNode.Children = umsMenuNodeChildren
	return umsMenuNode
}
