/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"GO_CLI/service"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "此命令用来创建代办事务",
	Run: func(cmd *cobra.Command, args []string) {

		TextInput_title := pterm.DefaultInteractiveTextInput

		TextInput_title.DefaultText = "输入标题"
		result, _ := TextInput_title.Show()

		fmt.Println(result) //

		TextInput_dsc := pterm.DefaultInteractiveTextInput

		TextInput_dsc.DefaultText = "输入事项描述"

		result_dsc, _ := TextInput_dsc.Show()

		fmt.Println(result_dsc)

		service.Create(result, result_dsc, false)

		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
