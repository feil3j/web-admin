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
	"/index.html": []interface{}{"GET", player.IndexHandler},
	"/add.html":   []interface{}{"GET", player.AddHandler},
	"/post.html":  []interface{}{"GET", player.PostmeHandler},
}

//注册路由
func RegisterHandlers(r *gin.RouterGroup) {
	for path, data := range routers {
		if len(data) != 2 {
			log.Fatalf("RegisterHandlers: len(data) is error, len=%d.", len(data))
		}
		h, ok := data[0].(gin.HandlerFunc)
		if !ok {
			log.Fatalf("RegisterHandlers: handle is error, path=%s, data=%v.", path, data)
		}
		tmp, ok := data[1].(string)
		if !ok {
			log.Fatalf("RegisterHandlers: method is not string, path=%s, data=%v.", path, data)
		}
		method := strings.ToUpper(tmp)
		if method == "Get" {
			r.GET(path, h)
		} else {
			r.POST(path, h)
		}
	}
}
