package main

/*
类型声明和方法
当你通过把一个现有（非interface）的类型定义为一个新的类型时，新的类型不会继承现有类型的方法。
*/

// 失败的例子：

/*
import "sync"
type myMutex sync.Mutex
func main() {
	var mtx myMutex
	mtx.Lock() //error
	mtx.Unlock() //error
}
*/

// print::
/*
command-line-arguments
M-gotchas-and-common-mistakes\9-Type-declaration-and-method.go:14: mtx.Lock undefined (type myMutex has no field or method Lock)
M-gotchas-and-common-mistakes\9-Type-declaration-and-method.go:15: mtx.Unlock undefined (type myMutex has no field or method Unlock)
*/

// 如果你确实需要原有类型的方法，你可以定义一个新的struct类型，用匿名方式把原有类型嵌入其中。 比如：

/*
import "sync"
type myLocker struct {
	sync.Mutex
}
func main() {
	var lock myLocker
	lock.Lock() //ok
	lock.Unlock() //ok
}

*/

// 还有另一种方法， interface类型的声明也会保留它们的方法集合。
import "sync"
type myLocker sync.Locker
func main() {
	var lock myLocker = new(sync.Mutex)
	lock.Lock() //ok
	lock.Unlock() //ok
}



