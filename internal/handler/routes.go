// 处理website站点路由请求的handler，考虑这里使用了模板，结合MVC使用的习惯，故以controller来命名。
// 当然，API也可以都放在这里处理，但为了结构清晰，不提倡这么做。
package handler

import (
	"log"
	"src/web-admin/internal/handler/player"
	"strings"

	"github.com/gin-gonic/gin"
)

//路由表
var routers = map[string][]interface{}{
	"/index.html":         {"GET", player.IndexHandler},
	"/player/:id/*action": {"GET", player.AddHandler},
	"/player/post.html":   {"POST", player.PostmeHandler},
}

//注册路由
func RegisterHandlers(r *gin.RouterGroup) {
	for path, data := range routers {
		if len(data) != 2 {
			log.Fatalf("RegisterHandlers: len(data) is error, len=%d.", len(data))
		}
		tmp, ok := data[0].(string)
		if !ok {
			log.Fatalf("RegisterHandlers: method is not string, path=%s, data=%v.", path, data)
		}
		method := strings.ToUpper(tmp)

		h, ok := data[1].(gin.HandlerFunc)
		if !ok {
			log.Fatalf("RegisterHandlers: handle is error, path=%s, data=%v.", path, data)
		}
		if method == "GET" {
			r.GET(path, h)
		} else {
			r.POST(path, h)
		}
	}
}

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
