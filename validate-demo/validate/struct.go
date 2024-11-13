package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// 结构绑定格式验证

type User struct {
	Name           string         `v:"required,alphaunicode"`
	Age            uint8          `v:"gte=10,lte=30"`
	Phone          string         `v:"required,e164"`
	Email          string         `v:"required,email"`
	FavoriteColor1 string         `v:"iscolor"`
	FavoriteColor2 string         `v:"hexcolor|rgb|rgba|hsl|hsla"`
	Address        *Address       `v:"required"`
	ContactUser    []*ContactUser `v:"required,gte=1,dive"`
	Hobby          []string       `v:"required,gte=2,dive,required,gte=2,alphaunicode"`
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=20,lte=30"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Phone Email"`
}

func StructValidate() {
	v := validate
	address := &Address{
		Province: "广东省",
		City:     "广州市",
	}

	contactUser1 := &ContactUser{
		Name:    "张三",
		Age:     30,
		Phone:   "+8613800138000",
		Email:   "nick@gmail.com",
		Address: address,
	}

	contactUser2 := &ContactUser{
		Name:    "李四",
		Age:     20,
		Phone:   "+8613800138000",
		Email:   "nick@gmail.com",
		Address: address,
	}

	user := User{
		Name:           "nick",
		Age:            18,
		Phone:          "+8613800138000",
		Email:          "nick@gmail.com",
		FavoriteColor1: "#fff",
		FavoriteColor2: "rgb(255, 255, 255)",
		Address:        address,
		ContactUser:    []*ContactUser{contactUser1, contactUser2},
		Hobby:          []string{"篮球", "足球"},
	}

	err := v.Struct(user)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
