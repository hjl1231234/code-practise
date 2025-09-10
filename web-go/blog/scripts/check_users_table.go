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

	// 首先查询users表是否存在
	var tableExists bool
	database.DB.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'users')").Scan(&tableExists)
	fmt.Printf("users表是否存在: %v\n", tableExists)

	// 如果表存在，检查表结构
	if tableExists {
		// 检查所有字段
		fmt.Println("\nusers表所有字段信息:")
		var columns []map[string]interface{}
		database.DB.Raw("SELECT column_name, data_type, is_nullable FROM information_schema.columns WHERE table_name = 'users' ORDER BY ordinal_position").Scan(&columns)
		for _, col := range columns {
			fmt.Printf("字段名: %v, 数据类型: %v, 是否允许NULL: %v\n",
				col["column_name"], col["data_type"], col["is_nullable"])
		}

		// 检查记录数量
		var totalCount int64
		database.DB.Table("users").Count(&totalCount)
		fmt.Printf("\nusers表总记录数: %d\n", totalCount)

		// 如果有记录，显示前几行记录的基本信息
		if totalCount > 0 {
			fmt.Println("\nusers表前3条记录的部分信息:")
			type UserPreview struct {
				ID uint
			}
			var users []UserPreview
			database.DB.Table("users").Select("id").Limit(3).Scan(&users)
			for i, user := range users {
				fmt.Printf("记录 %d: ID=%d\n", i+1, user.ID)
			}
		}

		// 分析错误原因并提供修复建议
		fmt.Println("\n错误分析:")
		fmt.Println("1. 从之前的错误信息来看，迁移过程中尝试添加username字段时失败")
		fmt.Println("2. PostgreSQL在表有数据的情况下添加非空字段时需要默认值")
		fmt.Println("3. 我们之前移除了not null约束是正确的解决方案")

		fmt.Println("\n解决方案:")
		fmt.Println("1. 当前已经移除了models/user.go中username字段的not null约束")
		fmt.Println("2. 现在可以尝试重新运行应用，GORM应该能够成功添加username字段")
		fmt.Println("3. 添加完成后，可以根据需要手动为现有记录设置username值")
	}
}
