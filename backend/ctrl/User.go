package ctrl

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"status": "ok",
		"path":   "admin/user/index",
	})
}

func (u *User) Login(c *gin.Context) {
	session := sessions.Default(c)

	if c.Query("login") == "1" {
		session.Set("user", "killuachen")
		session.Save()
		c.JSON(200, gin.H{
			"status": "ok",
			"path":   "登录成功了",
		})
	}
	fmt.Println(session.Get("user"))

	c.JSON(200, gin.H{
		"status": "ok",
		"path":   "admin/user/login",
	})
}
