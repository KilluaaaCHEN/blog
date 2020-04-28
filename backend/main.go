package main

import (
	"blog/backend/routes"
	"blog/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	backendRoutes.LoadRoute(r) //加载路由
	config.InitConfig()        //初始化配置文件

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
