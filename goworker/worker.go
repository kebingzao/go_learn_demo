package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

// run=====
// go build
// goworker.exe -queues=myqueue,myqueue2

func main() {
	fmt.Println("===========start=========")
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}