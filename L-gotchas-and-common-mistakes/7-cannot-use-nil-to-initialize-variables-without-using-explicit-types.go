package main

/*
不使用显式类型，无法使用“nil”来初始化变量
nil标志符用于表示interface、函数、maps、slices和channels的“零值”。
如果你不指定变量的类型，编译器将无法编译你的代码，因为它猜不出具体的类型。
*/

//失败的例子：
/*func main() {
	var x = nil //error
	_ = x
}*/

/*
print::
gotchas-and-common-mistakes\7-cannot-use-nil-to-initialize-variables-without-using-explicit-types.go:11:
use of untyped nil*/

//正确的是:
func main() {
	var x interface{} = nil
	_ = x
}

