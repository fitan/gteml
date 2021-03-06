package k8s

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fitan/magic/pkg/types"
	servicesTypes "github.com/fitan/magic/services/types"
	"github.com/gorilla/websocket"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	appv1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/kubectl/pkg/cmd/cp"
	"k8s.io/kubectl/pkg/cmd/exec"
	"net/http"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type K8s struct {
	k8sClient       *kubernetes.Clientset
	runtimeClient   client.Client
	informerFactory informers.SharedInformerFactory
	cfg             *rest.Config
	core            types.ServiceCore
}

func NewK8s(k8sClient *kubernetes.Clientset, runtimeClient client.Client, cfg *rest.Config, core types.ServiceCore) *K8s {
	return &K8s{
		k8sClient:     k8sClient,
		runtimeClient: runtimeClient,
		cfg:           cfg,
		core:          core,
	}
}

func (k *K8s) CreateWorker(worker *servicesTypes.Worker) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.CreateWorker")
	defer func() {
		log.Debug(
			"CreateWorkerMsg",
			zap.Any("methodIn", map[string]interface{}{"worker": worker}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"worker": worker}))
		}

		log.Sync()
	}()

	return k.runtimeClient.Create(k.core.GetTrace().Ctx(), worker.ToWorker())

}

func (k *K8s) UpdateWorker(worker *servicesTypes.Worker) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.UpdateWorker")
	defer func() {
		log.Debug(
			"UpdateWorkerMsg",
			zap.Any("methodIn", map[string]interface{}{"worker": worker}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"worker": worker}))
		}

		log.Sync()
	}()

	old, err := k.GetApp(worker.Metadata)
	if err != nil {
		return err
	}
	w := worker.ToWorker()
	w.SetResourceVersion(old.ResourceVersion)
	return k.runtimeClient.Update(k.core.GetTrace().Ctx(), w, &client.UpdateOptions{})

}

func (k *K8s) ApplyWorker(worker *servicesTypes.Worker) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.CreateWorker")
	defer func() {
		log.Debug(
			"CreateAppMsg",
			zap.Any("methodIn", map[string]interface{}{"worker": worker}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"worker": worker}))
		}

		log.Sync()
	}()
	//b, _ := json.Marshal(worker)
	//fmt.Println(string(b))
	//
	//w, _ := json.Marshal(worker.ToWorker())
	//fmt.Println(string(w))
	old, err := k.GetApp(worker.Metadata)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return k.runtimeClient.Create(k.core.GetTrace().Ctx(), worker.ToWorker())
		}
		return err
	}
	w := worker.ToWorker()
	w.SetResourceVersion(old.ResourceVersion)
	log.Info("ToWorker", zap.Any("worker", w))
	return k.runtimeClient.Update(k.core.GetTrace().Ctx(), w, &client.UpdateOptions{})
}

func (k *K8s) CreateApp(request servicesTypes.CreateAppRequest) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.CreateApp")
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

	return k.runtimeClient.Create(k.core.GetTrace().Ctx(), request.ToApplication())
}

func (k *K8s) UpdateApp(app *appv1beta1.Application) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.UpdateApp")
	defer func() {
		log.Debug(
			"UpdateAppMsg",
			zap.Any("methodIn", map[string]interface{}{"app": app}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"app": app}))
		}

		log.Sync()
	}()
	return k.runtimeClient.Update(k.core.GetTrace().Ctx(), app)
}

func (k *K8s) GetApps(keys []servicesTypes.K8sKey) (res *appv1beta1.ApplicationList, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetApps")
	defer func() {
		log.Debug(
			"GetAppsMsg",
			zap.Any("methodIn", map[string]interface{}{"keys": keys}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"keys": keys}))
		}

		log.Sync()
	}()
	apps := new(appv1beta1.ApplicationList)
	opts := make([]client.ListOption, len(keys))
	for _, v := range keys {
		opt := &client.ListOptions{
			LabelSelector: client.MatchingLabelsSelector{
				Selector: labels.SelectorFromSet(labels.Set{oam.LabelAppName: v.Name}),
			},
			Namespace: v.Namespace,
		}
		opts = append(opts, opt)
	}

	err = k.runtimeClient.List(k.core.GetTrace().Ctx(), apps, opts...)
	return apps, err
}

func (k *K8s) GetApp(key servicesTypes.K8sKey) (res *appv1beta1.Application, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetApp")
	defer func() {
		log.Debug(
			"GetAppMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
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

func (k *K8s) DeleteApp(key servicesTypes.K8sKey) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.DeleteApp")
	defer func() {
		log.Debug(
			"DeleteAppMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	app, err := k.GetApp(key)
	if err != nil {
		return err
	}
	return k.runtimeClient.Delete(k.core.GetTrace().Ctx(), app)
}

func (k *K8s) GetPodsByDeploymentKey(key servicesTypes.K8sKey) (pods *v12.PodList, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetPodsByDeploymentKey")
	defer func() {
		log.Debug(
			"GetPodsByDeploymentKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"pods": pods, "err": err}),
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

	asMap, err := v1.LabelSelectorAsMap(deployment.Spec.Selector)
	if err != nil {
		return nil, err
	}

	return k.k8sClient.CoreV1().Pods(key.Namespace).List(k.core.GetTrace().Ctx(), v1.ListOptions{
		LabelSelector: labels.SelectorFromSet(asMap).String(),
	})
}

func (k *K8s) GetPodByKey(key servicesTypes.K8sKey) (res *v12.Pod, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetPodByKey")
	defer func() {
		log.Debug(
			"GetPodByKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().Pods(key.Namespace).Get(k.core.GetTrace().Ctx(), key.Name, v1.GetOptions{})
}

func (k *K8s) DeletePodByKey(key servicesTypes.K8sKey) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.DeletePodByKey")
	defer func() {
		log.Debug(
			"DeletePodByKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().Pods(key.Namespace).Delete(k.core.GetTrace().Ctx(), key.Name, v1.DeleteOptions{})
}

func (k *K8s) CreateConfMap(key servicesTypes.K8sKey, data map[string]string) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.CreateConfMap")
	defer func() {
		log.Debug(
			"CreateConfMapMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key, "data": data}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key, "data": data}))
		}

		log.Sync()
	}()
	return k.runtimeClient.Create(k.core.GetTrace().Ctx(), &appv1beta1.Application{
		TypeMeta: v1.TypeMeta{
			Kind: servicesTypes.ApplicationKindName,
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      key.Name + "-ConfMap",
			Namespace: key.Namespace,
		},
		Spec: appv1beta1.ApplicationSpec{
			Components: []common.ApplicationComponent{common.ApplicationComponent{
				Name: key.Name + "-ConfMap",
				Type: "raw",
				Properties: util.Object2RawExtension(v12.ConfigMap{
					ObjectMeta: v1.ObjectMeta{Name: key.Name},
					Data:       data,
				}),
			}},
		},
	})
	//return k.k8sClient.CoreV1().ConfigMaps(key.Namespace).Create(k.core.GetTrace().Ctx(), &v12.ConfigMap{
	//	ObjectMeta: v1.ObjectMeta{Name: key.Name},
	//	Data:       data,
	//}, v1.CreateOptions{})
}

func (k *K8s) GetConfigMapByKey(key servicesTypes.K8sKey) (res *v12.ConfigMap, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetConfigMapByKey")
	defer func() {
		log.Debug(
			"GetConfigMapByKeyMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	ctx := k.core.GetTrace().Ctx()
	return k.k8sClient.CoreV1().ConfigMaps(key.Namespace).Get(ctx, key.Name, v1.GetOptions{})
}

func (k *K8s) GetConfigMapsByDeploymentKey(key servicesTypes.K8sKey) (res *v12.ConfigMapList, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetConfigMapsByDeploymentKey")
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
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetDeployment")
	defer func() {
		log.Debug(
			"GetDeploymentMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	return k.k8sClient.AppsV1().Deployments(key.Namespace).Get(k.core.GetTrace().Ctx(), key.Name, v1.GetOptions{})
}

func (k *K8s) CreatePvc(
	key servicesTypes.K8sKey, volumeName, storageClassName, limits, requests string,
) (res *v12.PersistentVolumeClaim, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.CreatePvc")
	defer func() {
		log.Debug(
			"CreatePvcMsg",
			zap.Any(
				"methodIn", map[string]interface{}{
					"key": key, "volumeName": volumeName, "storageClassName": storageClassName, "limits": limits, "requests": requests,
				},
			),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(
				err.Error(), zap.Any(
					"methodIn", map[string]interface{}{
						"key": key, "volumeName": volumeName, "storageClassName": storageClassName, "limits": limits, "requests": requests,
					},
				),
			)
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().PersistentVolumeClaims(key.Namespace).Create(k.core.GetTrace().Ctx(), &v12.PersistentVolumeClaim{
		ObjectMeta: v1.ObjectMeta{Name: key.Name},
		Spec: v12.PersistentVolumeClaimSpec{
			AccessModes: []v12.PersistentVolumeAccessMode{v12.ReadWriteMany},
			Resources: v12.ResourceRequirements{
				Limits:   v12.ResourceList{v12.ResourceStorage: resource.MustParse(limits)},
				Requests: v12.ResourceList{v12.ResourceStorage: resource.MustParse(requests)},
			},
			VolumeName:       volumeName,
			StorageClassName: &storageClassName,
		},
	}, v1.CreateOptions{})
}

func (k *K8s) UpdatePvc(pvc *v12.PersistentVolumeClaim) (claim *v12.PersistentVolumeClaim, err error) {
	return k.k8sClient.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(k.core.GetTrace().Ctx(), pvc, v1.UpdateOptions{})
}

func (k *K8s) GetPvc(key servicesTypes.K8sKey) (res *v12.PersistentVolumeClaim, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.GetPvc")
	defer func() {
		log.Debug(
			"GetPvcMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"res": res, "err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().PersistentVolumeClaims(key.Namespace).Get(k.core.GetTrace().Ctx(), key.Name, v1.GetOptions{})
}

func (k *K8s) DeletePvc(key servicesTypes.K8sKey) (err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.DeletePvc")
	defer func() {
		log.Debug(
			"DeletePvcMsg",
			zap.Any("methodIn", map[string]interface{}{"key": key}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(err.Error(), zap.Any("methodIn", map[string]interface{}{"key": key}))
		}

		log.Sync()
	}()
	return k.k8sClient.CoreV1().PersistentVolumeClaims(key.Namespace).Delete(k.core.GetTrace().Ctx(), key.Name, v1.DeleteOptions{})
}

func (k *K8s) GetPodLogs(key servicesTypes.K8sKey, podName, containerName string, tailLines *int64) (string, error) {
	req := k.k8sClient.CoreV1().Pods(key.Namespace).GetLogs(podName, &v12.PodLogOptions{
		Container: containerName,
		TailLines: tailLines,
		//Timestamps: true,
	})
	readCloser, err := req.Stream(k.core.GetTrace().Ctx())
	if err != nil {
		return "", err
	}

	defer readCloser.Close()

	b, err := ioutil.ReadAll(readCloser)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (k *K8s) WatchPodLogs(key servicesTypes.K8sKey, podName, containerName string, c chan string) error {
	var taillines int64 = 10
	//var byteReadLimit int64 = 500000
	req := k.k8sClient.CoreV1().Pods(key.Namespace).GetLogs(podName, &v12.PodLogOptions{
		Container: containerName,
		Follow:    true,
		//Timestamps: true,
		//LimitBytes: &byteReadLimit,
		TailLines: &taillines,
	})
	readCloser, err := req.Stream(k.core.GetTrace().Ctx())
	if err != nil {
		return err
	}

	defer readCloser.Close()
	defer close(c)
	r := bufio.NewReader(readCloser)
	for {
		bytes, err := r.ReadBytes('\n')
		c <- string(bytes)
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
	}
}

func (k *K8s) PodCopyFile(src string, dest string, containername string) (in *bytes.Buffer, out *bytes.Buffer, errOut *bytes.Buffer, err error) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.PodCopyFile")
	defer func() {
		log.Debug(
			"PodCopyFileMsg",
			zap.Any("methodIn", map[string]interface{}{"src": src, "dest": dest, "containername": containername}),
			zap.Any("methodOut", map[string]interface{}{"in": in, "out": out, "errOut": errOut, "err": err}),
		)

		if err != nil {
			log.Error(
				err.Error(),
				zap.Any("methodIn", map[string]interface{}{"src": src, "dest": dest, "containername": containername}),
			)
		}

		log.Sync()
	}()

	ioStreams, in, out, errOut := genericclioptions.NewTestIOStreams()
	copyOptions := cp.NewCopyOptions(ioStreams)
	copyOptions.Clientset = k.k8sClient
	copyOptions.ClientConfig = k.cfg
	copyOptions.Container = containername
	err = copyOptions.Run([]string{src, dest})
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Could not run copy operation: %v", err)
	}
	return in, out, errOut, nil
}

func (k *K8s) PodCopyFileV2(key servicesTypes.K8sKey, containerName string, src string) (
	res *io.PipeReader, err error,
) {
	log := k.core.GetCoreLog().TraceLog("services.k8s.PodCopyFileV2")
	defer func() {
		log.Debug(
			"PodCopyFileV2",
			zap.Any("methodIn", map[string]interface{}{"key": servicesTypes.K8sKey{}, "containername": containerName, "src": src}),
			zap.Any("methodOut", map[string]interface{}{"err": err}),
		)

		if err != nil {
			log.Error(
				err.Error(),
				zap.Any("methodIn", map[string]interface{}{"key": servicesTypes.K8sKey{}, "containername": containerName, "src": src}),
			)
		}

		log.Sync()
	}()
	reader, outStream := io.Pipe()
	cmdArr := []string{"tar", "cf", "-", src}
	req := k.k8sClient.RESTClient().Get().Namespace(key.Namespace).Resource("pods").Name(key.Name).SubResource("exec").VersionedParams(
		&v12.PodExecOptions{
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
			Container: containerName,
			Command:   cmdArr,
		}, v1.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(k.cfg, "POST", req.URL())
	if err != nil {
		log.Error("NewSPDYExecutor", zap.Error(err))
		return nil, err
	}
	go func() {
		defer outStream.Close()
		err = exec.Stream(remotecommand.StreamOptions{
			Stdin:  os.Stdin,
			Stdout: outStream,
			Stderr: os.Stderr,
			Tty:    false,
		})
		if err != nil {
			log.Error("exec.Stream error", zap.Error(err))
		}
	}()
	return reader, nil
}

func (k *K8s) PortForward(key servicesTypes.K8sKey, podName string, ports []string, down <-chan struct{}) error {
	req := k.k8sClient.CoreV1().RESTClient().Post().Namespace(key.Namespace).Resource("pods").Name(podName).SubResource("portforward")
	//signals := make(chan os.Signal, 1)
	//StopChannel := make(chan struct{}, 1)
	ReadyChannel := make(chan struct{})

	//defer signal.Stop(signals)

	transport, upgrader, err := spdy.RoundTripperFor(k.cfg)
	if err != nil {
		return err
	}

	stream := genericclioptions.NewTestIOStreamsDiscard()

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())
	fw, err := portforward.NewOnAddresses(dialer, []string{"0.0.0.0"}, ports, down, ReadyChannel, stream.Out, stream.ErrOut)
	if err != nil {
		return err
	}

	return fw.ForwardPorts()
}

func (k *K8s) Exec(
	key servicesTypes.K8sKey,
	podName, containerName string,
	cmd []string,
	runError *string,
) *io.PipeReader {

	reader, writer := io.Pipe()

	in := &bytes.Buffer{}
	errOut := &bytes.Buffer{}

	options := &exec.ExecOptions{
		StreamOptions: exec.StreamOptions{
			IOStreams: genericclioptions.IOStreams{
				In:     in,
				Out:    writer,
				ErrOut: errOut,
			},

			ContainerName: containerName,
			Namespace:     key.Namespace,
			PodName:       podName,
		},

		// TODO: Improve error messages by first testing if 'tar' is present in the container?
		Command:  cmd,
		Executor: &exec.DefaultRemoteExecutor{},
	}

	options.Config = k.cfg
	options.PodClient = k.k8sClient.CoreV1()
	go func() {
		err := options.Run()
		if err != nil {
			*runError = err.Error()
		}
		writer.Close()
	}()
	return reader
}

func (k *K8s) SSH(key servicesTypes.K8sKey, podName, containerName string, ws *websocket.Conn) error {
	defer ws.Close()
	req := k.k8sClient.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(key.Namespace).SubResource("exec")
	req.VersionedParams(&v12.PodExecOptions{
		TypeMeta: v1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
		Container: containerName,
		Command:   []string{"clear; (bash || ash || sh)"},
	}, v1.ParameterCodec)
	executor, err := remotecommand.NewSPDYExecutor(k.cfg, "POST", req.URL())
	if err != nil {
		return err
	}

	handler := &TerminalSession{
		wsConn:   ws,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		Tty:               true,
		TerminalSizeQueue: handler,
	})
	if err != nil {
		return err
	}
	return nil
}
