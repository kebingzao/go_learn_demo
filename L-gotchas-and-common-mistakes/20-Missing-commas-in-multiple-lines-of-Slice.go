package main

/*
在多行的Slice、Array和Map语句中遗漏逗号
*/

//失败的例子：
/*func main() {
	x := []int{
		1,
		2 //error
	}
	_ = x
}*/

/*
print::
gotchas-and-common-mistakes\20-Missing-commas-in-multiple-lines-of-Slice.go:11:
syntax error: unexpected semicolon or newline, expecting comma or }
*/

// 正确的做法：
func main() {
	x := []int{
		1,
		2,
	}
	x = x
	y := []int{3,4,} //no error
	y = y
}

// 当你把声明折叠到单行时，如果你没加末尾的逗号，你将不会得到编译错误。