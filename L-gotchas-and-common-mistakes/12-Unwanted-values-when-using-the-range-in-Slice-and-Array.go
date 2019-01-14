package main

import "fmt"

/*
在Slice和Array使用“range”语句时的出现的不希望得到的值
如果你在其他的语言中使用“for-in”或者“foreach”语句时会发生这种情况。
Go中的“range”语法不太一样。它会得到两个值：第一个值是元素的索引，而另一个值是元素的数据。
*/

//失败的例子：
/*func main() {
	x := []string{"a","b","c"}
	for v := range x {
		fmt.Println(v) //prints 0, 1, 2
	}
}*/

// 正确的做法：
func main() {
	x := []string{"a","b","c"}
	for _, v := range x {
		fmt.Println(v) //prints a, b, c
	}
}