package main

/*
Map的容量
你可以在map创建时指定它的容量，但你无法在map上使用cap()函数。
*/

//失败的例子：
/*func main() {
    m := make(map[string]int,99)
    cap(m) //error
}*/

/*
print::
gotchas-and-common-mistakes\9-Map-capacity.go:11: invalid argument m (type map[string]int) for cap
*/

// 正确的做法：
import "fmt"
func main() {
    m := make(map[string]int,99)
    fmt.Println(len(m)) // 输出0，这个是因为里面确实没有数据，但是没法得到 容量99
}