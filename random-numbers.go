package main

import "fmt"
import "math/rand"

func main() {
	p := fmt.Println
	for i:=1 ; i < 100; i++ {
		// 生成 0 <= n < 50 的随机数
		p(rand.Intn(50))
	}
	for i:=1 ; i < 100; i++ {
		// 生成 0.0 <= f < 1.0 之间的随机数
		p(rand.Float64())
	}
}