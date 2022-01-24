package micro

import (
	microConsul "github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/hashicorp/consul/api"
	"go-micro.dev/v4/registry"
	"sync"
)

var consulRegistry registry.Registry
var once sync.Once

func ConsulRegistry(addr string) registry.Registry {
	once.Do(
		func() {
			consulConf := api.DefaultConfig()
			consulConf.Address = addr
			consulRegistry = microConsul.NewRegistry(microConsul.Config(consulConf))
		})
	return consulRegistry
}
