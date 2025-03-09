/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	initdb "GO_CLI/initDB"
	"log"

	"github.com/spf13/cobra"
)

// parentCmd represents the parent command
var parentCmd = &cobra.Command{
	Use:   "parent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("parent called")

		// textInput := pterm.DefaultInteractiveTextInput

		// textInput.DefaultText = "请输入文字"
		// result, _ := textInput.Show()
		// fmt.Println(result)

		db := initdb.ReturnDB
		db.Create(&initdb.User{Name: "Alice"})

		// 查询数据
		var user initdb.User
		db.First(&user, 1) // 查询 ID 为 1 的用户
		log.Println("User:", user)

	},
}

func init() {
	rootCmd.AddCommand(parentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
