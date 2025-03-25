// @Author moruikang
// @Date 2025/3/24 21:01:00
// @Desc

package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "macro-mall/doc/mall_admin"
	controller "macro-mall/internal/admin/controller/v1"
	"macro-mall/internal/admin/store/mysql"
	"macro-mall/internal/pkg/middleware"
)

func Router() *gin.Engine {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())
	// TODO
	//engine.Use(middleware.Recovery())
	storeIns, err := mysql.GetMySQLFactoryOr(nil)
	if err != nil {
		log.Fatal("failed to get mysql store factory, error: " + err.Error())
	}
	authMiddleware, err := middleware.NewGinJwtMiddleware()
	if err != nil {
		log.Fatal("failed to get gin jwt middleware, error: " + err.Error())
	}
	// 后台管理员接口
	adminGroup := engine.Group("/admin")
	{
		adminController := controller.NewUmsAdminController(storeIns)
		noAuthGroup := adminGroup.Group("")
		{
			noAuthGroup.POST("/register", adminController.Register)
			noAuthGroup.POST("/login", authMiddleware.LoginHandler)
		}
		//adminGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
		adminGroup.Use(authMiddleware.MiddlewareFunc())
		{
			adminGroup.POST("/logout", authMiddleware.LogoutHandler)
			adminGroup.POST("/refreshToken", authMiddleware.RefreshHandler)

			adminGroup.GET("/info", adminController.GetAdminInfo)
			adminGroup.GET("/list", adminController.ListAdmin)
			adminGroup.GET("/:id", adminController.GetAdmin)
			adminGroup.POST("/update/:id", adminController.UpdateAdmin)
			adminGroup.POST("/:id", adminController.DeleteAdmin)
			adminGroup.PUT("/updateStatus/:id", adminController.UpdateStatus)
			adminGroup.PUT("/updatePassword", adminController.UpdatePassword)
			adminGroup.PUT("/role/update", adminController.UpdateAdminRole)
		}
	}

	// 后台角色接口
	roleGroup := adminGroup.Group("/role")
	roleGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
	{
		roleController := controller.NewUmsRoleController(storeIns)
		roleGroup.POST("/create", roleController.Create)
		roleGroup.GET("/list", roleController.ListRole)
		roleGroup.GET("/listAll", roleController.ListAllRole)
		roleGroup.POST("/update/:id", roleController.UpdateRole)
		roleGroup.POST("/delete", roleController.DeleteRole)
		roleGroup.POST("/updateStatus/:id", roleController.UpdateRoleStatus)
		roleGroup.GET("/listMenu/:roleId", roleController.ListRoleMenu)
		roleGroup.GET("/listResource/:roleId", roleController.ListRoleResource)
		roleGroup.POST("/allocMenu", roleController.AllocRoleMenu)
		roleGroup.POST("/allocResource", roleController.AllocRoleResource)

	}
	// 后台菜单接口
	menuGroup := adminGroup.Group("/menu")
	menuGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
	{
		menuController := controller.NewUmsMenuController(storeIns)
		menuGroup.POST("/create", menuController.Create)
		menuGroup.GET("/list/:parentId", menuController.ListMenu)
		menuGroup.POST("/update/:id", menuController.UpdateMenu)
		menuGroup.POST("/delete/:id", menuController.DeleteMenu)
		menuGroup.POST("/updateHidden/:id", menuController.UpdateMenuHidden)
		menuGroup.GET("/treeList", menuController.MenuTreeList)
		menuGroup.GET("/:id", menuController.GetMenu)
	}
	// 后台资源接口
	resourceGroup := adminGroup.Group("/resource")
	resourceGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
	{
		resourceController := controller.NewUmsResourceController(storeIns)
		resourceGroup.POST("/create", resourceController.Create)
		resourceGroup.GET("/list", resourceController.ListResource)
		resourceGroup.GET("/:id", resourceController.GetResource)
		resourceGroup.POST("/update/:id", resourceController.UpdateResource)
		resourceGroup.POST("/delete/:id", resourceController.DeleteResource)
		resourceGroup.GET("/listAll", resourceController.ListAllResource)
	}
	// 后台资源分类接口
	resourceCategoryGroup := adminGroup.Group("/resourceCategory")
	resourceCategoryGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
	{
		resourceCategoryController := controller.NewUmsResourceCategoryController(storeIns)
		resourceCategoryGroup.POST("/create", resourceCategoryController.Create)
		resourceCategoryGroup.GET("/listAll", resourceCategoryController.ListAllResourceCategory)
		resourceCategoryGroup.GET("/:id", resourceCategoryController.GetResourceCategory)
		resourceCategoryGroup.POST("/update/:id", resourceCategoryController.UpdateResourceCategory)
		resourceCategoryGroup.POST("/delete/:id", resourceCategoryController.DeleteResourceCategory)
	}

	docs.SwaggerInfo.Title = "mall_admin"
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return engine
}
