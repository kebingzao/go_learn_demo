package main

/*
字符串不总是UTF8文本
字符串的值不需要是UTF8的文本。它们可以包含任意的字节。只有在string literal使用时，字符串才会是UTF8。
即使之后它们可以使用转义序列来包含其他的数据。
为了知道字符串是否是UTF8，你可以使用“unicode/utf8”包中的ValidString()函数。
*/

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true
	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
}