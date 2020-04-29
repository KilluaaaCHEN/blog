package validates

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
}

type ListParam struct {
	PageIndex int    `json:"page" binding:"required" default:"1"`
	Query     string `json:"query" default:"0"`
}

func (*Post) ValidateList(c *gin.Context) {
	var lp ListParam
	if err := c.ShouldBindJSON(&lp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "param_error", "err_msg": err.Error()})
		c.Abort()
		return
	}
	c.Set("context", lp)
}
