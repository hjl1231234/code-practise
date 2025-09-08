package database

import (
	"fmt"
	"web-go/blog/config"
	"web-go/blog/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DatabaseHost,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Log.Errorf("数据库连接失败: %v", err)
		return fmt.Errorf("数据库连接失败: %w", err)
	}
	// 添加连接测试
	sqlDB, err := DB.DB()
	if err == nil {
		if pingErr := sqlDB.Ping(); pingErr != nil {
			logger.Log.Errorf("数据库连接测试失败: %v", err)
			return fmt.Errorf("数据库连接测试失败: %w", pingErr)
		}
	}

	logger.Log.Info("数据库连接成功")
	return nil
}
