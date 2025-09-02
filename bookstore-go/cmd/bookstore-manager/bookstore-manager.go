package main

import (
	"bookstore/config"
	"bookstore/global"
	"bookstore/web/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// 创建一个通道来接收操作系统信号
	quit := make(chan os.Signal, 1)
	// 监听中断信号和终止信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 启动一个goroutine来监听信号
	go func() {
		// 阻塞直到收到信号
		<-quit
		log.Println("接收到关闭信号，正在关闭服务...")

		// 创建一个5秒超时的context
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 优雅地关闭HTTP服务器
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("服务器关闭错误:", err)
		}
	}()

	log.Printf("服务启动成功: %s\n", addr)

	// 启动HTTP服务器
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("启动服务失败: %v", err)
	}

	// 清理资源
	cleanResources()
}

func cleanResources() {
	if global.RedisClient != nil {
		log.Println("关闭Redis")
		global.CloseRedis()
	}
	if global.DBClient != nil {
		log.Println("关闭数据库")
		global.CloseDB()
	}

	time.Sleep(1 * time.Second)
	log.Println("关闭资源成功")
}
