package _json

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"` // 如果字段值为空，则忽略
	Address Address `json:"address"`
	Ignore  string  `json:"-"` // 忽略此字段
}

func Nested() {
	jsonData := `{
		"id": 1,
		"name": "John Doe",
		"address": {"city": "New York", "zip_code": "10001"}
	}`
	var employee Employee
	// json 字符串转 结构体
	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employee)

	jsonOutput, _ := json.Marshal(employee)
	fmt.Println(string(jsonOutput))
}
