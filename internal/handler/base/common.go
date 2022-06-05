package base

import (
	"net/http"
	"src/web-admin/internal/config"

	"github.com/gin-gonic/gin"
)

func Display(c *gin.Context, tplName string, data gin.H) {
	header := gin.H{
		"title": config.GlobalConfig.Name,
	}
	footer := gin.H{
		"intro": "at end...",
	}
	content := ""
	if tplName != "" {
		content = config.GetTemplateData(tplName, data)
	}
	params := gin.H{
		"header":     header,
		"menu":       menu,
		"footer":     footer,
		"content":    content,
		"PATHPREFIX": "/",
	}
	c.HTML(http.StatusOK, "index.html", params)
}
