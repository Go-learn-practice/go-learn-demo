package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
标志类型：
持久标志：当前命令以及子命令与其继承命令
全局标志：由持久标志衍生
本地标志：当前命令可用
*/

var rootCmd = &cobra.Command{
	Use:   "mycobra",
	Short: "简短的描述",
	Long:  "详细的描述",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec mycobra command")
	},
}

func Execute() {
	rootCmd.Execute()
}

var userLicense string

func init() {
	rootCmd.PersistentFlags().Bool("viper", true, "是否采用viper作配置文件读取")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "作者名称")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "授权信息")
	rootCmd.Flags().StringP("source", "s", "", "来源")
}
