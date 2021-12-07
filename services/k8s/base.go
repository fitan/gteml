package k8s

import (
	"github.com/fitan/magic/pkg/types"
	servicesTypes "github.com/fitan/magic/services/types"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	appv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"go.uber.org/zap"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8s struct {
	k8sClient     kubernetes.Clientset
	runtimeClient client.Client
	core          types.ServiceCore
}

func (k *K8s) CreateApp(request servicesTypes.CreateAppRequest) (err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.CreateApp")
	defer func() {
		log.Debug(
			"CreateAppMsg",
			zap.Any("methodIn", map[string]interface{}{"request": request}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"request": request}))
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

func (k *K8s) GetApp(namespace, name string) (app *appv1beta1.Application, err error) {
	err = k.runtimeClient.Get(
		k.core.GetTrace().Ctx(), client.ObjectKey{
			Namespace: namespace,
			Name:      name,
		}, app,
	)
	return
}

func (k *K8s) GetPodsByDeploymentName(namespace string, deploymentName string) (pods *v12.PodList, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetPodsByDeploymentName")
	defer func() {
		log.Debug(
			"GetPodsByDeploymentNameMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "deploymentName": deploymentName}),
			zap.Any("methodOut", map[string]interface{}{"pods": pods, "err": err}),
		)

		if err != nil {
			log.Error(
				err.Error(),
				zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "deploymentName": deploymentName}),
			)
		}

		log.Sync()
	}()
	deployment, err := k.GetDeployment(namespace, deploymentName)
	if err != nil {
		return nil, err
	}
	deployment.Spec.Selector.String()
	return k.k8sClient.CoreV1().Pods(namespace).List(k.core.GetTrace().Ctx(), v1.ListOptions{
		LabelSelector: deployment.Spec.Selector.String(),
	})
}

func (k *K8s) GetPod(namespace string, name string) (res *v12.Pod, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetPod")
	defer func() {
		log.Debug(
			"GetPodMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().Pods(namespace).Get(k.core.GetTrace().Ctx(), name, v1.GetOptions{})
}

func (k *K8s) CreateConfMap(namespace string, name string, data map[string]string) (res *v12.ConfigMap, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.CreateConfMap")
	defer func() {
		log.Debug(
			"CreateConfMapMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name, "data": data}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(
				err.Error(),
				zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name, "data": data}),
			)
		}

		log.Sync()
	}()

	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(namespace).Create(ctx, &v12.ConfigMap{
		ObjectMeta: v1.ObjectMeta{Name: name},
		Data:       data,
	}, v1.CreateOptions{})
}

func (k *K8s) GetConfigMap(namespace, name string) (res *v12.ConfigMap, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetConfigMap")
	defer func() {
		log.Debug(
			"GetConfigMapMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}))
		}

		log.Sync()
	}()
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(namespace).Get(ctx, name, v1.GetOptions{})
}

func (k *K8s) GetConfigMapsByDeployment(namespace string, deploymentName string) (res *v12.ConfigMapList, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetConfMaps")
	defer func() {
		log.Debug(
			"GetConfMapsMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "deploymentName": deploymentName}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"namespace": namespace}))
		}

		log.Sync()
	}()
	deployment, err := k.GetDeployment(namespace, deploymentName)
	if err != nil {
		return nil, err
	}
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(namespace).List(ctx, v1.ListOptions{
		LabelSelector: deployment.Spec.Selector.String(),
	})
}

func (k *K8s) GetDeployment(namespace, name string) (res *v13.Deployment, err error) {
	log := k.core.GetCoreLog().ApmLog("services.k8s.GetDeployment")
	defer func() {
		log.Debug(
			"GetDeploymentMsg",
			zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"namespace": namespace, "name": name}))
		}

		log.Sync()
	}()

	return k.k8sClient.AppsV1().Deployments(namespace).Get(k.core.GetTrace().Ctx(), name, v1.GetOptions{})
}
