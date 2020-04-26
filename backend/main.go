package main

import (
	"gin/backend/routes"
	"gin/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	backendRoutes.LoadRoute(r) //加载路由
	config.InitConfig() //初始化配置文件

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))

	//r.GET("/hello", func(c *gin.Context) {
	//	session := sessions.Default(c)
	//	name := c.Query("name")
	//	fmt.Println(name)
	//	session.Set("hello", name)
	//	session.Save()
	//	c.JSON(200, gin.H{"hello": session.Get("hello")})
	//})
	//
	//r.GET("/hello2", func(c *gin.Context) {
	//	session := sessions.Default(c)
	//	c.JSON(200, gin.H{"hello": session.Get("hello")})
	//})

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
