package main

/*
偶然的变量隐藏Accidental Variable Shadowing
短式变量声明的语法如此的方便（尤其对于那些使用过动态语言的开发者而言），很容易让人把它当成一个正常的分配操作。
如果你在一个新的代码块中犯了这个错误，将不会出现编译错误，但你的应用将不会做你所期望的事情。
*/

//失败的例子：
import "fmt"
func main() {
	x := 1
	fmt.Println(x)     //prints 1
	{
		fmt.Println(x) //prints 1
		x := 2
		fmt.Println(x) //prints 2
	}
	fmt.Println(x)     //prints 1 (bad if you need 2)
}

/*
print::
1
1
2
1
*/

/*即使对于经验丰富的Go开发者而言，这也是一个非常常见的陷阱。这个坑很容易挖，但又很难发现。

你可以使用 vet命令来发现一些这样的问题。 默认情况下， vet不会执行这样的检查，你需要设置-shadow参数：
go tool vet -shadow your_file.go。*/


