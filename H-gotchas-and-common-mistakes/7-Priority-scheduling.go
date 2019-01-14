package main

/*
优先调度

有可能会出现这种情况，一个无耻的goroutine阻止其他goroutine运行。
当你有一个不让调度器运行的for循环时，这就会发生。
*/

/*import "fmt"
func main() {
	done := false
	go func(){
		done = true
	}()
	for !done {
	}
	fmt.Println("done!")
}*/

// 试了一下，好像不会出现无限循环，但是换成这种方式，就会知道

/*
import (
	"fmt"
	"runtime"
)
func main() {
	done := false
	go func(){
		done = true
	}()
	for !done {
		fmt.Println("not done!") //not inlined
	}
	fmt.Println("done!")
}
*/

/*
not done!
not done!
。
。
。
not done!
not done!
not done!
done!

*/

// 还有一种方式就是显示的唤醒调度器。你可以使用“runtime”包中的Goshed()函数。

import (
"fmt"
"runtime"
)
func main() {
	done := false
	go func(){
		done = true
	}()
	for !done {
		fmt.Println("not done!") //not inlined
		runtime.Gosched()
	}
	fmt.Println("done!")
}

// 这种情况下，只会执行一次 for 里面的东西，然后就会重新调度了

/*
not done!
done!
*/