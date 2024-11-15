package _case

import (
	"fmt"
	"time"
)

func Crud() {
	temp := Teacher{
		Name:     "nick",
		Age:      40,
		Role:     []string{"普通用户", "会员"},
		Birthday: time.Now().Unix(),
		Salary:   12345.123,
		Email:    "nick@gmail.com",
	}
	t := temp
	res := DB.Create(&t)
	fmt.Println(res.RowsAffected, res.Error, t.ID)

	t1 := Teacher{}
	DB.First(&t1)
	fmt.Println(t1)

	t1.Name = "king"
	t1.Age = 31
	DB.Save(&t1)

	DB.Delete(&t1)
}
