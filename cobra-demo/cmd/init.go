package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "init short",
	Long:    "init long",
	Aliases: []string{"create"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("执行 init 子命令 start")
		fmt.Println(
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("license").Value,
		)
		fmt.Println("执行 init 子命令 end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
