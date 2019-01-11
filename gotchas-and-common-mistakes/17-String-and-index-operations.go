package main

import "fmt"

/*
String和索引操作
字符串上的索引操作返回一个byte值，而不是一个字符（和其他语言中的做法一样）。
*/

func main() {
	x := "text"
	fmt.Println(x[0]) //print 116
	fmt.Printf("%T",x[0]) //prints uint8
}

/*
如果你需要访问特定的字符串“字符”（unicode编码的points/runes），使用for range。
官方的“unicode/utf8”包和实验中的utf8string包（golang.org/x/exp/utf8string）也可以用。
utf8string包中包含了一个很方便的At()方法。把字符串转换为rune的切片也是一个选项。*/
