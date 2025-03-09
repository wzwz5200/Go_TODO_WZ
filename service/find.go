package service

import (
	initdb "GO_CLI/initDB"
	"GO_CLI/model"
	"fmt"

	"github.com/pterm/pterm"
)

func ListTask() {

	db := initdb.ReturnDB

	tableDate := pterm.TableData{
		{"ID", "创建时间", "标题", "描述", "状态"}, //
	}
	var newM []model.Task
	result := db.Find(&newM)

	if result.Error != nil {

		fmt.Println("检索所有数据时出错")

	}

	for _, Date := range newM {
		status := "❌ 未完成"
		if Date.Status {
			status = "✅ 已完成"
		}

		tableDate = append(tableDate, []string{
			fmt.Sprintf("%d", Date.ID),                   // int 转 string
			Date.CreatedAt.Format("2006-01-02 15:04:05"), // 时间格式化
			Date.Title,
			Date.Descripion, // 修正拼写错误
			status,          // bool 转 string
		})
	}

	
	pterm.DefaultTable.WithHasHeader().WithData(tableDate).Render()
}
