package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Port     int            `yaml:"port"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

var AppConfig *Config

func InitConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("读取配置文件失败:", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalln("解析配置文件失败:", err)
	}

	AppConfig = &cfg
	log.Println("配置文件读取成功")
}
