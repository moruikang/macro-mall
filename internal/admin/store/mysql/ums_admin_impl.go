// @Author moruikang
// @Date 2025/3/16 10:45:00
// @Desc

package mysql

import (
	"context"
	gorm "gorm.io/gorm"
	"macro-mall/internal/admin/store/models"
	"time"
)

type umsAdmin struct {
	db *gorm.DB
}

func newUmsAdmin(ds *datastore) *umsAdmin {
	return &umsAdmin{ds.db}
}

func (u *umsAdmin) Create(ctx context.Context, user *models.UmsAdmin) error {
	return u.db.Create(user).Error
}

func (u *umsAdmin) Update(ctx context.Context, admin *models.UmsAdmin) error {
	return u.db.Save(admin).Error
}

func (u *umsAdmin) DeleteCollection(ctx context.Context, ids []int64) error {
	return u.db.Where("id IN ?", ids).Delete(&models.UmsAdmin{}).Error
}

func (u *umsAdmin) GetById(ctx context.Context, id int64) (*models.UmsAdmin, error) {
	var admin models.UmsAdmin
	err := u.db.First(&admin, id).Error
	return &admin, err
}

func (u *umsAdmin) GetByUserName(ctx context.Context, username string) (*models.UmsAdmin, error) {
	var admin models.UmsAdmin
	err := u.db.Where("username = ? and status = 1", username).First(&admin).Error
	return &admin, err
}

func (u *umsAdmin) List(ctx context.Context, keyword string, pageSize, pageNum int) (int64, []*models.UmsAdmin, error) {

	query := u.db.Model(&models.UmsAdmin{})

	if keyword != "" {
		query = query.Where("username LIKE ? or nick_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	var totalCount int64
	umsAdmins := make([]*models.UmsAdmin, 0)

	err := query.Offset(pageSize * (pageNum - 1)).
		Limit(pageSize).
		Order("created_at desc").
		Find(&umsAdmins).
		Offset(-1).
		Limit(-1).
		Count(&totalCount).
		Error
	return totalCount, umsAdmins, err
}

func (u *umsAdmin) ListAll(ctx context.Context) ([]*models.UmsAdmin, error) {
	var admins []*models.UmsAdmin
	err := u.db.Find(&admins).Error
	return admins, err
}

func (u *umsAdmin) UpdatePassword(ctx context.Context, userId int64, password string) error {
	return u.db.Model(&models.UmsAdmin{}).Where("id = ?", userId).Update("password", password).Error
}

func (u *umsAdmin) UpdateLoginTime(ctx context.Context, userId int64, time *time.Time) error {
	return u.db.Model(&models.UmsAdmin{}).Where("id = ?", userId).Update("login_time", time).Error
}

func (u *umsAdmin) UpdateStatus(ctx context.Context, userId int64, status int) error {
	return u.db.Model(&models.UmsAdmin{}).Where("id = ?", userId).Update("status", status).Error
}

func (u *umsAdmin) ClearRoles(ctx context.Context, userId int64) error {
	return u.db.Where("admin_id = ?", userId).Delete(&models.UmsAdminRoleRelation{}).Error
}

func (u *umsAdmin) UpdateRoles(ctx context.Context, userId int64, list []*models.UmsAdminRoleRelation) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("admin_id = ?", userId).Delete(&models.UmsAdminRoleRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(list).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *umsAdmin) GetUserRoles(ctx context.Context, userId int64) ([]*models.UmsRole, error) {
	var roles []*models.UmsRole
	err := u.db.Joins("LEFT JOIN ums_admin_role_relation ON ums_admin_role_relation.role_id = ums_role.id").
		Where("ums_admin_role_relation.admin_id = ?", userId).Find(&roles).Error
	return roles, err
}
