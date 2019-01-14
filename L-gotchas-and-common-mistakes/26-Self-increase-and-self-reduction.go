package main

import "fmt"

/*
自增和自减
许多语言都有自增和自减操作。不像其他语言，Go不支持前置版本的操作。你也无法在表达式中使用这两个操作符。
*/

//失败的例子：
/*func main() {
	data := []int{1,2,3}
	i := 0
	++i //error
	fmt.Println(data[i++]) //error
}*/
/*
print::
gotchas-and-common-mistakes\26-Self-increase-and-self-reduction.go:14: syntax error: unexpected ++, expecting }
*/

// 正确的做法：
func main() {
	data := []int{1,2,3}
	i := 0
	i++
	fmt.Println(data[i])
}