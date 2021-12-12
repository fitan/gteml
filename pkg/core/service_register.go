package core

import (
	"flag"
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services"
	core_oam_dev "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev"
	kubevelaapistandard "github.com/oam-dev/kubevela-core-api/apis/standard.oam.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var scheme = runtime.NewScheme()

func init() {
	_ = core_oam_dev.AddToScheme(scheme)
	_ = kubevelaapistandard.AddToScheme(scheme)
}

type ServiceRegister struct {
	k8sClient     *kubernetes.Clientset
	runtimeClient client.Client
	reload        bool
}

func NewServiceRegister() *ServiceRegister {
	return &ServiceRegister{reload: true}
}

func (s *ServiceRegister) Get() *ServiceRegister {
	if s.reload {

		k8sconfig := flag.String("k8sconfig", "/root/.kube/config", "kubernetes config file path")
		flag.Parse()
		config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
		if err != nil {
			log.Panicln(err)
		}
		k8sClient, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Panicln(err)
		}
		s.k8sClient = k8sClient

		runtimeClient, err := client.New(config, client.Options{Scheme: scheme})
		if err != nil {
			log.Println(err)
		}

		s.runtimeClient = runtimeClient
	}

	s.reload = false
	//if s.enforver == nil {
	//	a, err := gormadapter.NewAdapter("mysql", ConfReg.Confer.GetMyConf().Mysql.Url)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//	e, err := casbin.NewEnforcer("./rbac_model.conf", a)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//
	//	s.enforver = e
	//}
	return s
}

func (s *ServiceRegister) With(o ...types.Option) types.Register {
	return s
}

func (s *ServiceRegister) Reload(c *types.Core) {
	s.reload = true
}

func (s *ServiceRegister) Set(c *types.Core) {
	c.Services = services.NewServices(c, nil, s.Get().k8sClient, s.Get().runtimeClient)
}

func (s *ServiceRegister) Unset(c *types.Core) {
	return
}
