package routes

import (
	"github.com/gin-gonic/gin"
	"nano/src/api/handlers"
	"nano/src/api/middlewares"
)

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine) {
	// API分组
	api := router.Group("/api/v1")

	// 健康检查
	api.GET("/health", handlers.HealthCheck)

	// App端接口
	app := api.Group("/")
	{
		// 账号管理
		app.POST("/users/register", handlers.Register)
		app.POST("/users/login", handlers.Login)
		
		// 需要认证的接口
		appAuth := app.Group("/")
		appAuth.Use(middlewares.AuthMiddleware())
		{
			// 账号信息
			appAuth.GET("/users/info", handlers.GetUserInfo)
			appAuth.POST("/users/info", handlers.UpdateUserInfo)

			// 设备管理
			appAuth.POST("/devices/bind", handlers.BindDevice)
			appAuth.POST("/devices/unbind", handlers.UnbindDevice)
			appAuth.GET("/devices", handlers.ListDevices)
			appAuth.GET("/devices/:device_id", handlers.GetDeviceInfo)
			appAuth.POST("/devices/:device_id/control", handlers.ControlDevice)
			appAuth.POST("/devices/:device_id", handlers.UpdateDeviceInfo)

			// OTA管理
			appAuth.GET("/ota/latest", handlers.GetLatestOTA)
			appAuth.GET("/devices/:device_id/ota", handlers.GetDeviceOTA)
			appAuth.POST("/devices/:device_id/ota/upgrade", handlers.UpgradeDeviceOTA)
		}
	}

	// 机器人端接口
	device := api.Group("/")
	device.Use(middlewares.DeviceAuthMiddleware())
	{
		// 设备注册
		device.POST("/devices/register", handlers.DeviceRegister)
		
		// 设备状态上报
		device.POST("/devices/status", handlers.ReportDeviceStatus)
		
		// OTA相关
		device.GET("/ota/check", handlers.CheckOTA)
		device.GET("/ota/download/:version", handlers.DownloadOTA)
		device.POST("/devices/ota/status", handlers.ReportOTAStatus)
	}

	// Admin后台接口
	admin := api.Group("/admin")
	admin.Use(middlewares.AdminAuthMiddleware())
	{
		// 用户管理
		admin.GET("/users", handlers.ListUsers)
		admin.POST("/users", handlers.ManageUser)

		// 设备管理
		admin.GET("/devices", handlers.ListAllDevices)
		admin.POST("/devices", handlers.ManageDevice)

		// OTA管理
		admin.POST("/ota/upload", handlers.UploadOTA)
		admin.POST("/ota", handlers.ManageOTA)
		admin.GET("/ota", handlers.ListOTA)
	}
}
