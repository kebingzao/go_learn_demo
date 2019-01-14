package main

/*
更新Map的值
如果你有一个struct值的map，你无法更新单个的struct值。
*/

//失败的例子：
/*
type data struct {
	name string
}
func main() {
	m := map[string]data {"x":{"one"}}
	m["x"].name = "two" //error
}
*/

//H-gotchas-and-common-mistakes\2-update-map-value.go:14: cannot assign to struct field m["x"].name in map

// 这个操作无效是因为map元素是无法取址的。
// 而让Go新手更加困惑的是slice元素是可以取址的。

/*
import "fmt"
type dataStruct struct {
	name string
}
func main() {
	s := []dataStruct{{"one"}}
	s[0].name = "two" //ok
	fmt.Println(s)    //prints: [{two}]
}
*/

// 针对上面的问题， 第一个有效的方法是使用一个临时变量。

/*
import "fmt"
type data2 struct {
	name string
}
func main() {
	m := map[string]data2 {"x":{"one"}}
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v",m) //prints: map[x:{two}]
}
*/

// 另一个有效的方法是使用指针的map。
/*
import "fmt"
type data2 struct {
	name string
}
func main() {
	m := map[string]*data2 {"x":{"one"}}
	m["x"].name = "two" //ok
	fmt.Println(m["x"]) //prints: &{two}
}

*/

// 如果是下面这种情况，就会出现空指针问题

type data2 struct {
	name string
}
func main() {
	m := map[string]*data2 {"x":{"one"}}
	m["z"].name = "what?" //???
}

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x0 pc=0x40114a]

goroutine 1 [running]:
panic(0x45aac0, 0xc042004090)
	C:/Go/src/runtime/panic.go:500 +0x1af
main.main()
	F:/airdroid_code/go/src/go_learn_demo/H-gotchas-and-common-mistakes/2-update-map-value.go:71 +0x10a

*/
