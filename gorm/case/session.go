package _case

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Session() {
	tx := DB.Session(&gorm.Session{
		PrepareStmt:              true,
		SkipHooks:                true,
		DisableNestedTransaction: true,
		Logger:                   DB.Logger.LogMode(logger.Error),
	})
	t := Teacher{
		Name:     "nick",
		Age:      40,
		Role:     []string{"普通用户", "会员"},
		Birthday: time.Now().Unix(),
		Salary:   12345.123,
		Email:    "nick@gmail.com",
	}
	tx.Create(&t)
}
