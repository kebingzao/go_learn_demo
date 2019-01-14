package main

/*
从Panic中恢复
recover()函数可以用于获取/拦截panic。仅当在一个defer函数中被完成时，调用recover()将会完成这个小技巧。
*/

// 失败的例子：
/*
import "fmt"
func main() {
	recover() //doesn't do anything
	panic("not good")
	recover() //won't be executed :)
	fmt.Println("ok")
}
*/

// print::
/*
panic: not good

goroutine 1 [running]:
panic(0x47f3a0, 0xc04203a1d0)
C:/Go/src/runtime/panic.go:500 +0x1af
main.main()
F:/airdroid_code/go/src/go_learn_demo/M-gotchas-and-common-mistakes/4-recover-from-panic.go:12 +0x82
*/

// 正确的方式：

/*
import "fmt"
func main() {
	defer func() {
		fmt.Println("recovered:",recover())
	}()
	panic("not good")
}
*/

//recover()的调用仅当它在defer函数中被直接调用时才有效。不能在其他的函数里面

import "fmt"
func doRecover() {
	fmt.Println("recovered =>",recover()) //prints: recovered => <nil>
}
func main() {
	defer func() {
		doRecover() //panic is not recovered
	}()
	panic("not good")
}





