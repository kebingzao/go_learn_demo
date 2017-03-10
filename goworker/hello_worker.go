package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		// 这边如果不指定的话，那么跑goworker.exe 就要加 -queues 参数来指定要处理的队列
		//Queues:         []string{"myqueue", "myqueue2", "myqueue3"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("worker_Hello", Worker_hello)
	goworker.Register("worker_World", Worker_world)
}

func Worker_hello(queue string, args ...interface{}) error {
	fmt.Printf("Hello,  --> From %s, %v\n", queue, args)
	return nil
}

func Worker_world(queue string, args ...interface{}) error {
	fmt.Printf("world,  --> From %s, %v\n", queue, args)
	return nil
}