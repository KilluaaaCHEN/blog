package ctrl

import (
	"blog/common/models"
	"blog/common/tools"
	"blog/common/validates"
	"blog/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

type Post struct {
}

func (u *Post) List(c *gin.Context) {

	contextInfo, _ := c.Get("context")
	context := contextInfo.(validates.ListParam)
	redis := dao.InstRedis()
	cacheKey := fmt.Sprintf("post:list:%d_%s", context.PageIndex, context.Query)
	cache, err := redis.Get(cacheKey).Result()
	data := gin.H{}
	if err != nil {
		pageSize, totalCount := 5, 0
		db, p, posts := dao.InstDB(), tools.Paginate{}, make([]models.Post, pageSize)

		query := db.Model(&models.Post{}).Where("state_id = 10")
		query.Count(&totalCount)

		paginate := p.Init(pageSize, context.PageIndex, totalCount)

		query.Limit(pageSize).Offset(paginate["offset"]).Find(&posts)
		for i, v := range posts {
			posts[i].TagsData = strings.Split(strings.Trim(v.Tags, ","), ",")
		}
		data = gin.H{"list": posts, "paginate": paginate}
		dataByte, _ := json.Marshal(data)
		cache := string(dataByte)
		redis.Set(cacheKey, cache, time.Minute)
	} else {
		str := []byte(cache)
		if err := json.Unmarshal(str, &data); err != nil {
			log.Fatal(err)
		}
	}
	c.JSON(200, gin.H{"status": "ok", "data": data})
}
