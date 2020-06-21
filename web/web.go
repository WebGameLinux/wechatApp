package main

import (
		"github.com/WebGameLinux/wechatApp/plugins"
		"github.com/kataras/iris"
		"github.com/kataras/iris/context"
		log "github.com/micro/go-micro/v2/logger"
		"github.com/micro/go-micro/v2/web"
)

const (
		ServiceVersion = "1.0.0"
		ServiceName    = "WeChat.service.wechatAppWeb"
)

func main() {
		// New Service

		service := web.NewService(
				web.Name(ServiceName),
				web.Version(ServiceVersion),
				web.Registry(plugins.GetRegistry()),
				web.Address(":18088"),
		)

		// Initialise service #接收命令行参数
		err := service.Init()
		if err != nil {
				log.Error(err)
		}

		app := iris.New()

		app.Get("/test", func(ctx context.Context) {
				ctx.JSON(iris.Map{
						"code": 0,
						"data": ctx.URLParams(),
						"msg":  "测试",
				})
		})
		_ = app.Build()

		service.Handle("/", app)
		// Run service
		if err = service.Run(); err != nil {
				log.Fatal(err)
		}
}
