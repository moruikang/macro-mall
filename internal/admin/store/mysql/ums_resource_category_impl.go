package mysql

import (
	"context"
	gorm "gorm.io/gorm"
	"macro-mall/internal/admin/store/models"
)

type umsResourceCategory struct {
	db *gorm.DB
}

func newUmsResourceCategory(ds *datastore) *umsResourceCategory {
	return &umsResourceCategory{ds.db}
}

func (u *umsResourceCategory) Create(ctx context.Context, category *models.UmsResourceCategory) error {
	return u.db.Create(category).Error
}

func (u *umsResourceCategory) Update(ctx context.Context, category *models.UmsResourceCategory) error {
	return u.db.Save(category).Error
}

func (u *umsResourceCategory) DeleteCollection(ctx context.Context, ids []int64) error {
	return u.db.Where("id IN ?", ids).Delete(&models.UmsResourceCategory{}).Error
}

func (u *umsResourceCategory) GetById(ctx context.Context, id int64) (*models.UmsResourceCategory, error) {
	var category models.UmsResourceCategory
	err := u.db.First(&category, id).Error
	return &category, err
}

/*func (u *umsResourceCategory) List(ctx context.Context, keyword string, pageSize, pageNum int) (int64, []*models.UmsResourceCategory, error) {
	query := u.db.Model(&models.UmsResourceCategory{})

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	var totalCount int64
	umsResourceCategories := make([]*models.UmsResourceCategory, 0)

	err := query.Offset(pageSize * (pageNum - 1)).
		Limit(pageSize).
		Order("created_at desc").
		Find(&umsResourceCategories).
		Offset(-1).
		Limit(-1).
		Count(&totalCount).
		Error
	return totalCount, umsResourceCategories, err
}*/

func (u *umsResourceCategory) ListAll(ctx context.Context) ([]*models.UmsResourceCategory, error) {
	var categories []*models.UmsResourceCategory
	err := u.db.Order("created_at desc").Find(&categories).Error
	return categories, err
}
