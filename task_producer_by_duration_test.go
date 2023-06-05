package task_producer

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewTaskProducerByDuration(t *testing.T) {

	c := atomic.Int64{}
	producer := NewTaskProducerByDuration(time.Second*3, func() {
		c.Add(1)
	})

	err := producer.Run()
	assert.Nil(t, err)

	err = producer.Await()
	assert.Nil(t, err)

	t.Log(c.Load())
	assert.True(t, c.Load() > 10000)
}

func TestTaskProducerByDuration_Stop(t *testing.T) {
	c := atomic.Int64{}
	producer := NewTaskProducerByDuration(time.Second*10, func() {
		c.Add(1)
	})
	err := producer.Run()
	assert.Nil(t, err)

	time.Sleep(time.Millisecond * 10)

	err = producer.Stop()
	assert.Nil(t, err)

	err = producer.Await()
	assert.Nil(t, err)

	assert.True(t, c.Load() > 10000)
}
