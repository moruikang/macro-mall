// @Author moruikang
// @Date 2025/3/16 09:45:00
// @Desc

package store

import (
	"context"
	"macro-mall/internal/admin/store/models"
)

// UmsMenuStore 定义后台用户菜单接口
type UmsMenuStore interface {
	Create(ctx context.Context, menu *models.UmsMenu) error
	Update(ctx context.Context, menu *models.UmsMenu) error
	DeleteCollection(ctx context.Context, ids []int64) error
	GetById(ctx context.Context, id int64) (*models.UmsMenu, error)
	List(ctx context.Context, parentId int64, pageSize, pageNum int) (int64, []*models.UmsMenu, error)
	ListAll(ctx context.Context) ([]*models.UmsMenu, error)

	// 获取菜单树形结构
	Tree(ctx context.Context) ([]*models.UmsMenu, error)
	// 修改菜单显示状态
	UpdateHidden(ctx context.Context, menuId, hidden int64) error
}
