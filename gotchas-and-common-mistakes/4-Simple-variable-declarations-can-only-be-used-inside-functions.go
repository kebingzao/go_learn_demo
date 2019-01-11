package main

/*
简式的变量声明仅可以在函数内部使用
*/

//失败的例子：
/*
myvar := 1 //error
func main() {
}
*/

/*
print::
gotchas-and-common-mistakes\4-Simple-variable-declarations-can-only-be-used-inside-functions.go:8: syntax error: non-declaration statement outside function body
*/

//正确的是:
var myvar = 1
func main() {
}

