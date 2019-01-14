package main

/*
向已关闭的Channel发送会引起Panic
从一个关闭的channel接收是安全的。在接收状态下的ok的返回值将被设置为false，这意味着没有数据被接收。
如果你从一个有缓存的channel接收，你将会首先得到缓存的数据，一旦它为空，返回的ok值将变为false。

向关闭的channel中发送数据会引起panic。
这个行为有文档说明，但对于新的Go开发者的直觉不同，他们可能希望发送行为与接收行为很像。
*/
/*
import (
	"fmt"
	"time"
)
func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get the first result
	fmt.Println(<-ch)
	close(ch) //not ok (you still have other senders)，因为这时候关闭的情况下，gorouter还在往channel发送数据，会导致panic
	//do other work
	time.Sleep(2 * time.Second)
}*/
// print::
/*
2
panic: send on closed channel

goroutine 7 [running]:
panic(0x49ad40, 0xc04206a000)
C:/Go/src/runtime/panic.go:500 +0x1af
main.main.func1(0xc042038180, 0x2)
F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/32-Sending-to-a-closed-Channel-will-cause-Panic.go:20 +0x58
created by main.main
F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/32-Sending-to-a-closed-Channel-will-cause-Panic.go:21 +0x7a
panic: send on closed channel

goroutine 6 [running]:
panic(0x49ad40, 0xc0420042b0)
C:/Go/src/runtime/panic.go:500 +0x1af
main.main.func1(0xc042038180, 0x1)
F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/32-Sending-to-a-closed-Channel-will-cause-Panic.go:20 +0x58
created by main.main
F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/32-Sending-to-a-closed-Channel-will-cause-Panic.go:21 +0x7a
*/

/*
根据不同的应用，修复方法也将不同。可能是很小的代码修改，也可能需要修改应用的设计。
无论是哪种方法，你都需要确保你的应用不会向关闭的channel中发送数据。

上面那个有bug的例子可以通过使用一个特殊的废弃的channel来向剩余的worker发送不再需要它们的结果的信号来修复。
*/

import (
	"fmt"
	"time"
)
func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2: fmt.Println(idx,"sent result")
			case <- done: fmt.Println(idx,"exiting")
			}
		}(i)
	}
	//get first result
	fmt.Println("result:",<-ch)
	close(done)
	//do other work
	time.Sleep(3 * time.Second)
}

// 这个输出的结果也是不确定的，就是看这三个gorouter的执行顺序是如何






