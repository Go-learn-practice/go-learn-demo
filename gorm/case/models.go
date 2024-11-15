package _case

import "gorm.io/gorm"

func init() {
	err := DB.Migrator().AutoMigrate(&Teacher{})
	if err != nil {
		return
	}
}

type Roles []string
type Teacher struct {
	gorm.Model
	Name     string  `gorm:"size:256"`
	Age      uint8   `gorm:"check:age>30"`
	Email    string  `gorm:"size:256"`
	Salary   float64 `gorm:"scale:2;precision:7"`
	Birthday int64   `gorm:"serializer:unixtime;type:time"`
	Role     Roles   `gorm:"serializer:json"`
}
