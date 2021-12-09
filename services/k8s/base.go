package k8s

import (
	"github.com/fitan/magic/pkg/types"
	servicesTypes "github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	appv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"go.uber.org/zap"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8s struct {
	k8sClient     *kubernetes.Clientset
	runtimeClient client.Client
	core          types.ServiceCore
}

func NewK8s(k8sClient *kubernetes.Clientset, runtimeClient client.Client, core types.ServiceCore) *K8s {
	return &K8s{k8sClient: k8sClient, runtimeClient: runtimeClient, core: core}
}

func (k *K8s) CreateApp(request servicesTypes.CreateAppRequest) (err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.CreateApp")
	defer func() {
		log.Debug(
			"CreateAppMsg",
			zap.Any("methodIn", map[string]interface{}{"request" :request}),
			zap.Any("methodOut", map[string]interface{}{"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"request" :request}))
		}

		log.Sync()
	}()

	components := make([]common.ApplicationComponent, 0, len(request.Components))
	for _, v := range request.Components {
		components = append(components, common.ApplicationComponent{
			Name:       v.Name,
			Type:       v.Type,
			Properties: util.Object2RawExtension(map[string]interface{}{}),
		})
	}

	return k.runtimeClient.Create(k.core.GetTrace().Ctx(), &appv1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind:       "Application",
			APIVersion: "core.oam.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Namespace: request.Metadata.NameSpace,
			Name:      request.Metadata.Name,
		},
		Spec: appv1beta1.ApplicationSpec{
			Components: []common.ApplicationComponent{
				{
					Name: "web",
					Type: "webservice",
					Properties: util.Object2RawExtension(map[string]interface{}{
						"image": "nginx:1.14.0",
					}),
					Traits: []common.ApplicationTrait{
						{
							Type: "labels",
							Properties: util.Object2RawExtension(map[string]interface{}{
								"hello": "world",
							}),
						}},
				}},
		},
	})
}

func (k *K8s) GetApps(keys []servicesTypes.K8sKey) (res *appv1beta1.ApplicationList,err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetApps")
	defer func() {
		log.Debug(
			"GetAppsMsg",
			zap.Any("methodIn", map[string]interface{}{"keys" :keys}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"keys" :keys}))
		}

		log.Sync()
	}()
	apps := new(appv1beta1.ApplicationList)
	opts := make([]client.ListOption,len(keys))
	for _, v := range keys {
		opt := &client.ListOptions{
			LabelSelector: client.MatchingLabelsSelector{
				Selector: labels.SelectorFromSet(labels.Set{oam.LabelAppName: v.Name}),
			},
			Namespace:     v.Namespace,
		}
		opts = append(opts,opt)
	}

	err = k.runtimeClient.List(k.core.GetTrace().Ctx(), apps, opts...)
	return apps, err
}

func (k *K8s) GetApp(key servicesTypes.K8sKey) (res *appv1beta1.Application, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetApp")
	defer func() {
		log.Debug(
			"GetAppMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key}))
		}

		log.Sync()
	}()
	app := &appv1beta1.Application{}
	err = k.runtimeClient.Get(
		k.core.GetTrace().Ctx(), client.ObjectKey{
			Namespace: key.Namespace,
			Name:      key.Name,
		}, app,
	)
	return app, err
}

func (k *K8s) GetPodsByDeploymentKey(key servicesTypes.K8sKey) (pods *v12.PodList, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetPodsByDeploymentKey")
	defer func() {
		log.Debug(
			"GetPodsByDeploymentKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key}),
			zap.Any("methodOut", map[string]interface{}{"pods" :pods,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key}))
		}

		log.Sync()
	}()
	deployment, err := k.GetDeployment(key)
	if err != nil {
		return nil, err
	}
	return k.k8sClient.CoreV1().Pods(key.Namespace).List(k.core.GetTrace().Ctx(), v1.ListOptions{
		LabelSelector: deployment.Spec.Selector.String(),
	})
}

func (k *K8s) GetPodByKey(key servicesTypes.K8sKey) (res *v12.Pod, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetPodByKey")
	defer func() {
		log.Debug(
			"GetPodByKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().Pods(key.Namespace).Get(k.core.GetTrace().Ctx(), key.Name, v1.GetOptions{})
}

func (k *K8s) CreateConfMap(key servicesTypes.K8sKey, data map[string]string) (res *v12.ConfigMap, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.CreateConfMap")
	defer func() {
		log.Debug(
			"CreateConfMapMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key,"data" :data}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key,"data" :data}))
		}

		log.Sync()
	}()
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(key.Namespace).Create(ctx, &v12.ConfigMap{
		ObjectMeta: v1.ObjectMeta{Name: key.Name},
		Data:       data,
	}, v1.CreateOptions{})
}

func (k *K8s) GetConfigMapByKey(key servicesTypes.K8sKey) (res *v12.ConfigMap, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetConfigMapByKey")
	defer func() {
		log.Debug(
			"GetConfigMapByKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key}))
		}

		log.Sync()
	}()
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(key.Namespace).Get(ctx, key.Name, v1.GetOptions{})
}

func (k *K8s) GetConfigMapsByDeploymentKey(key servicesTypes.K8sKey) (res *v12.ConfigMapList, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetConfigMapsByDeploymentKey")
	defer func() {
		log.Debug(
			"GetConfigMapsByDeploymentKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	deployment, err := k.GetDeployment(key)
	if err != nil {
		return nil, err
	}
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(key.Namespace).List(ctx, v1.ListOptions{
		LabelSelector: deployment.Spec.Selector.String(),
	})
}

func (k *K8s) GetDeployment(key servicesTypes.K8sKey) (res *v13.Deployment, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetDeployment")
	defer func() {
		log.Debug(
			"GetDeploymentMsg",
			zap.Any("methodIn", map[string]interface{}{"key" :key}),
			zap.Any("methodOut", map[string]interface{}{"res" :res,"err" :err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key" :key}))
		}

		log.Sync()
	}()
	return k.k8sClient.AppsV1().Deployments(key.Namespace).Get(k.core.GetTrace().Ctx(), key.Name, v1.GetOptions{})
}

func (k *K8s) CreatePv(key servicesTypes.K8sKey, )  {
	k.k8sClient.CoreV1().PersistentVolumes().Create(k.core.GetTrace().Ctx(), &v12.PersistentVolume{
		ObjectMeta: v1.ObjectMeta{Namespace: key.Namespace, Name: key.Name},
		Spec:       v12.PersistentVolumeSpec{
			Capacity:                      nil,
			AccessModes:                   nil,
			StorageClassName:              "",
			MountOptions:                  nil,
			VolumeMode:                    nil,
			NodeAffinity:                  nil,
		},
	}, v1.CreateOptions{})
}

func (k *K8s) GetPvs {
	
}
