package main

/*
在Slice中"隐藏"数据
当你重新划分一个slice时，新的slice将引用原有slice的数组。
如果你忘了这个行为的话，在你的应用分配大量临时的slice用于创建新的slice来引用原有数据的一小部分时，会导致难以预期的内存使用。
*/

// 例子：
/*
import "fmt"
func get() []byte {
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
	return raw[:3]
}
func main() {
	data := get()
	fmt.Println(len(data),cap(data),&data[0]) //prints: 3 10000 <byte_addr_x> , 容量还是那么大有 10000，太浪费了
}
*/

// 为了避免这个陷阱，你需要从临时的slice中拷贝数据（而不是重新划分slice）。

import "fmt"
func get() []byte {
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
	res := make([]byte,3)
	copy(res,raw[:3])
	return res
}
func main() {
	data := get()
	fmt.Println(len(data),cap(data),&data[0]) //prints: 3 3 <byte_addr_y>
}
