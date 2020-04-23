package routes

import (
	"gin/ctrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRoute(r *gin.Engine) {

	r.LoadHTMLGlob("backend/public/index.html")
	r.StaticFS("/admin", http.Dir("./backend/public"))

	user := r.Group("user")
	{
		var userCtrl ctrl.User
		user.GET("index", userCtrl.Index)
	}

	post := r.Group("post")
	{
		var postCtrl ctrl.Post
		post.GET("index", postCtrl.Index)
		post.Any("view", postCtrl.View)
	}
}
