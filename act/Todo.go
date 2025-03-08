package act

import (
	"day/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 1.为什么我要用[]数组而不是切片{}
func GenRand() int {

	n := rand.Intn(100)
	return n
}

func GetNextId(T *[]model.Tasks) int {

	MaxID := 1
	for _, T := range *T {
		if T.Id > MaxID {
			MaxID = T.Id

		}

	}
	return MaxID + 1

}

func AddToDo(Title string, Dsc string) {

	var tasks []model.Tasks
	// 新任务数据
	fileData, err := os.ReadFile("task.json")
	if err == nil {
		// 文件存在，解析 JSON 数据
		if err := json.Unmarshal(fileData, &tasks); err != nil {
			fmt.Println("解析 JSON 错误:", err)
			return
		}
	} else if !os.IsNotExist(err) {
		// 其他错误处理
		fmt.Println("读取文件错误:", err)
		return
	}
	newTask := model.Tasks{GetNextId(&tasks), Title, Dsc, time.Now(), "NO"}

	// 尝试读取现有文件数据

	fmt.Println("事项ID：", newTask.Id)
	// 将新任务添加到现有任务数组中
	tasks = append(tasks, newTask)

	// 将任务数组序列化为 JSON 数据
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("序列化 JSON 错误:", err)
		return
	}

	// 安全写入文件
	if err := os.WriteFile("task.json", data, 0644); err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}
}

func ListTask() {
	fileDate, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("读取文件错误")

	}
	var task []model.Tasks
	if err := json.Unmarshal(fileDate, &task); err != nil {
		fmt.Println("解析 JSON 错误:", err)
		return
	}

	for _, taskz := range task {
		fmt.Println("📋 \033[1;35m任务详情\033[0m 📋")
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Printf("| 🆔 \033[1;36mID\033[0m              | %-40d |\n", taskz.Id)
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Printf("| 📌 \033[1;36m标题\033[0m            | %-40s |\n", taskz.Title)
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Printf("| 📝 \033[1;36m描述\033[0m            | %-40s |\n", taskz.Description)
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Printf("| 📅 \033[1;36m日期\033[0m            | %-40s |\n", taskz.Date.Format(time.RFC3339))
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Printf("| 🚦 \033[1;36m状态\033[0m            | %-40s |\n", taskz.Status)
		fmt.Println("+------------------+------------------------------------------+")
		fmt.Println("🎉 \033[1;32m任务展示完毕\033[0m 🎉")
		fmt.Println() // 空行分隔每个任务
	}

}

func DoneToDo(ID int) {
	// 读取 JSON 文件
	fileData, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	// 解析 JSON
	var tasks []model.Tasks
	if err := json.Unmarshal(fileData, &tasks); err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	// 查找任务并修改状态
	found := false
	for i, task := range tasks {
		if task.Id == ID {
			tasks[i].Status = "YES"
			found = true
			break
		}
	}

	// 如果没有找到任务
	if !found {
		fmt.Printf("任务 ID %d 不存在\n", ID)
		return
	}

	// 将修改后的任务列表写回 JSON 文件
	newFileData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}

	err = os.WriteFile("task.json", newFileData, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Printf("任务 ID %d 已标记为完成 ✅\n", ID)
}
