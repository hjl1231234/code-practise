package main

import (
	"fmt"
	"web-go/blog/config"
	"web-go/blog/database"
	"web-go/blog/logger"
	"web-go/blog/middleware"
	"web-go/blog/models"
	"web-go/blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	logger.InitLogger(nil)
	// 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		// logger.Log.Fatalf("加载配置失败: %v", err)
		// 这样会有nil指针问题
		logger.Log.Fatalf("加载配置失败: %v", err)
		fmt.Printf("加载配置失败: %v", err)
		return
	}
	logger.Log.Info("配置加载成功")
	fmt.Println("配置加载成功")

	logger.InitLogger(&cfg)

	database.InitDB(&cfg)
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		logger.Log.Fatalf("数据库迁移失败: %v", err)
	}
	logger.Log.Info("数据库迁移完成")

	router := gin.New()

	router.Use(middleware.RecoveryMiddleware())

	// 全局JWT中间件，也可以在路由组中单独设置
	// router.Use(middleware.JWTUserMiddleware(&cfg))

	routes.SetupRoutes(router, &cfg)

	logger.Log.Infof("服务器启动，监听端口 %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Log.Fatalf("服务启动失败: %v", err)
	}

}
