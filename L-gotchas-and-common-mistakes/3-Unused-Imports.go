package main

/*
未使用的Imports
如果你引入一个包，而没有使用其中的任何函数、接口、结构体或者变量的话，代码将会编译失败。
你可以使用goimports来增加引入或者移除未使用的引用：
$ go get golang.org/x/tools/cmd/goimports
如果你真的需要引入的包，你可以添加一个下划线标记符，_，来作为这个包的名字，从而避免编译失败。下滑线标记符用于引入，但不使用。
*/

//失败的例子：
/*
import (
	"fmt"
	"log"
	"time"
)

func main() {

}
*/

/*
print::
gotchas-and-common-mistakes\3-Unused-Imports.go:4: imported and not used: "fmt"
gotchas-and-common-mistakes\3-Unused-Imports.go:5: imported and not used: "log"
gotchas-and-common-mistakes\3-Unused-Imports.go:6: imported and not used: "time"
*/

//正确的是:
import (
	_ "fmt"
	"log"
	"time"
)
var _ = log.Println
func main() {
	_ = time.Now
}
// 另一个选择是移除或者注释掉未使用的imports ：-）