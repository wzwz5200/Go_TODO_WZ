package service

import (
	initdb "GO_CLI/initDB"
	"GO_CLI/model"
	"fmt"
)

func Create(T string, D string, S bool) {

	newTask := model.Task{Title: T, Descripion: D, Status: S}

	db := initdb.ReturnDB

	result := db.Create(&newTask)

	if result.Error != nil {

		fmt.Println("创建记录错误")

	}

	// var user model.Task
	// db.First(&user, 1) // 查询 ID 为 1 的用户
	// log.Println("User:", user)
}
