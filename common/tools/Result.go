package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func Result(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": data})
}

func Error(c *gin.Context, code string, errMsg string) {
	c.JSON(http.StatusOK, gin.H{"status": code, "errMsg": errMsg})
}
