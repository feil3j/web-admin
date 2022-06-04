package player

import (
	"src/web-admin/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	base.Display(c, "", gin.H{
		"Data": "hello world",
	})
}
