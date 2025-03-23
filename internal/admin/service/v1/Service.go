// @Author moruikang
// @Date 2025/3/16 17:29:00
// @Desc

package v1

import "macro-mall/internal/admin/store"

type Service interface {
	UmsAdmins() UmsAdminService
	UmsRoles() UmsRoleService
	UmsMenus() UmsMenuService
	UmsResources() UmsResourceService
	UmsResourceCategorys() UmsResourceCategoryService
}

type service struct {
	store store.Factory
}

var _ Service = (*service)(nil)

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) UmsAdmins() UmsAdminService {
	return NewUmsAdminService(s.store)
}

func (s *service) UmsRoles() UmsRoleService {
	return NewUmsRoleService(s.store)
}

func (s *service) UmsMenus() UmsMenuService {
	return NewUmsMenuService(s.store)
}

func (s *service) UmsResources() UmsResourceService {
	return NewUmsResourceService(s.store)
}

func (s *service) UmsResourceCategorys() UmsResourceCategoryService {
	return NewUmsResourceCategoryService(s.store)
}
