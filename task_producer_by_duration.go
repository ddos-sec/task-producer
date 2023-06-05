package task_producer

import (
	"sync"
	"sync/atomic"
	"time"
)

// TaskProducerByDuration 根据持续时间来生产任务
type TaskProducerByDuration struct {
	duration     time.Duration
	makeTaskFunc MakeTaskFunc
	wg           sync.WaitGroup
	isRunning    atomic.Bool
}

var _ TaskProducer = &TaskProducerByDuration{}

func NewTaskProducerByDuration(duration time.Duration, makeTaskFunc MakeTaskFunc) *TaskProducerByDuration {
	return &TaskProducerByDuration{
		duration:     duration,
		makeTaskFunc: makeTaskFunc,
		wg:           sync.WaitGroup{},
		isRunning:    atomic.Bool{},
	}
}

func (x *TaskProducerByDuration) Run() error {

	// 参数检查
	if x.duration == 0 {
		return ErrDurationInvalid
	} else if x.makeTaskFunc == nil {
		return ErrMakeTaskFuncInvalid
	}

	x.wg.Add(1)
	x.isRunning.Store(true)
	go func() {
		defer x.wg.Done()
		defer x.isRunning.Store(false)

		begin := time.Now()

		for x.isRunning.Load() && time.Now().Sub(begin) < x.duration {
			x.makeTaskFunc()
		}
	}()

	return nil
}

func (x *TaskProducerByDuration) Stop() error {
	x.isRunning.Store(false)
	return nil
}

func (x *TaskProducerByDuration) Await() error {
	x.wg.Wait()
	return nil
}
