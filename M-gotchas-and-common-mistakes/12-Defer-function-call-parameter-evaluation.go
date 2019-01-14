package main

/*
Defer函数调用参数的求值
被defer的函数的参数会在defer声明时求值（而不是在函数实际执行时）。
Arguments for a deferred function call are evaluated when the defer statement is evaluated (not when the function is actually executing).
*/

import "fmt"
func main() {
	var i int = 1
	defer fmt.Println("result =>",func() int { return i * 2 }())  // 执行到这边的时候，其实就执行了，所以这时候 i 就是 1，而不是后面的 2
	i++
	//prints: result => 2 (not ok if you expected 4)
}


