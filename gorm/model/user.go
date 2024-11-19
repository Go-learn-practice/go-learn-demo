package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	Age      uint8  `gorm:"check:age>18;not null" json:"age"`
	Email    string `gorm:"size:255" json:"email"`
	Password string `gorm:"size:255" json:"password"`
}
