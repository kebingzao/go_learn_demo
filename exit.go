package main

import "fmt"
import "os"

// GO 退出程序实例
// http://www.yiibai.com/go/golang-exit.html

func main() {

	// `defer`s will _not_ be run when using `os.Exit`, so
	// this `fmt.Println` will never be called.
	// 使用 os.Exit 退出的时候， defer 不会被执行
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer
// return value from `main` to indicate exit status. If
// you'd like to exit with a non-zero status you should
// use `os.Exit`.