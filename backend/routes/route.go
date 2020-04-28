package backendRoutes

import (
	"blog/backend/ctrl"
	"github.com/gin-contrib/sessions"
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

	//store := cookie.NewStore([]byte("secret"))
	//r.Use(sessions.Sessions("login_token", store), AuthMiddleware)

	var siteCtrl ctrl.Site

	r.GET("oauth2/callback", siteCtrl.Callback)

	//site := r.Group("site")
	//{
	//	site.POST("index", siteCtrl.Index)
	//	site.POST("login", siteCtrl.Login)
	//	site.POST("logout", siteCtrl.Login)
	//	site.POST("setting", siteCtrl.Setting)
	//}

}
