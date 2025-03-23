// @Author moruikang
// @Date 2025/3/16 09:40:00
// @Desc

package store

var client Factory

// 工厂定义 存储接口
type Factory interface {
	UmsAdmins() UmsAdminStore
	UmsRoles() UmsRoleStore
	UmsMenus() UmsMenuStore
	UmsResources() UmsResourceStore
	UmsResourceCategorys() UmsResourceCategoryStore
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
