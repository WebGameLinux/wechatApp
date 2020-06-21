package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	wechatApp "github.com/WebGameLinux/wechatApp/proto/wechatApp"
)

type WechatApp struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *WechatApp) Call(ctx context.Context, req *wechatApp.Request, rsp *wechatApp.Response) error {
	log.Info("Received WechatApp.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *WechatApp) Stream(ctx context.Context, req *wechatApp.StreamingRequest, stream wechatApp.WechatApp_StreamStream) error {
	log.Infof("Received WechatApp.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&wechatApp.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *WechatApp) PingPong(ctx context.Context, stream wechatApp.WechatApp_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&wechatApp.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
