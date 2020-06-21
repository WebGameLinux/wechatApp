module github.com/WebGameLinux/wechatApp

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/kataras/iris v11.1.1+incompatible
	github.com/micro/go-micro/v2 v2.9.0
	github.com/webGameLinux/kits v0.0.0-20200620181511-e2648daf6c7b // indirect
	google.golang.org/protobuf v1.24.0
)
