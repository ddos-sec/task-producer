package main

import (
	"fmt"
	task_producer "github.com/ddos-sec/task-producer"
	"sync/atomic"
)

func main() {
	c := &atomic.Int64{}
	producer := task_producer.NewTaskProducerByCount(10000, func() {
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

	fmt.Println(c.Load()) // Output: 10000
}
