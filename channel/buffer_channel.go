package main

import (
	"fmt"
)

// 无缓存的 channel

func writeRoutine(test_chan chan int, value int) {

	test_chan <- value
}

func readRoutine(test_chan chan int) {

	<-test_chan

	return
}

func main() {

	c := make(chan int)

	x := 100

	//readRoutine(c)
	//go writeRoutine(c, x)
	//
	//writeRoutine(c, x)
	//go readRoutine(c)
	//
	//go readRoutine(c)
	//writeRoutine(c, x)

	go writeRoutine(c, x)
	readRoutine(c)

	fmt.Println(x)
}