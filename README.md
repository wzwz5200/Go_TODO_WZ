# Go_TODO_WZ

这是一个Go语言初学者练习项目，主要包含以下功能：

## 项目结构

- main.go: 包含主函数和随机数生成功能
  - GenRand(): 生成随机数
  - main(): 程序入口

- SayHello.go: 实现CLI问候功能
  - GreetCommand(): 创建CLI命令
  - greetAction(): 处理问候命令

- act/Todo.go: TODO功能实现
- model/task.go: 任务模型定义

## 使用方法

1. 安装依赖：
   ```bash
   go mod tidy
   ```

2. 运行程序：
   ```bash
   go run main.go
   ```

## 学习目标

- 掌握Go语言基础语法
- 理解CLI应用程序开发
- 学习模块化编程
- 实践Go项目结构组织
