// @Author moruikang
// @Date 2025/3/16 10:45:00
// @Desc

package mysql

import (
	"context"
	gorm "gorm.io/gorm"
	"macro-mall/internal/admin/store/models"
)

type umsResource struct {
	db *gorm.DB
}

func newUmsResource(ds *datastore) *umsResource {
	return &umsResource{ds.db}
}

func (u umsResource) Create(ctx context.Context, resource *models.UmsResource) error {
	return u.db.Create(resource).Error
}

func (u umsResource) Update(ctx context.Context, resource *models.UmsResource) error {
	return u.db.Save(resource).Error
}

func (u umsResource) DeleteCollection(ctx context.Context, ids []int64) error {
	return u.db.Where("id IN ?", ids).Delete(&models.UmsResource{}).Error
}

func (u umsResource) GetById(ctx context.Context, id int64) (*models.UmsResource, error) {
	var resource models.UmsResource
	err := u.db.First(&resource, id).Error
	return &resource, err
}

func (u umsResource) List(ctx context.Context, categoryId int64, nameKeyword, urlKeyword string, pageSize, pageNum int) (int64, []*models.UmsResource, error) {

	query := u.db.Model(&models.UmsResource{})

	if categoryId != 0 {
		query = query.Where("category_id = ?", categoryId)
	}
	if nameKeyword != "" {
		query = query.Where("name LIKE ?", "%"+nameKeyword+"%")
	}
	if urlKeyword != "" {
		query = query.Where("url LIKE ?", "%"+urlKeyword+"%")
	}
	var totalCount int64
	umsResources := make([]*models.UmsResource, 0)

	err := query.
		Offset(pageSize * (pageNum - 1)).
		Limit(pageSize).
		Order("created_at desc").
		Find(&umsResources).
		Offset(-1).
		Limit(-1).
		Count(&totalCount).
		Error
	return totalCount, umsResources, err
}

func (u umsResource) ListAll(ctx context.Context) ([]*models.UmsResource, error) {
	var resources []*models.UmsResource
	err := u.db.Order("created_at desc").Find(resources).Error
	return resources, err
}

func (u umsResource) ListUmsAdminResource(ctx context.Context, userId int64) ([]*models.UmsResource, error) {
	var resources []*models.UmsResource
	err := u.db.Joins("LEFT JOIN ums_role_resource_relation ON ums_role_resource_relation.resource_id = ums_resource.id").
		Joins("LEFT JOIN ums_admin_role_relation ON ums_admin_role_relation.role_id = ums_role_resource_relation.role_id").
		Where("ums_admin_role_relation.admin_id = ?", userId).
		Find(&resources).
		Error
	return resources, err
}
