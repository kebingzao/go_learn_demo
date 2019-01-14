package main

/*
失败的类型断言
失败的类型断言返回断言声明中使用的目标类型的“零值”。这在与隐藏变量混合时，会发生未知情况。
*/
// 错误的方式：
/*
import "fmt"
func main() {
	var data interface{} = "great"
	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>",data)
	} else {
		fmt.Println("[not an int] value =>",data)
		//prints: [not an int] value => 0 (not "great")
	}
}
*/

// 正确的方法： 要重新赋值出来
import "fmt"
func main() {
	var data interface{} = "great"
	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>",res)
	} else {
		fmt.Println("[not an int] value =>",data)
		//prints: [not an int] value => great (as expected)
	}
}




