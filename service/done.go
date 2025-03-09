package service

import (
	initdb "GO_CLI/initDB"
	"GO_CLI/model"
	"fmt"
)

func DoneTask(ID int) {

	db := initdb.ReturnDB

	var Task model.Task
	result := db.First(&Task, ID)

	if result.Error != nil {
		fmt.Println("未找到任务:", result.Error)
		return
	}

	// 切换状态（如果是 false 变 true，true 变 false）
	Task.Status = !Task.Status

	// 更新数据库
	db.Save(&Task)

	fmt.Printf("任务 ID %d 状态已更新为: %v\n", Task.ID, IsDone(Task))

}

func IsDone(Task model.Task) string {

	if Task.Status != true {
		return "未完成"

	}

	return "完成"
}
