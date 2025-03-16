// @Author moruikang
// @Date 2025/3/16 10:45:00
// @Desc

package mysql

import (
	"context"
	gorm "gorm.io/gorm"
	"macro-mall/internal/admin/store/models"
)

type umsMenu struct {
	db *gorm.DB
}

func newUmsMenu(ds *datastore) *umsMenu {
	return &umsMenu{ds.db}
}

func (u *umsMenu) Create(ctx context.Context, menu *models.UmsMenu) error {
	return u.db.Create(menu).Error
}

func (u *umsMenu) Update(ctx context.Context, menu *models.UmsMenu) error {
	return u.db.Save(menu).Error
}

func (u *umsMenu) DeleteCollection(ctx context.Context, ids []int64) error {
	return u.db.Where("id IN ?", ids).Delete(&models.UmsMenu{}).Error
}

func (u *umsMenu) GetById(ctx context.Context, id int64) (*models.UmsMenu, error) {
	var menu models.UmsMenu
	err := u.db.First(&menu, id).Error
	return &menu, err
}

func (u *umsMenu) List(ctx context.Context, parentId int64, pageSize, pageNum int) (int64, []*models.UmsMenu, error) {

	query := u.db.Model(&models.UmsMenu{})

	var totalCount int64
	umsMenus := make([]*models.UmsMenu, 0)

	err := query.Where("parent_id = ?", parentId).
		Offset(pageSize * (pageNum - 1)).
		Limit(pageSize).
		Order("created_at desc").
		Find(&umsMenus).
		Offset(-1).
		Limit(-1).
		Count(&totalCount).
		Error
	return totalCount, umsMenus, err
}

func (u *umsMenu) ListAll(ctx context.Context) ([]*models.UmsMenu, error) {
	var menus []*models.UmsMenu
	err := u.db.Order("created_at desc").Find(&menus).Error
	return menus, err
}

func (u *umsMenu) Tree(ctx context.Context) ([]*models.UmsMenu, error) {
	panic("implement me")
}

func (u *umsMenu) UpdateHidden(ctx context.Context, menuId, hidden int64) error {
	return u.db.Model(&models.UmsMenu{}).Where("id = ?", menuId).Update("hidden", hidden).Error
}
