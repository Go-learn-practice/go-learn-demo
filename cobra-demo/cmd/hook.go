package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var hookRootCmd = &cobra.Command{
	Use: "hook",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("钩子案例 run 函数")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// run 函数之前执行
		fmt.Println("PersistentPreRun")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// run 函数之后执行
		fmt.Println("PersistentPostRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// run 函数之前执行
		fmt.Println("PreRun")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// run 函数之后执行
		fmt.Println("PostRun")
	},
}

func init() {
	rootCmd.AddCommand(hookRootCmd)
}
