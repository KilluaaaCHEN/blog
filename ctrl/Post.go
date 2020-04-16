package ctrl

import (
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
}

func (*Post) View(c *gin.Context) {
	type View struct {
		Id int `json:"id" form:"id" binding:"required"`
	}
	var param View
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	post = post.GetPost(param.Id)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (*Post) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "path": "post_index2"})
}
