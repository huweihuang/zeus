package controller

import (
	"fmt"
	"time"

	"github.com/huweihuang/golib/kube"
	log "github.com/huweihuang/golib/logger/logrus"
	"github.com/huweihuang/zeus/pkg/constant"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"

	"github.com/huweihuang/zeus/pkg/types"
)

const (
	// maxRetries is the number of times a job will be retried before it is dropped out of the queue.
	maxRetries = 5
)

// WorkerController is responsible for synchronizing job objects stored
// in the system with actual running replica sets and pods.
type WorkerController struct {
	client clientset.Interface
	// Workers that need to be synced
	queue workqueue.RateLimitingInterface
}

// NewWorkerController creates a new WorkerController.
func NewWorkerController(kubeConfig string) (*WorkerController, error) {
	client, err := kube.NewKubeClient(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to new kube client, err: %v", err)
	}
	c := &WorkerController{
		client: client,
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "job"),
	}
	return c, nil
}

// Run begins watching and syncing.
func (c *WorkerController) Run(workers int) {
	defer c.queue.ShutDown()

	for i := 0; i < workers; i++ {
		go wait.Forever(func() {
			if err := c.syncWorker(); err != nil {
				log.Logger.WithError(err).Errorln("failed to sync worker controller")
			}
		}, time.Second)
	}
}

func (c *WorkerController) syncWorker() (err error) {
	job, quit := c.queue.Get()
	if quit {
		return fmt.Errorf("quit workqueue")
	}
	defer c.queue.Done(job)

	switch job.(types.Instance).Status.JobState {
	case constant.JobStateCreating:
		err = c.createWorker()
	case constant.JobStateUpdating:
		err = c.updateWorker()
	case constant.JobStateDeleting:
		err = c.deleteWorker()
	}

	c.handleErr(err, job)
	return nil
}

// handleErr 错误没有超过重试次数则重新入队列
func (c *WorkerController) handleErr(err error, job interface{}) {
	if err == nil {
		c.queue.Forget(job)
		return
	}

	// 没有超过重试次数，则重新入队列
	if c.queue.NumRequeues(job) < maxRetries {
		log.Logger.WithError(err).WithField("job", job).Info("Error syncing job")
		c.queue.AddRateLimited(job)
		return
	}

	// 超过重试次数则丢弃任务，打印错误日志
	log.Logger.WithError(err).WithField("job", job).Info("Dropping job out of the queue")
	c.queue.Forget(job)
}

func (c *WorkerController) enqueue(job interface{}) {
	c.queue.Add(job)
}

func (c *WorkerController) enqueueAfter(job interface{}, after time.Duration) {
	c.queue.AddAfter(job, after)
}

func (c *WorkerController) enqueueRateLimited(job interface{}) {
	c.queue.AddRateLimited(job)
}
