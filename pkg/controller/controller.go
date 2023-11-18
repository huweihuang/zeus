package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/huweihuang/golib/kube"
	log "github.com/huweihuang/golib/logger/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"

	"github.com/huweihuang/zeus/pkg/constant"
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
func (c *WorkerController) Run(ctx context.Context, workers int) {
	defer c.queue.ShutDown()

	for i := 0; i < workers; i++ {
		go wait.Forever(c.worker, time.Second)
	}
	log.Logger().Infof("worker controller is running, workers: [%d]", workers)

	<-ctx.Done()
}

func (c *WorkerController) worker() {
	for c.processNextWorkItem() {
	}
}

func (c *WorkerController) processNextWorkItem() bool {
	key, quit := c.queue.Get()
	log.Logger().With("key", key).With("quit", quit).Debug("processNextWorkItem begin")
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.syncHandler(key.(string))
	c.handleErr(err, key)
	return true
}

func (c *WorkerController) syncHandler(job string) (err error) {
	ins, err := ConvertJobToInstance(job)
	if err != nil {
		return err
	}

	switch ins.Status.JobState {
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
		log.Logger().With("err", err).With("job", job).Info("Error syncing job")
		c.queue.AddRateLimited(job)
		return
	}

	// 超过重试次数则丢弃任务，打印错误日志
	log.Logger().With("err", err).With("job", job).Info("Dropping job out of the queue")
	c.queue.Forget(job)
}

func (c *WorkerController) enqueue(ins *types.Instance) error {
	job, err := ConvertInstanceToJob(ins)
	if err != nil {
		return err
	}
	c.queue.Add(job)
	return nil
}

func (c *WorkerController) enqueueAfter(ins *types.Instance, after time.Duration) error {
	job, err := ConvertInstanceToJob(ins)
	if err != nil {
		return err
	}
	c.queue.AddAfter(job, after)
	return nil
}

func (c *WorkerController) enqueueRateLimited(ins *types.Instance) error {
	job, err := ConvertInstanceToJob(ins)
	if err != nil {
		return err
	}
	c.queue.AddRateLimited(job)
	return nil
}

func ConvertInstanceToJob(ins *types.Instance) (job string, err error) {
	jobByte, err := json.Marshal(ins)
	if err != nil {
		return "", fmt.Errorf("json marshal error, %v", err)
	}
	job = string(jobByte)
	return job, nil
}

func ConvertJobToInstance(job string) (ins *types.Instance, err error) {
	err = json.Unmarshal([]byte(job), &ins)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error, %v", err)
	}
	return ins, nil
}
