// @Author moruikang
// @Date 2025/3/16 10:45:00
// @Desc

package mysql

import (
	"context"
	gorm "gorm.io/gorm"
	"macro-mall/internal/admin/store/models"
)

type umsRole struct {
	db *gorm.DB
}

func newUmsRole(ds *datastore) *umsRole {
	return &umsRole{ds.db}
}

func (u umsRole) Create(ctx context.Context, role *models.UmsRole) error {
	return u.db.Create(role).Error
}

func (u umsRole) Update(ctx context.Context, role *models.UmsRole) error {
	return u.db.Save(role).Error
}

func (u umsRole) DeleteCollection(ctx context.Context, ids []int64) error {
	return u.db.Where("id IN ?", ids).Delete(&models.UmsRole{}).Error
}

func (u umsRole) GetById(ctx context.Context, id int64) (*models.UmsRole, error) {
	var role models.UmsRole
	err := u.db.First(&role, id).Error
	return &role, err
}

func (u umsRole) List(ctx context.Context, keyword string, pageSize, pageNum int) (int64, []*models.UmsRole, error) {
	query := u.db.Model(&models.UmsRole{})

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	var totalCount int64
	umsRoles := make([]*models.UmsRole, 0)

	err := query.Offset(pageSize * (pageNum - 1)).
		Limit(pageSize).
		Order("created_at desc").
		Find(&umsRoles).
		Offset(-1).
		Limit(-1).
		Count(&totalCount).
		Error
	return totalCount, umsRoles, err
}

func (u umsRole) ListAll(ctx context.Context) ([]*models.UmsRole, error) {
	var roles []*models.UmsRole
	err := u.db.Order("created_at desc").Find(&roles).Error
	return roles, err
}

func (u umsRole) UpdateStatus(ctx context.Context, id int64, status int) error {
	return u.db.Model(&models.UmsRole{}).Where("id = ?", id).Update("status", status).Error
}

func (u umsRole) GetMenuList(ctx context.Context, adminId int64) ([]*models.UmsMenu, error) {
	var menus []*models.UmsMenu
	/*	err := u.db.Joins("LEFT JOIN ums_role_menu_relation ON ums_role_menu_relation.menu_id = ums_menu.id").
		Joins("LEFT JOIN ums_admin_role_relation ON ums_admin_role_relation.role_id = ums_role_menu_relation.role_id").
		Where("ums_admin_role_relation.admin_id = ?", adminId).
		Find(&menus).Error*/
	err := u.db.
		Joins("LEFT JOIN ums_role r ON arr.role_id = r.id").
		Joins("LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id").
		Joins("LEFT JOIN ums_menu m ON rmr.menu_id = m.id").
		Where("arr.admin_id = ? AND m.id IS NOT NULL", adminId).
		Group("m.id").
		Find(&menus).Error
	return menus, err
}

func (u umsRole) ListMenus(ctx context.Context, roleId int64) ([]*models.UmsMenu, error) {
	var menus []*models.UmsMenu
	err := u.db.Joins("LEFT JOIN ums_role_menu_relation ON ums_role_menu_relation.menu_id = ums_menu.id").
		Where("ums_role_menu_relation.role_id = ?", roleId).Find(&menus).Error
	return menus, err
}

func (u umsRole) ListResources(ctx context.Context, roleId int64) ([]*models.UmsResource, error) {
	var resources []*models.UmsResource
	err := u.db.Joins("LEFT JOIN ums_role_resource_relation ON ums_role_resource_relation.resource_id = ums_resource.id").
		Where("ums_role_resource_relation.role_id = ?", roleId).Find(&resources).Error
	return resources, err
}

func (u umsRole) ListRoleResourceRelations(ctx context.Context, roleId int64) ([]*models.UmsRoleResourceRelation, error) {
	var relations []*models.UmsRoleResourceRelation
	err := u.db.Where("role_id = ?", roleId).Find(&relations).Error
	return relations, err
}

func (u umsRole) AllocMenus(ctx context.Context, roleId int64, menuIds []int64) error {
	/*	tx := u.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Where("role_id = ?", roleId).Delete(&models.UmsRoleMenuRelation{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		var relations []*models.UmsRoleMenuRelation
		for _, menuId := range menuIds {
			relations = append(relations, &models.UmsRoleMenuRelation{
				RoleID: roleId,
				MenuID: menuId,
			})
		}

		if err := tx.Create(relations).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error*/
	return u.db.Transaction(func(tx *gorm.DB) error {
		// 删除原有角色菜单关联
		if err := tx.Where("role_id = ?", roleId).Delete(&models.UmsRoleMenuRelation{}).Error; err != nil {
			return err
		}
		// 添加新的角色菜单关联
		roleMenuRelations := make([]*models.UmsRoleMenuRelation, 0)
		for _, menuId := range menuIds {
			roleMenuRelations = append(roleMenuRelations, &models.UmsRoleMenuRelation{
				RoleID: roleId,
				MenuID: menuId,
			})
		}

		return u.db.Create(roleMenuRelations).Error
	})
}

func (u umsRole) AllocResources(ctx context.Context, roleId int64, resourceIds []int64) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		// 删除原有角色资源关联
		if err := tx.Where("role_id = ?", roleId).Delete(&models.UmsRoleResourceRelation{}).Error; err != nil {
			return err
		}
		// 添加新的角色资源关联
		roleMenuRelations := make([]*models.UmsRoleResourceRelation, 0)
		for _, resourceId := range resourceIds {
			roleMenuRelations = append(roleMenuRelations, &models.UmsRoleResourceRelation{
				RoleID:     roleId,
				ResourceID: resourceId,
			})
		}

		return u.db.Create(roleMenuRelations).Error
	})
}
