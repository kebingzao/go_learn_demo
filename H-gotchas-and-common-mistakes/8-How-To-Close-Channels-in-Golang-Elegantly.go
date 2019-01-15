package main

/*
如何优雅地关闭Go channel -- https://www.jianshu.com/p/d24dfbb33781

关于 channel 关闭的几个问题：
1 在不能更改channel状态的情况下，没有简单普遍的方式来检查channel是否已经关闭了
2 关闭已经关闭的channel会导致panic，所以在closer(关闭者)不知道channel是否已经关闭的情况下去关闭channel是很危险的
3 发送值到已经关闭的channel会导致panic，所以如果sender(发送者)在不知道channel是否已经关闭的情况下去向channel发送值是很危险的
*/

// 需要一个简单的方法来检查 channel 是否已经关闭：

/*import "fmt"

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}*/

/*
上面已经提到了，没有一种适用的方式来检查channel是否已经关闭了。
但是，就算有一个简单的 closed(chan T) bool函数来检查channel是否已经关闭，它的用处还是很有限的，就像内置的len函数用来检查缓冲channel中元素数量一样。
原因就在于，已经检查过的channel的状态有可能在调用了类似的方法返回之后就修改了，因此返回来的值已经不能够反映刚才检查的channel的当前状态了。
尽管在调用closed(ch)返回true的情况下停止向channel发送值是可以的，
但是如果调用closed(ch)返回false，那么关闭channel或者继续向channel发送值就不安全了（会panic）。
*/
//=============================================================================================================
// The Channel Closing Principle -- 不要从接收端关闭channel，也不要关闭有多个并发发送者的channel
// 简单的来说，如果只有一个唯一的 sender，那么如果要关闭 channel，只能这个 sender 自己去关闭
// 如果sender(发送者)只是唯一的sender或者是channel最后一个活跃的sender，那么你应该在sender的goroutine关闭channel，
// 从而通知receiver(s)(接收者们)已经没有值可以读了。维持这条原则将保证永远不会发生向一个已经关闭的channel发送值或者关闭一个已经关闭的channel。

// 如果你因为某种原因从接收端（receiver side）关闭channel或者在多个发送者中的一个关闭channel，
// 那么你应该使用列在Golang panic/recover Use Cases的函数来安全地发送值到channel中（假设channel的元素类型是T）

/*
func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call
			closed = true
		}
	}()

	ch <- value // panic if ch is closed
	return false // <=> closed = false; return
}
*/

/*
如果channel ch没有被关闭的话，那么这个函数的性能将和ch <- value接近。对于channel关闭的时候，
SafeSend函数只会在每个sender goroutine中调用一次，因此程序不会有太大的性能损失。
同样的想法也可以用在从多个goroutine关闭channel中：
*/

/*
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}
*/

// 很多人喜欢用sync.Once来关闭channel： 这样只会关闭一次channel，不会重复关闭，从而导致 panic

/*
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func(){
		close(mc.C)
	})
}
*/

// 当然了，我们也可以用sync.Mutex来避免多次关闭channel：

/*
type MyChannel struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannel) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
*/

// 我们应该要理解为什么Go不支持内置SafeSend和SafeClose函数，
// 原因就在于并不推荐从接收端或者多个并发发送端关闭channel。
// Golang甚至禁止关闭只接收（receive-only）的channel。

//======================================一个优化的方案===================

/*
上面的SaveSend函数有一个缺点是，在select语句的case关键字后不能作为发送操作被调用（译者注：类似于 case SafeSend(ch, t):）。
另外一个缺点是，很多人，包括我自己都觉得上面通过使用panic/recover和sync包的方案不够优雅。
针对各种场景，下面介绍不用使用panic/recover和sync包，纯粹是利用channel的解决方案。
*/

// 第一种场景： 多个接收者，一个发送者，sender通过关闭data channel说“不再发送”
// 这是最简单的场景了，就只是当sender不想再发送的时候让sender关闭data 来关闭channel：

/*
import (
	"time"
	"math/rand"
	"sync"
	"log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
*/

// 第二种情况，一个接收者，多个发送者，解决方法就是 receiver 通过关闭一个额外的signal channel说“请停止发送”。
// 所以是接收者来控制关闭信号，只是发送信号而已，并不是去关闭channel， 主要就是通知 sender 去关闭

/*
import (
	"time"
	"math/rand"
	"sync"
	"log"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)

				select {
				case val := <- stopCh:
					fmt.Println("===============:", val)
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				// the receiver of the dataCh channel is
				// also the sender of the stopCh channel.
				// It is safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
*/


// 第三种情况： 多个接收者， 多个发送者。
/*
它们当中任意一个通过通知一个moderator（仲裁者）关闭额外的signal channel来说“让我们结束游戏吧”
这是最复杂的场景了。我们不能让任意的receivers和senders关闭data channel，
也不能让任何一个receivers通过关闭一个额外的signal channel来通知所有的senders和receivers退出游戏。
这么做的话会打破channel closing principle。但是，我们可以引入一个moderator来关闭一个额外的signal channel。
这个例子的一个技巧是怎么通知moderator去关闭额外的signal channel：
*/

/*
import (
	"time"
	"math/rand"
	"sync"
	"log"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its receivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its receivers is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <- toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					// 标记是由那个 sender 关闭的
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				// 标记是由哪个sender 捕获的
				select {
				case <- stopCh:
					return
				default:
				}

				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case <- stopCh:
					return
				default:
				}

				select {
				case <- stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber-1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						// 标记是由那个 receiver 断开的
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
*/

// 结果就是，有时候是由 sender 断开的，有时候是由 receiver 断开的。

//============================================结论==========
/*这里没有一种场景要求你去打破channel closing principle。如果你遇到了这种场景，请思考一下你的设计并重写你的代码。
用Go编程就像在创作艺术。*/












