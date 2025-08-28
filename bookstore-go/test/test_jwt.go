package main

import (
	"bookstore/config"
	"bookstore/global"
	"bookstore/jwt"
	"fmt"
	"log"
)

func main() {
	config.InitConfig("conf/config.yaml")
	global.InitRedis()
	res, err := jwt.GenerateTokenPair(1, "admin")
	if err != nil {
		log.Fatalln("err:", err)
	}
	fmt.Println("token:", res.AccessToken)
	claim, err := jwt.ParseToken(res.AccessToken)
	if err != nil {
		log.Fatalln("err:", err)
	}
	fmt.Println("claim:", claim.Username)
}
