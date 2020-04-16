package ctrl

import (
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {

	var user models.User
	for i := 0; i < 100; i++ {
		user.GetUser(13090)
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"path":    "book_index",
		"databse": viper.Get("database.db"),
		"user":    user,
	})
}
