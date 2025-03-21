// @Author moruikang
// @Date 2025/3/16 09:45:00
// @Desc

package store

import (
	"context"
	"macro-mall/internal/admin/store/models"
)

// UmsRoleStore 定义后台用户角色接口
type UmsRoleStore interface {
	Create(ctx context.Context, menu *models.UmsRole) error
	Update(ctx context.Context, menu *models.UmsRole) error
	DeleteCollection(ctx context.Context, ids []int64) error
	GetById(ctx context.Context, id int64) (*models.UmsRole, error)
	List(ctx context.Context, keywork string, pageSize, pageNum int) (int64, []*models.UmsRole, error)
	ListAll(ctx context.Context) ([]*models.UmsRole, error)
	UpdateStatus(ctx context.Context, id int64, status int) error
	// 根据管理员ID获取用户菜单列表
	GetMenuList(ctx context.Context, adminId int64) ([]*models.UmsMenu, error)
	// 根据云角色ID获取菜单列表
	ListMenus(ctx context.Context, roleId int64) ([]*models.UmsMenu, error)
	// 根据角色Id获取角色资源列表
	ListResources(ctx context.Context, roleId int64) ([]*models.UmsResource, error)
	// 根据角色Id获取角色和资源关联关系
	ListRoleResourceRelations(ctx context.Context, roleId int64) ([]*models.UmsRoleResourceRelation, error)
	// 给角色分配菜单
	AllocMenus(ctx context.Context, roleId int64, menuIds []int64) error
	// 给角色分配资源
	AllocResources(ctx context.Context, roleId int64, resourceIds []int64) error
}
