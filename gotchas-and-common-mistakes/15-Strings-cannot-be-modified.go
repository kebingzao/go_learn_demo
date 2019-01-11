package main

import "fmt"

/*
Strings无法修改
尝试使用索引操作来更新字符串变量中的单个字符将会失败。string是只读的byte slice（和一些额外的属性）。
如果你确实需要更新一个字符串，那么使用byte slice，并在需要时把它转换为string类型。
*/

//失败的例子：
/*func main() {
	x := "text"
	x[0] = 'T'
	fmt.Println(x)
}*/

/*
print::
gotchas-and-common-mistakes\15-Strings-cannot-be-modified.go:14: cannot assign to x[0]
*/

// 正确的做法：
func main() {
	x := "text"
	xbytes := []byte(x)
	xbytes[0] = 'T'
	fmt.Println(string(xbytes)) //prints Text
}