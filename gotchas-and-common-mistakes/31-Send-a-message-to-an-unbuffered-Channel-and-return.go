package main

/*
向无缓存的Channel发送消息，只要目标接收者准备好就会立即返回
发送者将不会被阻塞，除非消息正在被接收者处理。
根据你运行代码的机器的不同，接收者的goroutine可能会或者不会有足够的时间，在发送者继续执行前处理消息。
*/

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:",m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
}

// 这个很有意思，我执行了很多次，其中有些是：
//processed: cmd.1
// 有一些是：
//processed: cmd.1
//processed: cmd.2

//也就是 cmd.2 入channel之后，什么时候被输出，取决于主程序快还是慢
