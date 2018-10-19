package main

import (
	"fmt"
	"github.com/zhanghe06/gin_project/config"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/logs"
	"github.com/zhanghe06/gin_project/routers"
)

func main() {
	// 初始化配置
	config.Init()
	config.Watch()

	// 初始化日志
	logs.Init()

	// 初始化数据库
	err := dbs.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer dbs.Close()

	// 初始化路由
	router := routers.Init()

	// 启动服务
	router.Run()
}
