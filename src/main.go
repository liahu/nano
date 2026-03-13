package main

import (
	"fmt"
	"nano/src/api/routes"
	"nano/src/configs"
	"nano/src/db"
	"nano/src/models"
	"nano/src/services"
	"nano/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	if err := configs.LoadConfig("config.toml"); err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		return
	}

	// 初始化日志
	if err := utils.InitLog(); err != nil {
		fmt.Printf("日志初始化失败: %v\n", err)
		return
	}

	// 初始化数据库
	dbInstance, err := db.NewDB("data.db")
	if err != nil {
		fmt.Printf("数据库初始化失败: %v\n", err)
		return
	}

	// 自动迁移表结构
	err = db.AutoMigrate(dbInstance, &models.User{}, &models.Device{})
	if err != nil {
		fmt.Printf("数据库迁移失败: %v\n", err)
		return
	}

	// 初始化服务
	services.InitServices(dbInstance)

	// 初始化Gin
	router := gin.Default()

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务器
	config := configs.GetConfig()
	port := config.Server.Port
	fmt.Printf("服务启动在端口 %d\n", port)
	if err := router.Run(fmt.Sprintf("%s:%d", config.Server.IP, port)); err != nil {
		fmt.Printf("服务启动失败: %v\n", err)
	}
}
