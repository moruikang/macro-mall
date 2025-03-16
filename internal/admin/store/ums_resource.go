// @Author moruikang
// @Date 2025/3/16 09:45:00
// @Desc

package store

import (
	"context"
	"macro-mall/internal/admin/store/models"
)

// UmsResourceStore 定义后台用户资源接口
type UmsResourceStore interface {
	Create(ctx context.Context, menu *models.UmsResource) error
	Update(ctx context.Context, menu *models.UmsResource) error
	DeleteCollection(ctx context.Context, ids []int64) error
	GetById(ctx context.Context, id int64) (*models.UmsResource, error)
	List(ctx context.Context, categoryId int64, nameKeywork, urlKeyword string, pageSize, pageNum int) (int64, []*models.UmsResource, error)
	ListAll(ctx context.Context) ([]*models.UmsResource, error)
	ListUmsAdminResource(ctx context.Context, userId int64) ([]*models.UmsResource, error)
}
