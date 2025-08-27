package global

import (
	"bookstore/config"
	"context"
	"fmt"

	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB
var RedisClient *redis.Client

func InitMysql() {
	mysqlConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}

	DBClient = client
	log.Println("数据库连接成功")
}

func InitRedis() {
	redisConfig := config.AppConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	str, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalln("连接Redis失败:", err)
	}

	log.Println("Redis连接成功:", str)
}

func GetDB() *gorm.DB {
	return DBClient
}

func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err != nil {
			log.Println("关闭数据库失败:", err)
		}
		sqlDB.Close()
		DBClient = nil
	}
}
