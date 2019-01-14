package main

/*
使用指针接收方法的值的实例
只要值是可取址的，那在这个值上调用指针接收方法是没问题的。换句话说，在某些情况下，你不需要在有一个接收值的方法版本。

然而并不是所有的变量是可取址的。Map的元素就不是。通过interface引用的变量也不是。
*/

//失败的例子：

import "fmt"
type data struct {
	name string
}
func (p *data) print() {
	fmt.Println("name:",p.name)
}
type printer interface {
	print()
}
func main() {
	d1 := data{"one"}
	d1.print() //ok
	var in printer = data{"two"} //error
	in.print()
	m := map[string]data {"x":data{"three"}}
	m["x"].print() //error
}

/*
H-gotchas-and-common-mistakes\1-uses-a-pointer-to-receive-a-method.go:25: cannot use data literal (type data) as type printer in assignment:
	data does not implement printer (print method has pointer receiver)
H-gotchas-and-common-mistakes\1-uses-a-pointer-to-receive-a-method.go:28: cannot call pointer method on m["x"]
H-gotchas-and-common-mistakes\1-uses-a-pointer-to-receive-a-method.go:28: cannot take the address of m["x"]
*/




