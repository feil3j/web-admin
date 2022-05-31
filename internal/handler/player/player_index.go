package player

import (
	"src/web-admin/internal/common"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	common.Display(c, gin.H{
		"Data": "hello world",
	})
}
