package ctrl

import (
	"blog/common/models"
	"blog/common/tools"
	"blog/common/validates"
	"blog/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Post struct {
}

func (u *Post) List(c *gin.Context) {

	contextInfo, _ := c.Get("context")
	switch contextInfo.(type) {
	case error:
		c.JSON(http.StatusBadRequest, gin.H{"err": "param_error", "err_msg": contextInfo.(error).Error()})
		return
	}
	context := contextInfo.(validates.ListParam)

	pageSize, totalCount := 5, 0
	db, p, posts := dao.Instance(), tools.Paginate{}, make([]models.Post, pageSize)

	query := db.Model(&models.Post{}).Where("state_id = 10")
	query.Count(&totalCount)

	paginate := p.Init(pageSize, context.PageIndex, totalCount)

	query.Limit(pageSize).Offset(paginate["offset"]).Find(&posts)
	for i, v := range posts {
		posts[i].TagsData = strings.Split(strings.Trim(v.Tags, ","), ",")
	}
	c.JSON(200, gin.H{"status": "ok", "data": gin.H{"list": posts, "paginate": paginate}})
}
