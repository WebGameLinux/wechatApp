# WechatApp Service

This is the WechatApp service

Generated with

```
micro new --namespace=WeChat --type=service github.com/WebGameLinux/wechatApp
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: WeChat.service.wechatApp
- Type: service
- Alias: wechatApp

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./wechatApp-service
```

Build a docker image
```
make docker
```