package main

import "fmt"


/*
"switch"声明中的失效行为
在“switch”声明语句中的“case”语句块在默认情况下会break。这和其他语言中的进入下一个“next”代码块的默认行为不同。
*/

//失败的例子：
/*func main() {
	isSpace := func(ch byte) bool {
		switch(ch) {
		case ' ': //error
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //prints true (ok)
	fmt.Println(isSpace(' '))  //prints false (not ok)
}*/


// 正确的做法：
// 你可以通过在每个“case”块的结尾使用“fallthrough”，来强制“case”代码块进入。
// 你也可以重写switch语句，来使用“case”块中的表达式列表。

func main() {
	isSpace := func(ch byte) bool {
		switch(ch) {
		case ' ', '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //prints true (ok)
	fmt.Println(isSpace(' '))  //prints true (ok)
}