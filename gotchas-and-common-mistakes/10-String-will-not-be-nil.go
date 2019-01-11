package main

/*
字符串不会为nil
这对于经常使用nil分配字符串变量的开发者而言是个需要注意的地方。
*/

//失败的例子：
/*func main() {
    var x string = nil //error
    if x == nil { //error
        x = "default"
    }
}*/

/*
print::
gotchas-and-common-mistakes\10-String-will-not-be-nil.go:10: cannot use nil as type string in assignment
gotchas-and-common-mistakes\10-String-will-not-be-nil.go:11: invalid operation: x == nil (mismatched types string and nil)
*/

// 正确的做法：
func main() {
    var x string //defaults to "" (zero value)
    if x == "" {
        x = "default"
    }
}