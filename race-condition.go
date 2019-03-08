package main

// url: https://larrylu.blog/race-condition-in-golang-c49a6e242259

import (
	"fmt"
	"sync"
)
// 这时候因为有多核CPU，兩個 CPU 同時去拿變數 a 的值，各自加 1 後存回，導致 a 只被加了一次，因此結果（9903）會小於正確的 10000
/*func main() {
	a := 0
	times := 10000
	c := make(chan bool)

	for i := 0; i < times; i++ {
		go func() {
			a++
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}*/

// 使用互斥锁，可以保证同時只能有一個 goroutine 做 a++

func main() {
	a := 0
	times := 10000
	c := make(chan bool)

	var m sync.Mutex

	for i := 0; i < times; i++ {
		go func() {
			m.Lock() // 取得鎖
			a++
			m.Unlock() // 釋放鎖
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}


/*
F:\airdroid_code\go\src\go_learn_demo>go run -race test.go
==================
WARNING: DATA RACE
Read at 0x00c04203c1d0 by goroutine 7:
main.main.func1()
F:/airdroid_code/go/src/go_learn_demo/test.go:15 +0x46

Previous write at 0x00c04203c1d0 by goroutine 6:
main.main.func1()
F:/airdroid_code/go/src/go_learn_demo/test.go:15 +0x5f

Goroutine 7 (running) created at:
main.main()
F:/airdroid_code/go/src/go_learn_demo/test.go:17 +0xd2

Goroutine 6 (running) created at:
main.main()
F:/airdroid_code/go/src/go_learn_demo/test.go:17 +0xd2
==================
race: limit on 8192 simultaneously alive goroutines is exceeded, dying
exit status 66
*/

