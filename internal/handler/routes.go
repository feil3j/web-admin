// 处理website站点路由请求的handler，考虑这里使用了模板，结合MVC使用的习惯，故以controller来命名。
// 当然，API也可以都放在这里处理，但为了结构清晰，不提倡这么做。
package handler

import (
	"src/web-admin/internal/handler/player"
	"src/web-admin/internal/handler/world"

	"github.com/gin-gonic/gin"
)

//玩家信息路由表
var playerRouters = map[string]gin.HandlerFunc{
	"index": player.IndexHandler,
	"basic": player.BasicHandler,
	"task":  player.TaskHandler,
}

//世界信息路由表
var worldRouters = map[string]gin.HandlerFunc{
	"index": world.IndexHandler,
}

//注册路由
func RegisterHandlers(r *gin.RouterGroup) {
	//player route
	for module, h := range playerRouters {
		path := "/player/" + module
		r.POST(path, h)
	}

	//world route
	for module, h := range worldRouters {
		path := "/world/" + module
		r.POST(path, h)
	}
}

// 定义全局的CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
