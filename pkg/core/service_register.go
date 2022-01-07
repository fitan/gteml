package core

import (
	"github.com/fitan/magic/pkg/types"
	"github.com/fitan/magic/services"
	core_oam_dev "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev"
	kubevelaapistandard "github.com/oam-dev/kubevela-core-api/apis/standard.oam.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	scheme2 "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
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
	k8sConfig     *rest.Config
	reload        bool
}

func NewServiceRegister() *ServiceRegister {
	return &ServiceRegister{reload: true}
}

func (s *ServiceRegister) Get() *ServiceRegister {
	if s.reload {
		var err error
		s.k8sConfig, err = clientcmd.BuildConfigFromFlags("", ConfReg.Confer.GetMyConf().K8sConf.ConfigPath)
		if err != nil {
			log.Panicln(err)
		}
		s.k8sConfig.APIPath = "/api"
		s.k8sConfig.GroupVersion = &schema.GroupVersion{Version: "v1"}
		s.k8sConfig.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme2.Codecs}
		k8sClient, err := kubernetes.NewForConfig(s.k8sConfig)
		if err != nil {
			log.Panicln(err)
		}
		s.k8sClient = k8sClient

		runtimeClient, err := client.New(s.k8sConfig, client.Options{Scheme: scheme})
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
	c.Services = services.NewServices(c, nil, s.Get().k8sClient, s.Get().runtimeClient, s.Get().k8sConfig)
}

func (s *ServiceRegister) Unset(c *types.Core) {
	return
}
