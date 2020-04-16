package routes

import (
	"gin/ctrl"
	"github.com/gin-gonic/gin"
)

func LoadRoute(r *gin.Engine) {

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
