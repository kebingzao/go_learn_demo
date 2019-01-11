package main

import "fmt"

/*
访问不存在的Map Keys
这对于那些希望得到“nil”标示符的开发者而言是个技巧（和其他语言中做的一样）。
如果对应的数据类型的“零值”是“nil”，那返回的值将会是“nil”，但对于其他的数据类型是不一样的。
检测对应的“零值”可以用于确定map中的记录是否存在，但这并不总是可信（比如，如果在二值的map中“零值”是false，这时你要怎么做）。
检测给定map中的记录是否存在的最可信的方法是，通过map的访问操作，检查第二个返回的值。
*/

//失败的例子：
/*func main() {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if v := x["two"]; v == "" { //incorrect
		fmt.Println("no entry")
	}
}*/

// 正确的做法：
func main() {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if _,ok := x["two"]; !ok {
		fmt.Println("no entry")
	}
}