package main

import (
	"fmt"
	"github.com/benmanns/goworker"
)

// run=====
// go build
// goworker.exe -queues=hello

func main() {
	fmt.Println("===========start=========")
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}