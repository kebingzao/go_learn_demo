package main

import (
	"fmt"
)

func main() {
	//var a [2]string
	//a[0] = "hello"
	//a[1] = "world"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)

	//p := []int{2,3,5,7,11,13}
	//fmt.Println("p ==", p)
	//
	//for i:=0; i<len(p); i++ {
	//	fmt.Printf("p[%d] == %d\n", i, p[i])
	//}
	//
	//fmt.Println("p[1:4] ==", p[1:4])
	//fmt.Println("p[:3] ==", p[:3])
	//fmt.Println("p[4:] ==", p[4:])


	//a := make([]int, 5)
	//fmt.Println("a", a)
	//b := make([]int, 0, 5)
	//fmt.Println("b", b)
	//c := b[:2]
	//fmt.Println("c", c)
	//d := c[2:5]
	//fmt.Println("d", d)


	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}