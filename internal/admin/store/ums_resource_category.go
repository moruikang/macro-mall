// @Author moruikang
// @Date 2025/3/16 09:45:00
// @Desc

package store

import (
	"context"
	"macro-mall/internal/admin/store/models"
)

// UmsResourceCategoryStore 定义后台用户资源-分类接口
type UmsResourceCategoryStore interface {
	Create(ctx context.Context, menu *models.UmsResourceCategory) error
	Update(ctx context.Context, menu *models.UmsResourceCategory) error
	DeleteCollection(ctx context.Context, ids []int64) error
	GetById(ctx context.Context, id int64) (*models.UmsResourceCategory, error)
	//List(ctx context.Context, keyword string, pageSize, pageNum int) (int64, []*models.UmsResourceCategory, error)
	ListAll(ctx context.Context) ([]*models.UmsResourceCategory, error)
}
