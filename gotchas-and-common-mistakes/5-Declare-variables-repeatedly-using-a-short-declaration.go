package main

/*
使用简式声明重复声明变量
你不能在一个单独的声明中重复声明一个变量，但在多变量声明中这是允许的，其中至少要有一个新的声明变量。
重复变量需要在相同的代码块内，否则你将得到一个隐藏变量。
*/

//失败的例子：
/*func main() {
    one := 0
    one := 1 //error
}*/

/*
print::
gotchas-and-common-mistakes\5-Declare-variables-repeatedly-using-a-short-declaration.go:12: no new variables on left side of :=
*/

//正确的是:
func main() {
	one := 0
	one, two := 1,2
	one,two = two,one
}

