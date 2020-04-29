package ctrl

import (
	"blog/common/models"
	"blog/common/tools"
	"blog/dao"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Post struct {
}

func (u *Post) List(c *gin.Context) {

	page := c.PostForm("page")
	if page == "" {
		page = "1"
	}
	pageIndex, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(200, gin.H{"err": "param_err", "err_msg": "page参数错误"})
		return
	}
	pageSize, totalCount := 5, 0
	db, p, posts := dao.Instance(), tools.Paginate{}, make([]models.Post, pageSize)

	query := db.Model(&models.Post{}).Where("state_id = 10")
	query.Count(&totalCount)

	paginate := p.Init(pageSize, pageIndex, totalCount)

	query.Limit(pageSize).Offset(paginate["offset"]).Find(&posts)
	for i, v := range posts {
		posts[i].TagsData = strings.Split(strings.Trim(v.Tags, ","), ",")
	}
	c.JSON(200, gin.H{"status": "ok", "data": gin.H{"list": posts, "paginate": paginate}})
}
