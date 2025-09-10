package main

import (
	"fmt"
	"web-go/blog/config"
	"web-go/blog/database"
	"web-go/blog/logger"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}

	logger.InitLogger(&cfg)

	// 初始化数据库连接
	database.InitDB(&cfg)

	// 查询users表结构
	fmt.Println("users表完整结构信息:")
	var columns []map[string]interface{}
	database.DB.Raw(`
		SELECT column_name, data_type, character_maximum_length, is_nullable, column_default
		FROM information_schema.columns 
		WHERE table_name = 'users' 
		ORDER BY ordinal_position
	`).Scan(&columns)

	for _, col := range columns {
		fmt.Printf("字段名: %-15v 数据类型: %-20v 长度: %-5v 允许NULL: %-5v 默认值: %v\n",
			col["column_name"],
			col["data_type"],
			col["character_maximum_length"],
			col["is_nullable"],
			col["column_default"])
	}

	// 查询是否有NULL值记录
	fmt.Println("\n检查NULL值记录:")
	for _, col := range columns {
		colName := col["column_name"].(string)
		var nullCount int64
		database.DB.Table("users").Where(fmt.Sprintf("%s IS NULL", colName)).Count(&nullCount)
		if nullCount > 0 {
			fmt.Printf("字段 '%s' 中有 %d 条NULL值记录\n", colName, nullCount)
		}
	}

	// 提供修复建议
	fmt.Println("\n修复建议:")
	fmt.Println("1. 根据实际表结构修改models/user.go文件，使用gorm的column标签映射字段")
	fmt.Println("2. 对于包含NULL值的字段，不要在模型中设置not null约束")
	fmt.Println("3. 如果需要添加约束，先为NULL值记录设置默认值")
}