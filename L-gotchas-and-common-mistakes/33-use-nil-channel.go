package main

/*
使用"nil" Channels
在一个nil的channel上发送和接收操作会被永久阻塞。这个行为有详细的文档解释，但它对于新的Go开发者而言是个惊喜。
*/


/*import (
	"fmt"
	"time"
)
func main() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	//get first result
	fmt.Println("result:",<-ch)
	//do other work
	time.Sleep(2 * time.Second)
}*/


// print::
/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive (nil chan)]:
main.main()
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:19 +0xd4

goroutine 5 [chan send (nil chan)]:
main.main.func1(0xc042024020, 0x0)
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:15 +0x5b
created by main.main
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:16 +0x7b

goroutine 6 [chan send (nil chan)]:
main.main.func1(0xc042024020, 0x1)
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:15 +0x5b
created by main.main
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:16 +0x7b

goroutine 7 [chan send (nil chan)]:
main.main.func1(0xc042024020, 0x2)
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:15 +0x5b
created by main.main
	F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/33-use-nil-channel.go:16 +0x7b
*/

/*
如果运行代码你将会看到一个runtime错误：死锁？？？
这个行为可以在select声明中用于动态开启和关闭case代码块的方法。
*/

import "fmt"
import "time"
func main() {
	inch := make(chan int)
	outch := make(chan int)
	go func() {
		var in <- chan int = inch
		var out chan <- int
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <- in:
				out = outch
				in = nil
			}
		}
	}()
	go func() {
		for r := range outch {
			fmt.Println("result:",r)
		}
	}()
	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}


