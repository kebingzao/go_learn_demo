package main

import (
	"fmt"
)



func main() {
	//defer fmt.Println("world")

	//fmt.Println("hello")

	fmt.Println("counting")

	for i:=1; i<10; i++ {
		defer fmt.Println(i)
	}
}