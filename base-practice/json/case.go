package _json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"` // 使用 `json` 标签指定 JSON 字段名
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (u *User) String() string {
	return "User: " + u.Name
}

func RunStruct2Json() {
	user := User{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	// 将结构体转为 JSON 字符串
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error converting struct to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func RunJson2Struct() {
	var jsonData string = `{"name": "Bob", "age": 30, "email": "bob@example.com"}`
	var user User
	// 将 JSON 字符串转为结构体
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error converting JSON to struct:", err)
	}
	fmt.Println(user)
}
