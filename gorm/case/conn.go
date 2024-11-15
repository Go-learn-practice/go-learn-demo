package _case

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB
var dsn = "root:123456@tcp(127.0.0.1)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// 预编译：不支持嵌套事务
		PrepareStmt: true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	setPool(DB)
}

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(10)
}
