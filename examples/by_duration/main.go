package main

import (
	"fmt"
	task_producer "github.com/ddos-sec/task-producer"
	"sync/atomic"
	"time"
)

func main() {
	c := atomic.Int64{}
	producer := task_producer.NewTaskProducerByDuration(time.Second*3, func() {
		c.Add(1)
	})

	err := producer.Run()
	if err != nil {
		panic(err)
	}

	err = producer.Await()
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Load()) // Output: 444424594
}
