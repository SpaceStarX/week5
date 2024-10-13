package main

import (
	"flag"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	workqueue "k8s.io/client-go/util/workqueue"
)

type Controller struct {
	indexer cache.Indexer
	queue workqueue.TypedRateLimitingInterface[string]
	informer cache.Controller
}

func NewController(queue workqueue.TypedRateLimitingInterface[string], indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		informer: informer,
		indexer: indexer,
		queue: queue,
	}
}

func (c *Controller) syncToStdout(key string) error{
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Printf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		fmt.Printf("Pod %s does not exist anymore\n", key)
	} else {
		pod := obj.(*corev1.Pod)
		fmt.Printf("Sync/Add/Update pod %s, status: %s\n", pod.Name, pod.Status.Phase)
	}
	return nil
}

func (c *Controller) handleErr(err error, key string){
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		fmt.Printf("Retry %d for key %s", c.queue.NumRequeues(key), key)
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	fmt.Printf("Dropping pod %q out of the queue: %v\n", key, err)
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	
	err := c.syncToStdout(key)
	c.handleErr(err, key)
	return true
}

func onAddPod(obj interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err == nil {
		queue.Add(key)
	}
}

func onUpdatePod(obj interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err == nil {
		queue.Add(key)
	}
}

func onDeletePod(obj interface{}, queue workqueue.TypedRateLimitingInterface[string]) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err == nil {
		queue.Add(key)
	}
}

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/ymin/.kube/docker-desktop", "Path to kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Hour * 2)
	queue := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[string]())

	deployInformer := informerFactory.Core().V1().Pods()
	informer := deployInformer.Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			onAddPod(obj, queue)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			onUpdatePod(newObj, queue)
		},
		DeleteFunc: func(obj interface{}) {
			onDeletePod(obj, queue)
		},
	})
	
	stopper := make(chan struct{})
	defer close(stopper)

	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)

	controller := NewController(queue, deployInformer.Informer().GetIndexer(), informer)	
	go func(){
		for {
			if !controller.processNextItem() {
				break
			}
		}
	}()

	<- stopper
}