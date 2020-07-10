package ctrl

import (
	"blog_api/common/models"
	"blog_api/common/tools"
	"blog_api/common/tools/oauth2/github"
	"github.com/gin-gonic/gin"
)

type Site struct {
}

func (u *Site) List(c *gin.Context) {
	items := make([]gin.H, 3)
	items[0] = gin.H{"id": 1, "title": "title111", "status": 1, "author": "killua", "display_time": "2020-06-03", "pageviews": 30}
	items[1] = gin.H{"id": 2, "title": "title222", "status": 1, "author": "killua", "display_time": "2020-06-03", "pageviews": 30}
	items[2] = gin.H{"id": 3, "title": "title333", "status": 1, "author": "killua", "display_time": "2020-06-03", "pageviews": 30}
	tools.Result(c, gin.H{"total": 3, "items": items})
}

/**
Github授权回调
*/
func (u *Site) Callback(c *gin.Context) {
	type CallbackForm struct {
		Code string `binding:"required"`
	}
	var cf CallbackForm
	if err := c.ShouldBindJSON(&cf); err != nil {
		tools.Error(c, "err", err.Error())
		return
	}
	var gh github.Github
	ghUser, err := gh.GetUserInfo(cf.Code)
	if err != nil {
		tools.Error(c, "err", err.Error())
		return
	}
	var user models.User
	user, err = user.LoginGithubUser(c, ghUser, c.ClientIP())
	if err != nil {
		tools.Error(c, "user_disable", err.Error())
		return
	}
	if user.Id != 1 {
		tools.Error(c, "no_auth", "您的账号未授权,不能登录管理后台")
		return
	}
	tools.Result(c, user)
}

func (u *Site) Logout(c *gin.Context) {
	tools.Result(c, "success")
}

func (u *Site) Info(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		tools.Error(c, "no_auth", "请重新登录")
		return
	}
	tools.Result(c, user)
}
