package main

import (
	"context"
	"flag"
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
	log.Printf("")

	router := gin.Default()

	// 静态资源加载，本例为css,js以及资源图片
	router.StaticFS("/public", http.Dir("./src/web-admin/internal/view/static"))
	router.StaticFile("/favicon.ico", "./src/web-admin/internal/view/static/icon/favicon.ico")

	// 导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("./src/web-admin/internal/view/tpl/*/*")

	// 使用全局CORS中间件。
	router.Use(handler.Cors())

	// website分组
	r := router.Group("/")
	handler.RegisterHandlers(r)

	// 前端走代理的话，不同分组可以考虑拆分为不同子域名
	// Listen and serve on 0.0.0.0:8080
	// router.Run(":80") 这样写就可以了，下面所有代码（go1.8+）是为了优雅处理重启等动作。可根据实际情况选择。

	srv := &http.Server{
		Addr:         ":80",
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

	// 优雅Shutdown（或重启）服务
	// 5秒后优雅Shutdown服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) //syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	// 1.fvbock/endless endless.ListenAndServe(":80", router) 能优雅restart or stop WEB服务
	// 2.http.ListenAndServe(":80", router) 直接使用http.ListenAndServe()，也可以gracehttp
	// 2.具体实现可以参考https://github.com/tabalt/gracehttp#demo
}
