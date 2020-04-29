package routes

import (
	"blog/api/ctrl"
	"blog/common/validates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	path := c.FullPath()
	session := sessions.Default(c)
	if session.Get("site") == nil && path != "/site/login" {
		c.JSON(200, gin.H{"status": "err", "err_msg": "请登录"})
		return
	}
}

func LoadRoute(r *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("login", store))

	var siteCtrl ctrl.Site
	var postCtrl ctrl.Post
	var vaPost validates.Post

	r.GET("/", siteCtrl.Index)
	r.GET("oauth2/callback", siteCtrl.Callback)
	r.POST("/posts", vaPost.ValidateList, postCtrl.List)

}
