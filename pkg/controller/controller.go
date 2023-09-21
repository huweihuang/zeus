package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	logger "k8s.io/klog/v2"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	// maxRetries is the number of times a deployment will be retried before it is dropped out of the queue.
	maxRetries = 15
)

// DeploymentController is responsible for synchronizing Deployment objects stored
// in the system with actual running replica sets and pods.
type DeploymentController struct {
	client clientset.Interface
	// Deployments that need to be synced
	queue workqueue.RateLimitingInterface
}

// NewDeploymentController creates a new DeploymentController.
func NewDeploymentController(ctx context.Context, client clientset.Interface) (*DeploymentController, error) {
	dc := &DeploymentController{
		client: client,
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "deployment"),
	}
	return dc, nil
}

// Run begins watching and syncing.
func (dc *DeploymentController) Run(ctx context.Context, workers int) {
	defer dc.queue.ShutDown()

	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, dc.worker, time.Second)
	}

	<-ctx.Done()
}

// worker runs a worker thread that just dequeues items, processes them, and marks them done.
// It enforces that the syncHandler is never invoked concurrently with the same key.
func (dc *DeploymentController) worker(ctx context.Context) {
	for dc.processNextWorkItem(ctx) {
	}
}

func (dc *DeploymentController) processNextWorkItem(ctx context.Context) bool {
	key, quit := dc.queue.Get()
	if quit {
		return false
	}
	defer dc.queue.Done(key)

	err := dc.syncHandler(ctx, key.(string))
	dc.handleErr(ctx, err, key)

	return true
}

// syncHandler will sync the deployment with the given key.
// This function is not meant to be invoked concurrently with the same key.
func (dc *DeploymentController) syncHandler(ctx context.Context, key string) error {
	return nil
}

func (dc *DeploymentController) handleErr(ctx context.Context, err error, key interface{}) {
	if err == nil {
		dc.queue.Forget(key)
		return
	}
	ns, name, keyErr := cache.SplitMetaNamespaceKey(key.(string))
	if keyErr != nil {
		logger.Error(err, "Failed to split meta namespace cache key", "cacheKey", key)
	}

	if dc.queue.NumRequeues(key) < maxRetries {
		logger.V(2).Info("Error syncing deployment", "deployment", logger.KRef(ns, name), "err", err)
		dc.queue.AddRateLimited(key)
		return
	}

	logger.V(2).Info("Dropping deployment out of the queue", "deployment", logger.KRef(ns, name), "err", err)
	dc.queue.Forget(key)
}

func (dc *DeploymentController) enqueue(deployment interface{}) {
	key := ""
	dc.queue.Add(key)
}
