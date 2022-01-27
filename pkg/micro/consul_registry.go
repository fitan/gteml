package micro

import (
	microConsul "github.com/asim/go-micro/plugins/registry/consul/v4"
	microetcd "github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/hashicorp/consul/api"
	"go-micro.dev/v4/registry"
	"sync"
)

var consulRegistry registry.Registry
var consulOnce sync.Once

func ConsulRegistry(addr string) registry.Registry {
	consulOnce.Do(
		func() {
			consulConf := api.DefaultConfig()
			consulConf.Address = addr
			consulRegistry = microConsul.NewRegistry(microConsul.Config(consulConf))
		})
	return consulRegistry
}

var etcdRegistry registry.Registry
var etcdOnce sync.Once

func EtcdRegistry(addr string) registry.Registry {
	etcdOnce.Do(
		func() {
			etcdRegistry = microetcd.NewRegistry(registry.Addrs(addr))
		})
	return etcdRegistry
}
