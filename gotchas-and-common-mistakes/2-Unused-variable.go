package main

import "fmt"

/*
未使用的变量
如果你有未使用的变量，代码将编译失败。当然也有例外。在函数内一定要使用声明的变量，但未使用的全局变量是没问题的。
如果你给未使用的变量分配了一个新的值，代码还是会编译失败。你需要在某个地方使用这个变量，才能让编译器愉快的编译。
*/

//失败的例子：
/*
var gvar int //not an error
func main() {
    var one int   //error, unused variable
    two := 2      //error, unused variable
    var three int //error, even though it's assigned 3 on the next line
    three = 3
}
*/
/*
print::
gotchas-and-common-mistakes\Unused-variable.go:11: one declared and not used
gotchas-and-common-mistakes\Unused-variable.go:12: two declared and not used
gotchas-and-common-mistakes\Unused-variable.go:13: three declared and not used
*/

//正确的是:
func main() {
	var one int
	_ = one
	two := 2
	fmt.Println(two)
	var three int
	three = 3
	one = three
	var four int
	four = four
}
//另一个选择是注释掉或者移除未使用的变量 ：-）