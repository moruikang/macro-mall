# 一个用 Go 仿写 [macrozheng mall](https://github.com/macrozheng/mall) 的开源项目

# 1. 项目介绍


### 技术选型
- [Go](https://go.dev/) 语言
- [Gin](https://github.com/gin-gonic/gin) 后台框架
- [Gin-jwt](https://github.com/appleboy/gin-jwt) JWT
- [Gorm](https://gorm.io/) ORM
- [MySQL]()
- [Redis](https://github.com/redis/go-redis) 缓存
- [Ladon](https://github.com/ory/ladon) 认证授权



### 技术实现
- mvc 使用了工厂方法模式
```azure
type Service interface {
	UmsAdmins() UmsAdminService
	UmsRoles() UmsRoleService
	UmsMenus() UmsMenuService
	UmsResources() UmsResourceService
	UmsResourceCategorys() UmsResourceCategoryService
}
```
通过service工厂方法创建具体的service实例，service实例通过repository层接口与数据库交互
```azure
type Factory interface {
	UmsAdmins() UmsAdminStore
	UmsRoles() UmsRoleStore
	UmsMenus() UmsMenuStore
	UmsResources() UmsResourceStore
	UmsResourceCategorys() UmsResourceCategoryStore
}
```
如果后续修改存储，只需要重新实现存储层Factory接口即可
- 认证授权使用 [Ladon](https://github.com/ory/ladon) 框架
具体实现授权方案:
```azure
列出所有role
获取每个role - resource 关系
以角色维度来组装DefaultPolicy： 以role id 为ladon subject、 resources url为ladon resource，组装ladon.DefaultPolicy(即一个用户可以访问那些资源)
调用ladon.Manager 创建ladon Policy pool：以role_id:resource_id:resource_url 为redis key，创建一条ladon.DefaultPolicy
缺点: 需要先通过user id获取user role列表，再通过role 列表鉴权
优点: 仅需要维护一个role - resource 关联关系，维护简单
```