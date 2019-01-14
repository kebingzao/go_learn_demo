package main

/*
从"for switch"和"for select"代码块中跳出
没有标签的“break”声明只能从内部的switch/select代码块中跳出来。
如果无法使用“return”声明的话，那就为外部循环定义一个标签是另一个好的选择。
*/

/*
import "fmt"
func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}
	fmt.Println("out!")
}
*/


// goto  也可以做到

import "fmt"
func main() {
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			goto loop
		}
	}
loop:
	fmt.Println("out!")
}

// 使用break label 和 goto label 都能跳出for循环；
// 不同之处在于：break标签只能用于for循环，且标签位于for循环前面，goto是指跳转到指定标签处