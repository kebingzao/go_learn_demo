package main

/*
"nil" Interfaces和"nil" Interfaces的值

这在Go中是第二最常见的技巧，因为interface虽然看起来像指针，但并不是指针。
interface变量仅在类型和值为“nil”时才为“nil”。

interface的类型和值会根据用于创建对应interface变量的类型和值的变化而变化。
当你检查一个interface变量是否等于“nil”时，这就会导致未预期的行为。

*/

//失败的例子：
/*
import "fmt"
func main() {
	var data *byte
	var in interface{}
	fmt.Println(data,data == nil) //prints: <nil> true
	fmt.Println(in,in == nil)     //prints: <nil> true
	in = data
	fmt.Println(in,in == nil)     //prints: <nil> false， 这时候 in 的值虽然是 nil，但是他的类型不是nil了，而是 *byte, 因此不等于nil
	//'data' is 'nil', but 'in' is not 'nil'
}
*/

// 另一个错误的例子就是：

/*
import "fmt"
func main() {
	doit := func(arg int) interface{} {
		var result *struct{} = nil
		if(arg > 0) {
			result = &struct{}{}
		}
		return result
	}
	if res := doit(-1); res != nil {
		fmt.Println("good result:",res) //prints: good result: <nil>
		//'res' is not 'nil', but its value is 'nil'
	}
}

*/








