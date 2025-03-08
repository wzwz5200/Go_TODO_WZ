package main

import (
	"day/act"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 生成随机数（可用于生成任务 ID）
func GenRand() int {
	rand.Seed(time.Now().UnixNano()) // 确保随机数不同
	return rand.Intn(100)
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 确保至少有一个子命令
	if len(os.Args) < 2 {
		fmt.Println("请提供子命令: add | list | done")
		return
	}

	switch os.Args[1] {
	case "add":
		// 创建 add 子命令
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)

		// 解析 add 后面的参数
		addCmd.Parse(os.Args[2:])

		// 确保有足够的参数
		if addCmd.NArg() < 2 {
			fmt.Println("使用方式: add <标题> <描述>")
			return
		}

		// 获取标题和描述
		Title := addCmd.Arg(0)
		Description := addCmd.Arg(1)

		fmt.Println("添加事项 ✅")
		fmt.Println("标题:", Title)
		fmt.Println("描述:", Description)

		// 调用 act.AddToDo 处理任务添加
		act.AddToDo(Title, Description)

	case "list":
		// 列出所有任务
		act.ListTask()

	case "done":
		// 解析 done 后的参数
		if len(os.Args) < 3 {
			fmt.Println("使用方式: done <ID>")
			return
		}

		// 获取 ID 参数并转换
		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID 无效:", err)
			return
		}

		// 调用 act.DoneToDo 处理任务完成
		act.DoneToDo(ID)

	default:
		fmt.Println("未知子命令:", os.Args[1])
		fmt.Println("支持的命令: add | list | done")
	}
}
