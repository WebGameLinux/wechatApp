
GOPATH:=$(shell go env GOPATH)
SERVER_VERSION="1.0.0"
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/wechatApp/wechatApp.proto
    

.PHONY: build
build: proto

	go build -o wechatAppRpc-service  rpc/*.go
	go build -o wechatAppWeb-service  web/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t wechatAppRpc-service:${SERVER_VERSION}
