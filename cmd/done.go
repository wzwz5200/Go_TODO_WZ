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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "此命令用来完成某一个事项",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("done called")
		service.ListTask()
		InputText := pterm.DefaultInteractiveTextInput

		InputText.DefaultText = "请输入完成事项的ID"

		result, _ := InputText.Show()

		fmt.Println(result)

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
