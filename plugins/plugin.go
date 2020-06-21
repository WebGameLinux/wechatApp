package plugins

import (
		"github.com/micro/go-micro/v2/registry"
		"github.com/micro/go-micro/v2/registry/etcd"
		"os"
		"strings"
		"sync"
)

var (
		safeLock         sync.Once
		registryInstance registry.Registry
)

const (
		EnvRegistryAddrKey  = "registry_address"
		DefaultRegistryAddr = "127.0.0.1:2379"
)

// 初始相关
func init() {
		initRegistry()
}

// 获取注册中心
func GetRegistry() registry.Registry {
		if registryInstance == nil {
				initRegistry()
		}
		return registryInstance
}

// 初始化注册中心
func initRegistry() registry.Registry {
		if registryInstance != nil {
				return registryInstance
		}
		safeLock.Do(newRegistry)
		return registryInstance
}

// 创建注册中心
func newRegistry() {
		registryInstance = etcd.NewRegistry(initRegistryOption)
}

// 初始注册中心配置
func initRegistryOption(options *registry.Options) {
		var opt = getEnvRegistryOptions()
		if len(opt.Addrs) != 0 {
				options.Addrs = append(opt.Addrs, options.Addrs...)
		}
		if len(options.Addrs) == 0 {
				options.Addrs = []string{DefaultRegistryAddr}
		}
		options.Secure = opt.Secure
		if opt.Timeout != 0 {
				options.Timeout = opt.Timeout
		}
		if options.TLSConfig == nil && opt.TLSConfig != nil {
				options.TLSConfig = opt.TLSConfig
		}
		if options.Context == nil && opt.Context != nil {
				options.Context = opt.Context
		}
}

// 获取环境变量中的参数
func getEnvRegistryOptions() *registry.Options {
		var opt = new(registry.Options)
		addr := os.Getenv(EnvRegistryAddrKey)
		if addr != "" {
				opt.Addrs = strings.Split(addr, ",")
		}
		//@todo other params
		return opt
}
