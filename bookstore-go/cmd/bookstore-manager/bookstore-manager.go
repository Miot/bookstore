package main

import (
	"bookstore/config"
	"bookstore/global"
)

func main() {
	// 初始化配置
	config.InitConfig("conf/config.yaml")
	// 初始化数据库
	global.InitMysql()
	// 初始化Redis
	global.InitRedis()
}
