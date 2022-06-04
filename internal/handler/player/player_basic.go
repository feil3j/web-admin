package player

import (
	"src/web-admin/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func BasicHandler(c *gin.Context) {
	base.Display(c, "", gin.H{
		"Title": "basic",
		"token": "abcd",
	})
}
