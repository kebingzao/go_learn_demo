package main

/*
String在“range”语句中的迭代值
Logging库一般提供不同的log等级。与这些logging库不同，Go中log包在你调用它的Fatal*()和Panic*()函数时，可以做的不仅仅是log。
当你的应用调用这些函数时，Go也将会终止应用 :-)
*/

import "fmt"
func main() {
	data := "A\xfe\x02\xff\x04"
	for _,v := range data {
		fmt.Printf("%#x ",v)
	}
	//prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)
	fmt.Println()
	for _,v := range []byte(data) {
		fmt.Printf("%#x ",v)
	}
	//prints: 0x41 0xfe 0x2 0xff 0x4 (good)
}