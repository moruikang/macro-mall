// @Author moruikang
// @Date 2025/3/16 09:45:00
// @Desc

package store

import (
	"context"
	"macro-mall/internal/admin/store/models"
	"time"
)

// UmsAdminStore 定义后台用户接口
type UmsAdminStore interface {
	Create(ctx context.Context, admin *models.UmsAdmin) error
	Update(ctx context.Context, admin *models.UmsAdmin) error
	DeleteCollection(ctx context.Context, ids []int64) error
	GetById(ctx context.Context, id int64) (*models.UmsAdmin, error)
	GetByUserName(ctx context.Context, username string) (*models.UmsAdmin, error)
	List(ctx context.Context, keywork string, pageSize, pageNum int) (int64, []*models.UmsAdmin, error)
	ListAll(ctx context.Context) ([]*models.UmsAdmin, error)
	UpdatePassword(ctx context.Context, userId int64, password string) error
	UpdateLoginTime(ctx context.Context, userId int64, time *time.Time) error
	UpdateStatus(ctx context.Context, userId int64, status int) error
	ClearRoles(ctx context.Context, userId int64) error
	UpdateRoles(ctx context.Context, userId int64, list []*models.UmsAdminRoleRelation) error
	GetUserRoles(ctx context.Context, userId int64) ([]*models.UmsRole, error)
}
