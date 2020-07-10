package routes

import (
	"blog_api/backend/ctrl"
	"blog_api/common/models"
	"blog_api/common/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		token := c.Request.Header.Get("X-Token")
		if path != "/oauth2/callback" {
			if "" == token {
				tools.Error(c, "no_auth", "请登录")
				c.Abort()
				return
			}
			var userService models.User
			user, err := userService.GetUserByPwd(token)
			if err != nil {
				tools.Error(c, "no_auth", err.Error())
				c.Abort()
				return
			}
			if user.Id != 1 {
				tools.Error(c, "no_auth", "您的账号未授权,不能登录管理后台")
				c.Abort()
				return
			}

			if user.State != 1 {
				tools.Error(c, "no_auth", "账号违规,已被禁用")
				c.Abort()
				return
			}

			c.Set("token", token)
			c.Set("user", user)
		}
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
}

func LoadRoute(r *gin.Engine) {

	r.Use(Cors(), AuthMiddleware())

	var siteCtrl ctrl.Site
	r.POST("oauth2/callback", siteCtrl.Callback)
	r.POST("user/logout", siteCtrl.Logout)

	r.POST("user/info", siteCtrl.Info)

	r.GET("user/list", siteCtrl.List)

}
