package database

import (
	"fmt"
	"time"
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

		// 设置连接池配置
		sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetime) * time.Hour)

		logger.Log.Infof("数据库连接池配置: 最大打开连接数=%d, 最大空闲连接数=%d, 连接最大存活时间=%d小时",
			cfg.DBMaxOpenConns, cfg.DBMaxIdleConns, cfg.DBConnMaxLifetime)
	}

	// 打印所有数据库配置信息
	logger.Log.Infof("数据库配置信息: 主机=%s, 端口=%s, 用户名=%s, 数据库名=%s, SSL模式=disable, 时区=Asia/Shanghai",
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseUser,
		cfg.DatabaseName,
	)

	logger.Log.Info("数据库连接成功")
	return nil
}
