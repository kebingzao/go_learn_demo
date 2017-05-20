package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

// GO 信号实例
// http://www.yiibai.com/go/golang-signals.html

func main() {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	// 等待程序停止的时候，发送信号
	// step2
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		// 接收信号
		// step3
		sig := <-sigs
		// 打印
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	// 接下来等待程序被退出（可以手动退出）
	// step 1
	fmt.Println("awaiting signal")
	// 等待状态修改
	// step4
	<-done
	// step5
	fmt.Println("exiting")
}