package main

/*
字符串的长度
让我们假设你是Python开发者，你有下面这段代码：

data = u'♥'
print(len(data)) #prints: 1

当把它转换为Go代码时，你可能会大吃一惊。
*/

//例子：
/*
import "fmt"
func main() {
	data := "♥"
	fmt.Println(len(data)) //prints: 3
}
*/
// 这个是因为 内建的len()函数返回byte的数量，而不是像Python中计算好的unicode字符串中字符的数量。

// 要在Go中得到相同的结果，可以使用“unicode/utf8”包中的RuneCountInString()函数。
// 正确的做法：
/*
import (
	"fmt"
	"unicode/utf8"
)
func main() {
	data := "♥"
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1
}*/

//但是 理论上说RuneCountInString()函数并不返回字符的数量，因为单个字符可能占用多个rune。，比如一下这种情况
import (
	"fmt"
	"unicode/utf8"
)
func main() {
	data := "é"
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2
}