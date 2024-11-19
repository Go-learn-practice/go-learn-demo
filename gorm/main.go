package main

import (
	"context"
	"fmt"
	"gorm/config"
	"gorm/dao"
	"gorm/dao/model"
	"gorm/generator"
	"gorm/migrate"
	"os"
)

func main() {
	// 初始化数据库连接
	config.InitDB()

	// 判断命令行参数执行相应功能
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [migrate|generate|run]")
		return
	}

	switch os.Args[1] {
	case "migrate":
		migrate.Migrate()
	case "generate":
		generator.Generate()
	case "run":
		runBusinessLogic()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func runBusinessLogic() {
	// 使用生成的 Query API
	q := dao.Use(config.DB)

	// 新增用户
	newUser := model.User{
		Name:     "Alice",
		Age:      25,
		Email:    "alice@example.com",
		Password: "123456789",
	}
	if err := q.User.WithContext(context.Background()).Create(&newUser); err != nil {
		fmt.Println("Failed to create user:", err)
		return
	}
	fmt.Printf("User created: %+v\n", newUser)

	// 查询用户
	user, err := q.User.WithContext(context.Background()).Where(q.User.Name.Eq("Alice")).First()
	if err != nil {
		fmt.Println("Failed to query user:", err)
		return
	}
	fmt.Printf("User queried: %+v\n", user)

	// 更新用户
	user.Email = "newalice@example.com"
	if _, err := q.User.WithContext(context.Background()).Where(q.User.ID.Eq(user.ID)).Update(q.User.Email, user.Email); err != nil {
		fmt.Println("Failed to update user:", err)
		return
	}
	fmt.Println("User updated successfully!")

	// 删除用户
	//if _, err := q.User.WithContext(context.Background()).Where(q.User.ID.Eq(user.ID)).Delete(); err != nil {
	//	fmt.Println("Failed to delete user:", err)
	//	return
	//}
	//fmt.Println("User deleted successfully!")
}
