// @Author moruikang
// @Date 2025/3/24 21:01:00
// @Desc

package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		adminGroup.Use(authMiddleware.MiddlewareFunc(), middleware.SetJwtInfo())
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

	return engine
}
