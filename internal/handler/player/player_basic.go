package player

import (
	"net/http"

	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func BasicHandler(c *gin.Context) {
	time := time.Now().Unix()
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time, 10)))
	token := hex.EncodeToString(h.Sum(nil))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "basic",
		"token": token,
	})
}
