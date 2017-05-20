package main

import "fmt"
import "math/rand"

func main() {
	p := fmt.Println
	for i:=1 ; i < 100; i++ {
		p(rand.Intn(100))
	}
	for i:=1 ; i < 100; i++ {
		p(rand.Float64() * 10)
	}
}