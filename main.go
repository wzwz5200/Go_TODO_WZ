/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"GO_CLI/cmd"
	initdb "GO_CLI/initDB"
)


func main() {
	initdb.InitDB()
	cmd.Execute()
}
