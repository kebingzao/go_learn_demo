package main

/*
使用“nil” Slices and Maps
在一个nil的slice中添加元素是没问题的，但对一个map做同样的事将会生成一个运行时的panic。
*/

//失败的例子：
/*func main() {
	var m map[string]int
	m["one"] = 1 //error
}*/

/*
print::
panic: assignment to entry in nil map

goroutine 1 [running]:
panic(0x45a540, 0xc04203a000)
	C:/Go/src/runtime/panic.go:500 +0x1af
main.main()
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/8-Use-nil-Slices-and-Maps.go:11 +0x6f
*/
//正确的是:

func main() {
    var s []int
    s = append(s,1)
}

