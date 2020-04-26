package backendRoutes

import (
	"fmt"
	"gin/backend/ctrl"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	path := c.FullPath()
	fmt.Println(path)

	session := sessions.Default(c)
	isLogin := session.Get("user")
	loginUrl := "/user/login"
	fmt.Println(isLogin)
	if isLogin != nil {
		if path == loginUrl {
			c.Redirect(301, "/user/index")
			return
		}
	} else {
		if path != loginUrl {
			c.Redirect(302, loginUrl)
			return
		}
	}
}

func LoadRoute(r *gin.Engine) {

	//r.LoadHTMLGlob("backend/public/index.html")
	//r.StaticFS("/admin", http.Dir("./backend/public"))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("login_token", store))

	user := r.Group("user")
	user.Use(AuthMiddleware)
	{
		var userCtrl ctrl.User
		user.GET("index", userCtrl.Index)
		user.GET("login", userCtrl.Login)
	}

}
