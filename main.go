package main

import (
	"go-blog/model"
	"go-blog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 启动路由
	routes.InitRouter()
}
