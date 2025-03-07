package main

import (
	"day/act"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func GenRand() int {

	n := rand.Intn(100)
	return n
}

func main() {
	// 定义参数：名称、默认值、帮助信息

	//ADD := flag.Bool("ADD", false, "提供此参数时打印 Hello")

	// 解析命令行参数
	flag.Parse()

	// 检查是否提供了 `-hello`
	// if *ADD {

	// } else {
	// 	fmt.Println("你好！这是一个终端任务代办事项应用📝")
	// }

	if len(os.Args) < 2 {
		fmt.Println("请提供子命令")
		return
	}

	switch os.Args[1] {
	case "Add":

		addCmd := flag.NewFlagSet("ADD", flag.ExitOnError)
		//	verbose := addCmd.Bool("v", false, "启用详细模式")

		// 解析 ADD 子命令的参数
		addCmd.Parse(os.Args[2:])

		fmt.Println("添加事项")

		flag.CommandLine.Parse(os.Args[2:])

		//解析 add 后面的参数
		if flag.NArg() < 2 {
			fmt.Println("使用方式: ADD 标题 描述")
			return
		}

		Title := flag.Arg(0)       // 第一个参数
		Description := flag.Arg(1) // 第二个参数

		fmt.Println("标题:", Title)
		fmt.Println("描述", Description)

		act.AddToDo(Title, Description)

	case "list":

		act.ListTask()

	default:
		fmt.Println("未知子命令")
	}

	// 解析命令行参数
	flag.Parse()

	// 使用参数（需解引用指针）

	//act.AddToDo()

}
