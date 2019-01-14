package main

/*
在Slice, Array, and Map "range"语句中更新引用元素的值
在“range”语句中生成的数据的值是真实集合元素的拷贝。它们不是原有元素的引用。
这意味着更新这些值将不会修改原来的数据。同时也意味着使用这些值的地址将不会得到原有数据的指针。
*/

// 例子：
/*
import "fmt"
func main() {
	data := []int{1,2,3}
	for _,v := range data {
		v *= 10 //original item is not changed
	}
	fmt.Println("data:",data) //prints data: [1 2 3]
}
*/

//如果你需要更新原有集合中的数据，使用索引操作符来获得数据。
/*
import "fmt"
func main() {
	data := []int{1,2,3}
	for i,_ := range data {
		data[i] *= 10
	}
	fmt.Println("data:",data) //prints data: [10 20 30]
}
*/

/*
如果你的集合保存的是指针，那规则会稍有不同。
如果要更新原有记录指向的数据，你依然需要使用索引操作，但你可以使用for range语句中的第二个值来更新存储在目标位置的数据。*/

import "fmt"
func main() {
	data := []*struct{num int} { {1},{2},{3} }
	for _,v := range data {
		v.num *= 10
	}
	fmt.Println(data[0],data[1],data[2]) //prints &{10} &{20} &{30}
}













