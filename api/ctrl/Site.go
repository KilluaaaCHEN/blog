package ctrl

import (
	"blog/common/models"
	"blog/common/tools/oauth2/github"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Site struct {
}

func (u *Site) Index(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	userName := session.Get("user_name")
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user_id": userId, "user_name": userName})
}

/*func (u *Site) Login(c *gin.Context) {
	method := c.Request.Method
	if method == "GET" {
		c.HTML(http.StatusOK, "login", gin.H{"a": 1})
	} else {
		userName, password := c.PostForm("user_name"), c.PostForm("password")
		if userName == "killua" && password == "admin888" {
			session := sessions.Default(c)
			session.Set("site", userName)
			session.Save()
			c.Redirect(301, "/site/index")
		} else {
			c.Redirect(301, "/site/login?error=1")
		}
	}
}*/

/**
Github授权回调
*/
func (u *Site) Callback(c *gin.Context) {
	code := c.Query("code")
	var gh github.Github
	ghUser, err := gh.GetUserInfo(code)
	if err != nil {
		c.JSON(200, gin.H{"err": "server_error", "err_msg": err.Error()})
		return
	}
	var user models.User
	user, err = user.LoginGithubUser(c, ghUser, c.ClientIP())
	if err != nil {
		c.JSON(200, gin.H{"err": "user_disable", "err_msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "user_info": user})
}
