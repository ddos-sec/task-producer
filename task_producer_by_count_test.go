package task_producer

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewTaskProducerByCount(t *testing.T) {

	c := &atomic.Int64{}
	producer := NewTaskProducerByCount(10000, func() {
		c.Add(1)
	})

	err := producer.Run()
	assert.Nil(t, err)

	err = producer.Await()
	assert.Nil(t, err)

	assert.Equal(t, int64(10000), c.Load())
}

func TestTaskProducerByCount_Stop(t *testing.T) {

	c := atomic.Int64{}
	producer := NewTaskProducerByCount(0xFFFFFFFFFFFF, func() {
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
