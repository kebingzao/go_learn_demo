package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		Queues:         []string{"myqueue", "delimited", "queues"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "resque:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("Hello", helloWorker)
}

func helloWorker(queue string, args ...interface{}) error {
	fmt.Printf("Hello, world! --> From %s, %v\n", queue, args)
	return nil
}