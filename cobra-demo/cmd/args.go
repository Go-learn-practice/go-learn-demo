package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// 自定义验证
var curArgsCheckCmd = &cobra.Command{
	Use:   "cus",
	Long:  "",
	Short: "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("自定义参数校验 start")
		fmt.Println(args)
		fmt.Println("自定义参数校验 end")
	},
}

// 无参数验证
var noArgsCmd = &cobra.Command{}

// 可以接收任何参数
var arbitrayArgsCmd = &cobra.Command{}

var onlyArgsCmd = &cobra.Command{}

var exactArgs = &cobra.Command{}

var maxArgs = &cobra.Command{}

func init() {
	rootCmd.AddCommand(curArgsCheckCmd)
}
