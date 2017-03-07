package main

import (
	"fmt"
)

// 此示例里面 select 会一直等待等到某个 case 语句完成， 也就是等到成功从 ch1 或者 ch2 中读到数据,
// 如果都不满足条件且存在default case, 那么default case会被执行。 则 select 语句结束。

func main(){
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	select {
	case e1 := <-ch1:
	//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v",e1)
	case e2 := <-ch2:
	//如果ch2通道成功读取数据，则执行该case处理语句
		fmt.Printf("2th case is selected. e2=%v",e2)
	default:
	//如果上面case都没有成功，则进入default处理流程
		fmt.Println("default!.")
	}
}