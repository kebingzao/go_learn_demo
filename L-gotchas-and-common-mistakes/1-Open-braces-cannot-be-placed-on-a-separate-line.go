package main

import "fmt"
/*
开大括号不能放在单独的一行
在大多数其他使用大括号的语言中，你需要选择放置它们的位置。Go的方式不同。你可以为此感谢下自动分号的注入（没有预读）。是的，Go中也是有分号的：-）
*/

//失败的例子：
/*
func main()
{ //error, can't have the opening brace on a separate line
	fmt.Println("hello there!")
}
*/
//print:: syntax error: unexpected semicolon or newline before {


//正确的是:
func main() {
	fmt.Println("works!")
}