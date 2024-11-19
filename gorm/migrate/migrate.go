package migrate

import (
	"fmt"
	"gorm/config"
	"gorm/model"
)

func Migrate() {
	// 自动迁移 创建表
	err := config.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	fmt.Println("migrate success")
}
