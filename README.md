# 终端待办事项应用

一个用Go语言编写的命令行待办事项管理工具，使用SQLite数据库存储数据。

## 功能特性

- 添加新待办事项
- 标记待办事项为已完成
- 列出所有待办事项
- 查看待办事项详情
- 删除待办事项

## 安装使用

1. 确保已安装Go 1.20+ 和 SQLite3
2. 克隆仓库：
   ```bash
   git clone https://github.com/yourusername/todo-cli.git
   ```
3. 进入项目目录：
   ```bash
   cd todo-cli
   ```
4. 构建项目：
   ```bash
   go build -o todo
   ```
5. 运行程序：
   ```bash
   ./todo
   ```

## 命令说明

- `add` - 添加新待办事项
- `list` - 列出所有待办事项
- `done` - 标记待办事项为已完成

## 项目结构

```
.
├── cmd/            # 命令行命令实现
│   ├── add.go      # 添加命令
│   ├── done.go     # 完成命令
│   ├── list.go     # 列表命令
│   └── root.go     # 根命令
├── initDB/         # 数据库初始化
│   └── sqlite.go   # SQLite数据库配置
├── model/          # 数据模型
│   └── strut.go    # 数据结构定义
├── service/        # 业务逻辑
│   ├── create.go   # 创建服务
│   ├── done.go     # 完成服务
│   └── find.go     # 查询服务
├── go.mod          # Go模块文件
├── go.sum          # Go模块校验文件
├── main.go         # 程序入口
```

## 许可证

本项目采用MIT许可证。详见[LICENSE](LICENSE)文件。