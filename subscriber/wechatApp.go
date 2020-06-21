package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	wechatApp "github.com/WebGameLinux/wechatApp/proto/wechatApp"
)

type WechatApp struct{}

func (e *WechatApp) Handle(ctx context.Context, msg *wechatApp.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *wechatApp.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
