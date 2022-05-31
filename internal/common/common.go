package common

import (
	"net/http"
	"src/web-admin/internal/config"

	"github.com/gin-gonic/gin"
)

// func Display(c *gin.Context, tplName string, content gin.H) {
func Display(c *gin.Context, content gin.H) {
	header := gin.H{
		"title": config.GlobalConfig.Name,
	}
	footer := gin.H{
		"intro": "at end...",
	}
	params := gin.H{
		"header":  header,
		"footer":  footer,
		"content": content,
	}
	c.HTML(http.StatusOK, "index.html", params)
}
