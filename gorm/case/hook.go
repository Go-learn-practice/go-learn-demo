package _case

import (
	"fmt"
	"gorm.io/gorm"
)

// 事务开始前
func (t *Teacher) BeforeSave(tx *gorm.DB) error {
	fmt.Println("Before BeforeSave")
	return nil
}

func (t *Teacher) AfterSave(tx *gorm.DB) error {
	fmt.Println("After AfterSave")
	return nil
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("Before BeforeCreate")
	return nil
}

func (t *Teacher) AfterCreate(tx *gorm.DB) error {
	fmt.Println("After AfterCreate")
	return nil
}

func (t *Teacher) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("Before BeforeUpdate")
	return nil
}

func (t *Teacher) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("After AfterUpdate")
	return nil
}

func (t *Teacher) AfterFind(tx *gorm.DB) error {
	fmt.Println("After AfterFind")
	return nil
}
