package controller

import (
	"context"
	"fmt"
	"time"

	rgp_client "github.com/apaarshrm39/rgp/pkg/client/clientset/versioned"
	rgp_informer "github.com/apaarshrm39/rgp/pkg/client/informers/externalversions/apaar.dev/v1alpha1"
	rgp_lister "github.com/apaarshrm39/rgp/pkg/client/listers/apaar.dev/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	// clientset for custom resource
	rclient rgp_client.Clientset
	// rgp cahce has synced ?
	hasSynced cache.InformerSynced
	// workqueue
	queue workqueue.RateLimitingInterface
	// lister for RGP
	rgpLister rgp_lister.RgpLister
}

func NewController(rclient rgp_client.Clientset, rinfiormer rgp_informer.RgpInformer) *Controller {
	c := &Controller{
		rclient:   rclient,
		rgpLister: rinfiormer.Lister(),
		hasSynced: rinfiormer.Informer().HasSynced,
		queue:     workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "queue"),
	}

	rinfiormer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.handleAdd,
			DeleteFunc: c.handleDelete,
		},
	)

	return c
}

func (c *Controller) Run(ch <-chan struct{}) {
	fmt.Println("Controller has started")

	if !cache.WaitForCacheSync(ch, c.hasSynced) {
		fmt.Println("could not sync cache")
	}

	go wait.Until(c.worker, 1*time.Second, ch)

	<-ch

}

func (c *Controller) worker() {
	for c.process() {

	}
}

func (c *Controller) process() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Forget(item)
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		fmt.Println(err)
		return false
	}

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		fmt.Println(err)
		return false
	}

	rgp, err := c.rclient.ApaarV1alpha1().Rgps(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(rgp.Name)

	return true
}

func (c *Controller) handleAdd(obj interface{}) {
	fmt.Println("Add even called")
	c.queue.Add(obj)
}

func (c *Controller) handleDelete(obj interface{}) {
	fmt.Println("Add even called")
	c.queue.Add(obj)
}
