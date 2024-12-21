package _flag

import (
	"fmt"
	"os"
)

func RunOsArgs() {
	// 打印所有命令行参数
	fmt.Println("All Args:", os.Args)

	// 打印程序名称
	fmt.Println("Program Name:", os.Args[0])

	// 打印参数个数
	fmt.Println("Number of Args:", len(os.Args)-1)

	// 打印每个参数
	for i, arg := range os.Args {
		fmt.Printf("Arg[%d]: %s\n", i, arg)
	}

	// 检查是否有足够的参数
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <param1> <param2>")
		return
	}

	// 获取第一个和第二个参数
	param1 := os.Args[1]
	param2 := os.Args[2]

	fmt.Println("First Parameter:", param1)
	fmt.Println("Second Parameter:", param2)
}
