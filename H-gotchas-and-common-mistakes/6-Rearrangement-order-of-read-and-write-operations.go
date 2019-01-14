package main

/*
读写操作的重排顺序

Go可能会对某些操作进行重新排序，但它能保证在一个goroutine内的所有行为顺序是不变的。然而，它并不保证多goroutine的执行顺序。
*/

import (
	"runtime"
	"time"
)
var _ = runtime.GOMAXPROCS(3)
var a, b int
func u1() {
	a = 1
	b = 2
}
func u2() {
	a = 3
	b = 4
}
func p() {
	println(a)
	println(b)
}
func main() {
	go u1()
	go u2()
	go p()
	time.Sleep(1 * time.Second)
}


// a和b最有趣的组合式是"02"。这表明b在a之前更新了。
//如果你需要在多goroutine内放置读写顺序的变化，你将需要使用channel，或者使用"sync"包构建合适的结构体。



