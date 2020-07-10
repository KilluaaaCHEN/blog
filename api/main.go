package main

import (
	"blog_api/api/routes"
	"blog_api/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	routes.LoadRoute(r) //加载路由
	config.InitConfig() //初始化配置文件

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
