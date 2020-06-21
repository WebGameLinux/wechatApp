package main

import (
		"github.com/WebGameLinux/wechatApp/handler"
		"github.com/WebGameLinux/wechatApp/plugins"
		wechatApp "github.com/WebGameLinux/wechatApp/proto/wechatApp"
		"github.com/WebGameLinux/wechatApp/subscriber"
		"github.com/micro/go-micro/v2"
		log "github.com/micro/go-micro/v2/logger"
)

const (
		ServiceVersion = "1.0.0"
		ServiceName    = "WeChat.service.wechatAppRpc"
)

func main() {

		// New Service
		service := micro.NewService(
				micro.Name(ServiceName),
				micro.Version(ServiceVersion),
				micro.Address("127.0.0.1:8888"),
				micro.Registry(plugins.GetRegistry()),
		)

		// Initialise service #接收命令行参数
		service.Init()

		// Register Handler
		err := wechatApp.RegisterWechatAppHandler(service.Server(), new(handler.WechatApp))
		if err != nil {
				log.Error(err)
		}

		// Register Struct as Subscriber # 消息订阅
		err = micro.RegisterSubscriber(ServiceName, service.Server(), new(subscriber.WechatApp))
		if err != nil {
				log.Error(err)
		}


		// Run service
		if err = service.Run(); err != nil {
				log.Fatal(err)
		}
}
