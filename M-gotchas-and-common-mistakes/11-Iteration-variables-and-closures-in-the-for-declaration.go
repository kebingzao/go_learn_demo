package main

/*
"for"声明中的迭代变量和闭包
这在Go中是个很常见的技巧。for语句中的迭代变量在每次迭代时被重新使用。
这就意味着你在for循环中创建的闭包（即函数字面量）将会引用同一个变量（而在那些goroutine开始执行时就会得到那个变量的值）。
*/

// 错误的例子：
/*
import (
	"fmt"
	"time"
)
func main() {
	data := []string{"one","two","three"}
	for _,v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}
*/

// print::
/*
three
three
three
*/

// 最简单的解决方法（不需要修改goroutine）是，在for循环代码块内把当前迭代的变量值保存到一个局部变量中。

/*
import (
	"fmt"
	"time"
)
func main() {
	data := []string{"one","two","three"}
	for _,v := range data {
		vcopy := v //
		go func() {
			fmt.Println(vcopy)
		}()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}
*/
// 另一个解决方法是把当前的迭代变量作为匿名goroutine的参数。
/*
import (
	"fmt"
	"time"
)
func main() {
	data := []string{"one","two","three"}
	for _,v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}
*/

// 下面还有比较复杂的错误版本：

/*
import (
	"fmt"
	"time"
)
type field struct {
	name string
}
func (p *field) print() {
	fmt.Println(p.name)
}
func main() {
	data := []field{ {"one"},{"two"},{"three"} }
	for _,v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}
*/

// 会输出 ：
/*
three
three
three
*/

// 那么改一下，改成这个：

/*
import (
	"fmt"
	"time"
)
type field struct {
	name string
}
func (p *field) print() {
	fmt.Println(p.name)
}
func main() {
	data := []field{ {"one"},{"two"},{"three"} }
	for _,v := range data {
		v := v
		go v.print()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}
*/

// 还可以换成这个：
/*
import (
	"fmt"
	"time"
)
type field struct {
	name string
}
func (p *field) print() {
	fmt.Println(p.name)
}
func main() {
	data := []*field{ {"one"},{"two"},{"three"} }
	for _,v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}
*/

// 上面那个方式也行，也就是都加指针。 另一种方式，都不加指针也行：

import (
	"fmt"
	"time"
)
type field struct {
	name string
}
func (p field) print() {
	fmt.Println(p.name)
}
func main() {
	data := []field{ {"one"},{"two"},{"three"} }
	for _,v := range data {
		go v.print()
	}
	time.Sleep(3 * time.Second)
	//goroutines print: one, two, three
}














