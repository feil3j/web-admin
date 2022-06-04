package player

import (
	"src/web-admin/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func TaskHandler(c *gin.Context) {
	base.Display(c, "player_task.html", gin.H{
		"Title": "basic",
		"token": "abcd",
	})
}
