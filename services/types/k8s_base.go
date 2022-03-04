package types

import (
	"bytes"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"io"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
)

//go:generate gowrap.exe gen -p ./ -i K8s -t ../../gowrap/templates/log -o k8s_with_log.go
//go:generate gowrap.exe gen -p ./ -i K8s -t ../../gowrap/templates/opentracing -o k8s_with_opentracing.go
//go:generate gowrap.exe gen -p ./ -i K8s -t ../../gowrap/templates/prometheus -o k8s_with_prometheus.go
//go:generate gowrap.exe gen -p ./ -i K8s -t ../../gowrap/templates/timeout -o k8s_with_timeout.go
type K8s interface {
	ApplyWorker(worker *Worker) (err error)
	CreateApp(request CreateAppRequest) (err error)
	UpdateApp(app *v1beta1.Application) (err error)
	GetApps(keys []K8sKey) (res *v1beta1.ApplicationList, err error)
	GetApp(key K8sKey) (res *v1beta1.Application, err error)
	DeleteApp(key K8sKey) (err error)
	GetPodsByDeploymentKey(key K8sKey) (pods *v1.PodList, err error)
	GetPodByKey(key K8sKey) (res *v1.Pod, err error)
	DeletePodByKey(key K8sKey) (err error)
	CreateConfMap(key K8sKey, data map[string]string) (err error)
	GetConfigMapByKey(key K8sKey) (res *v1.ConfigMap, err error)
	GetConfigMapsByDeploymentKey(key K8sKey) (res *v1.ConfigMapList, err error)
	GetDeployment(key K8sKey) (res *v12.Deployment, err error)
	CreatePvc(
		key K8sKey, volumeName, storageClassName, limits, requests string,
	) (res *v1.PersistentVolumeClaim, err error)
	UpdatePvc(pvc *v1.PersistentVolumeClaim) (claim *v1.PersistentVolumeClaim, err error)
	GetPvc(key K8sKey) (res *v1.PersistentVolumeClaim, err error)
	DeletePvc(key K8sKey) (err error)
	WatchPodLogs(key K8sKey, podName, containerName string, c chan string) error
	Exec(key K8sKey, podName, containerName string, cmd []string, runError *string) *io.PipeReader

	PodCopyFile(src string, dest string, containername string) (in *bytes.Buffer, out *bytes.Buffer, errOut *bytes.Buffer, err error)
	PortForward(key K8sKey, podName string, ports []string, down <-chan struct{}) error
	PodCopyFileV2(key K8sKey, containerName string, src string) (
		*io.PipeReader, error,
	)
	GetPodLogs(key K8sKey, podName, containerName string, tailLines *int64) (string, error)
}
