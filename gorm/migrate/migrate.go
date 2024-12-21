package migrate

import (
	"fmt"
	"gorm/config"
	"gorm/model"
)

// Migrate 自动迁移
func Migrate() {
	// 自动迁移 创建表
	err := config.DB.AutoMigrate(&model.User{}, &model.Students{})

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	fmt.Println("migrate success")
}
