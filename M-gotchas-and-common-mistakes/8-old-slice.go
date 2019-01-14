package main

/*
陈旧的(Stale)Slices
多个slice可以引用同一个数据。
比如，当你从一个已有的slice创建一个新的slice时，这就会发生。如果你的应用功能需要这种行为，那么你将需要关注下“走味的”slice。

在某些情况下，在一个slice中添加新的数据，在原有数组无法保持更多新的数据时，将导致分配一个新的数组。
而现在其他的slice还指向老的数组（和老的数据）。
*/

// 例子：
import "fmt"
func main() {
	s1 := []int{1,2,3}
	fmt.Println(len(s1),cap(s1),s1) //prints 3 3 [1 2 3]
	s2 := s1[1:]
	fmt.Println(len(s2),cap(s2),s2) //prints 2 2 [2 3]
	for i := range s2 { s2[i] += 20 }
	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]  这时候其实引用的还是 s1 的数组
	s2 = append(s2,4) // 因为原数组s1没法加入了，所以这时候 s2 其实就是全新的数组了，跟s1没有关系了
	for i := range s2 { s2[i] += 10 }
	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}