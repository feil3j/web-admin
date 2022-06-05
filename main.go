package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"src/web-admin/internal/config"
	"src/web-admin/internal/handler"

	"github.com/gin-gonic/gin"
)

var globalConfig *config.Config

func main() {
	//加载配置
	configFile := flag.String("f", "./src/web-admin/etc/admin.conf", "the config file")
	flag.Parse()
	globalConfig := config.LoadConfig(configFile)
	log.Printf("main: globalConfig=%#v", globalConfig)

	router := gin.Default()

	// 静态资源加载，例如css,js以及资源图片
	router.StaticFS("/public", http.Dir(globalConfig.AdminRootPath+"/internal/view/static"))
	router.StaticFile("/favicon.ico", globalConfig.AdminRootPath+"/internal/view/static/icon/favicon.ico")

	// 导入所有模板，多级目录结构需要这样写
	pattern := globalConfig.AdminRootPath + "/internal/view/tpl/*/*"
	templ := template.Must(template.New("").Delims("<{", "}>").ParseGlob(pattern))
	router.SetHTMLTemplate(templ)
	config.SetHtmlTemplate(templ)

	// 使用全局CORS中间件。
	router.Use(handler.Cors())

	// website分组
	r := router.Group("/")
	//注册路由
	handler.RegisterHandlers(r)

	//启动服务
	addr := globalConfig.Server.Host + ":" + globalConfig.Server.Port
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server start success, addr=%s.", addr)
	addr = ""

	// 5秒后优雅Shutdown服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill) //syscall.SIGINT
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
