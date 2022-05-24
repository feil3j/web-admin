# web-admin
refer: 
	github.com/ffhelicopter/tmm
	github.com/zeromicro/go-zero

快速搭建WEB站点以及提供RESTful接口 。

## 一：静态资源站点

	

``` go
	router := gin.Default()

	// 静态资源加载，本例为css,js以及资源图片
	router.StaticFS("/public", http.Dir("D:/goproject/src/src/web-admin/internal/website/static"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
```

## 二：动态站点
模板可调用静态资源站点的css，图片等

	

``` go
// 导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("website/tpl/*/*")
	
	// website分组
	v := router.Group("/")
	{

		v.GET("/index.html", handler.IndexHandler)
		v.GET("/add.html", handler.AddHandler)
		v.POST("/postme.html", handler.PostmeHandler)
	}
```

## 三：中间件的使用，在API中可能使用限流，身份验证等


	

``` go
	// 中间件 Go语言的net/http包特别容易构建中间件。
	// Gin提供了类似的中间件。需要注意的是中间件只对注册过的路由起作用。
	// 可以限定中间件的作用范围。大致分为全局中间件，单个处理程序中间件和组中间件。

	// 使用全局CORS中间件。
	// router.Use(Handler.Cors())
	// 即使是全局中间件，在use前的代码不受影响
	// 也可在handler中局部使用，见api.GetUser

	// 身份认证中间件，对于API，我们可以考虑JSON web tokens

	//rate-limit 限流中间件 
	lmt := tollbooth.NewLimiter(1, nil)
	//lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	//lmt.SetHeader("X-Access-Token", []string{"abc123", "xyz098"})
	//lmt.SetBasicAuthUsers([]string{"sansa"})
	lmt.SetMessage("服务繁忙，请稍后再试...")
	//tollbooth.LimitByKeys(lmt, []string{"127.0.0.1", "/"})
```

## 四：RESTful API接口

	

``` go
// 组路由以及版本控制
	v1 := router.Group("/v1")
	{
		// 下面是组中间件的用法
		// v1.Use(Handler.Cors())

		// 单个中间件的用法
		// v1.GET("/user/:id/*action",Handler.Cors(), api.GetUser)

		// rate-limit
		v1.GET("/user/:id/*action", LimitHandler(lmt), api.GetUser)

		//v1.GET("/user/:id/*action", Handler.Cors(), api.GetUser)
		// AJAX OPTIONS ，下面是有关OPTIONS用法的示例
		// v1.OPTIONS("/users", OptionsUser)      // POST
		// v1.OPTIONS("/users/:id", OptionsUser)  // PUT, DELETE
		/*
			// 对应的handler中增加处理
			func OptionsUser(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
			    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			    c.Next()
				...
			}

		*/
	}
```


## 五：服务器启动相关

	

``` go
	router := gin.Default()

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
```

