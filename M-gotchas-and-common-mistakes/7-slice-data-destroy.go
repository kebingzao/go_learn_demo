package main

/*
Slice的数据“毁坏”
比如说你需要重新一个路径（在slice中保存）。
你通过修改第一个文件夹的名字，然后把名字合并来创建新的路径，来重新划分指向各个文件夹的路径。
*/

// 例子：

/*
import (
	"fmt"
	"bytes"
)
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1,"suffix"...)
	path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB (not ok)
	fmt.Println("new path =>",string(path))
}
*/

/*
结果与你想的不一样。与"AAAAsuffix/BBBBBBBBB"相反，你将会得到"AAAAsuffix/uffixBBBB"。
这个情况的发生是因为两个文件夹的slice都潜在的引用了同一个原始的路径slice。
这意味着原始路径也被修改了。根据你的应用，这也许会是个问题。

通过分配新的slice并拷贝需要的数据，你可以修复这个问题。另一个选择是使用完整的slice表达式。
*/

import (
	"fmt"
	"bytes"
)
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex:sepIndex] //full slice expression
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1,"suffix"...)
	path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)
	fmt.Println("new path =>",string(path))
}

//完整的slice表达式中的额外参数可以控制新的slice的容量。
// 现在在那个slice后添加元素将会触发一个新的buffer分配，而不是覆盖第二个slice中的数据。




