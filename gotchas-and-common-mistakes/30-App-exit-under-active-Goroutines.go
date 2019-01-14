package main

/*
有活动的Goroutines下的应用退出
应用将不会等待所有的goroutines完成。这对于初学者而言是个很常见的错误。每个人都是以某个程度开始，因此如果犯了初学者的错误也没神马好丢脸的 :-)
*/


//import (
//	"fmt"
//	"time"
//)
/*func main() {
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		go doit(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}
func doit(workerId int) {
	fmt.Printf("[%v] is running\n",workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n",workerId)
}*/

/*
print::
[1] is running
[0] is running
all done!
*/

/*
一个最常见的解决方法是使用“WaitGroup”变量。
它将会让主goroutine等待所有的worker goroutine完成。
如果你的应用有长时运行的消息处理循环的worker，你也将需要一个方法向这些goroutine发送信号，让它们退出。
你可以给各个worker发送一个“kill”消息。另一个选项是关闭一个所有worker都接收的channel。这是一次向所有goroutine发送信号的简单方式。
*/

/*import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i,done,wg)
	}
	close(done)
	wg.Wait()
	fmt.Println("all done!")
}
func doit(workerId int,done <-chan struct{},wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n",workerId)
	defer wg.Done()
	<- done
	fmt.Printf("[%v] is done\n",workerId)
}*/

// print::
/*
[1] is running
[1] is done
[0] is running
[0] is done
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc04200429c)
C:/Go/src/runtime/sema.go:47 +0x37
sync.(*WaitGroup).Wait(0xc042004290)
C:/Go/src/sync/waitgroup.go:131 +0x9e
main.main()
F:/airdroid_code/go/src/go_learn_demo/gotchas-and-common-mistakes/30-App-exit-under-active-Goroutines.go:55 +0x101
*/

/*这边报了这个错，显示死锁？？ all goroutines are asleep - deadlock
这可不太好 :-) 发送了神马？为什么会出现死锁？worker退出了，它们也执行了wg.Done()。应用应该没问题啊。

死锁发生是因为各个worker都得到了原始的“WaitGroup”变量的一个拷贝。
当worker执行wg.Done()时，并没有在主goroutine上的“WaitGroup”变量上生效。*/

// 所以改成这样： wg 要传入指令，而不应该是值传递

import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i,wq,done,&wg)
	}
	for i := 0; i < workerCount; i++ {
		wq <- i
	}
	close(done)
	wg.Wait()
	fmt.Println("all done!")
}
func doit(workerId int, wq <-chan interface{},done <-chan struct{},wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n",workerId)
	defer func() {
		fmt.Printf("[%v] is defer\n",workerId)
		wg.Done()
	}()
	for {
		select {
		case m := <- wq:
			fmt.Printf("[%v] m => %v\n",workerId,m)
		case <- done:
			fmt.Printf("[%v] is done\n",workerId)
			return
		}
	}
}

// print::
/*[1] is running
[1] m => 0
[1] m => 1
[0] is running
[0] is done
[1] is done
all done!*/
