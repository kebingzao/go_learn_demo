package main

import "fmt"

/*
Array函数的参数
如果你是一个C或则C++开发者，那么数组对你而言就是指针。
当你向函数中传递数组时，函数会参照相同的内存区域，这样它们就可以修改原始的数据。
Go中的数组是数值，因此当你向函数中传递数组时，函数会得到原始数组数据的一份复制。如果你打算更新数组的数据，这将会是个问题。
*/

// 值传递的 例子：
/*
func main() {
    x := [3]int{1,2,3}
    // 这边的 x 数组是值传递
    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(x)
    fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])
}
*/

// 如果你需要更新原始数组的数据，你可以使用数组指针类型。
/*func main() {
    x := [3]int{1,2,3}
    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) //prints &[7 2 3]
        fmt.Println(*arr) //prints [7 2 3]
    }(&x)
    fmt.Println(x) //prints [7 2 3]
}*/


// 另一个选择是使用slice。即使你的函数得到了slice变量的一份拷贝，它依旧会参照原始的数据。
func main() {
    x := []int{1,2,3}
    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) //prints [7 2 3]
    }(x)
    fmt.Println(x) //prints [7 2 3]
}