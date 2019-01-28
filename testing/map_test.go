package main

import (
	"sync"
	"testing"
	"strconv"
	"fmt"
)

// M
/*
type M struct {
	Map    map[string]string
}

// Set ...
func (m *M) Set(key, value string) {
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}
*/

// 换成有锁的
// M
type M struct {
	Map    map[string]string
	lock sync.RWMutex // 加锁
}

// Set ...
func (m *M) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

// Get ...
func (m *M) Get(key string) string {
	return m.Map[key]
}


// TestMap  ...
func TestMap(t *testing.T) {
	c := M{Map: make(map[string]string)}
	wg := sync.WaitGroup{}
	// 当为 20 个的时候，goroutines 跑的时候是正常
	//for i := 0; i < 21; i++ {
/*
F:\airdroid_code\go\src\go_learn_demo\testing>go test map_test.go -v
	=== RUN   TestMap
	k=:20,v:=20
	k=:14,v:=14
	k=:10,v:=10
	k=:12,v:=12
	k=:13,v:=13
	k=:17,v:=17
	k=:11,v:=11
	k=:18,v:=18
	k=:16,v:=16
	k=:4,v:=4
	k=:0,v:=0
	k=:1,v:=1
	k=:2,v:=2
	k=:15,v:=15
	k=:3,v:=3
	k=:7,v:=7
	k=:5,v:=5
	k=:19,v:=19
	k=:8,v:=8
	k=:6,v:=6
	k=:9,v:=9
	ok finished.
	--- PASS: TestMap (0.00s)
	PASS
	ok      command-line-arguments  0.303s
*/

	// 但是一旦再多，比如 71 的时候，就会报读的并发错误
	/*
	fatal error: concurrent map writes

	goroutine 75 [running]:
	runtime.throw(0x13b2011, 0x15)
	*/
	// 因此这时候就要加锁
	for i := 0; i < 171; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c.Set(k, v)
			fmt.Printf("k=:%v,v:=%v\n", k, c.Get(k))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("ok finished.")
}