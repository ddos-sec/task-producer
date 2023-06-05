package task_producer

import (
	"sync"
	"sync/atomic"
)

// TaskProducerByCount 根据任务次数生产任务
type TaskProducerByCount struct {
	count        int
	makeTaskFunc MakeTaskFunc
	wg           *sync.WaitGroup
	isRunning    *atomic.Bool
}

var _ TaskProducer = &TaskProducerByCount{}

// NewTaskProducerByCount 根据任务个数生产任务
func NewTaskProducerByCount(count int, makeTaskFunc MakeTaskFunc) *TaskProducerByCount {
	return &TaskProducerByCount{
		count:        count,
		makeTaskFunc: makeTaskFunc,
		wg:           &sync.WaitGroup{},
		isRunning:    &atomic.Bool{},
	}
}

func (x *TaskProducerByCount) Run() error {

	// 参数检查
	if x.count <= 0 {
		return ErrCountInvalid
	} else if x.makeTaskFunc == nil {
		return ErrMakeTaskFuncInvalid
	}

	x.wg.Add(1)
	x.isRunning.Store(true)
	go func() {
		defer x.wg.Done()
		defer x.isRunning.Store(false)

		for i := 0; x.isRunning.Load() && i < x.count; i++ {
			x.makeTaskFunc()
		}
	}()

	return nil
}

func (x *TaskProducerByCount) Stop() error {
	x.isRunning.Store(false)
	return nil
}

func (x *TaskProducerByCount) Await() error {
	x.wg.Wait()
	return nil
}
