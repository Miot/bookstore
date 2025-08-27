package main

import (
	"bookstore/config"
	"bookstore/global"
	"bookstore/web/router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 初始化配置
	config.InitConfig("conf/config.yaml")
	// 初始化数据库
	global.InitMysql()
	// 初始化Redis
	global.InitRedis()
	// 初始化路由
	r := router.InitRouter()
	// 启动服务
	addr := fmt.Sprintf("%s:%d", "localhost", config.AppConfig.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("启动服务失败:", err)
		os.Exit(-1)
	}
}
