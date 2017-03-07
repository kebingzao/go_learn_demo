package main

import (
	"fmt"
	"time"
)

func main() {

	go say("world")
	say("hello")

	fmt.Println("---------------1")

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fmt.Println("---------------2")

	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	fmt.Println("---------------3")
	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}

	fmt.Println("---------------4")
	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("---------------4.5")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c4)
		}
		quit <- 0
		fmt.Println("---------------4.6")
	}()

	fmt.Println("---------------4.7")
	fibonacci2(c4, quit)

	fmt.Println("---------------5")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick. ")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("---------------x")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}