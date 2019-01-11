package main

/*
Slices和Arrays是一维的
看起来Go好像支持多维的Array和Slice，但不是这样的。尽管可以创建数组的数组或者切片的切片。
对于依赖于动态多维数组的数值计算应用而言，Go在性能和复杂度上还相距甚远。
你可以使用纯一维数组、“独立”切片的切片，“共享数据”切片的切片来构建动态的多维数组。
如果你使用纯一维的数组，你需要处理索引、边界检查、当数组需要变大时的内存重新分配。
使用“独立”slice来创建一个动态的多维数组需要两步。首先，你需要创建一个外部的slice。
然后，你需要分配每个内部的slice。内部的slice相互之间独立。你可以增加减少它们，而不会影响其他内部的slice。
*/

//例子：
/*func main() {
	x := 2
	y := 4
	table := make([][]int,x)
	for i:= range table {
		table[i] = make([]int,y)
	}
}*/

// 使用“共享数据”slice的slice来创建一个动态的多维数组需要三步。首先，你需要创建一个用于存放原始数据的数据“容器”。
// 然后，你再创建外部的slice。最后，通过重新切片原始数据slice来初始化各个内部的slice。

import "fmt"
func main() {
	h, w := 2, 4
	raw := make([]int,h*w)
	for i := range raw {
		raw[i] = i
	}
	fmt.Println(raw,&raw[4])
	//prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>
	table := make([][]int,h)
	for i:= range table {
		table[i] = raw[i*w:i*w + w]
	}
	fmt.Println(table,&table[1][0])
	//prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>
}

// 关于多维array和slice已经有了专门申请，但现在看起来这是个低优先级的特性。